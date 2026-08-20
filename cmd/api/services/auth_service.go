package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

const (
	defaultNonceExpiry = 5 * time.Minute
	// ecdsaSignatureLen is the expected length of an ECDSA signature: r (32) + s (32) + v (1).
	ecdsaSignatureLen = 65
	// ecdsaRecoveryIDOffset bridges MetaMask and go-ethereum conventions for the recovery ID (v).
	// MetaMask sends v as 27 or 28 (Ethereum convention), but go-ethereum's SigToPub expects 0 or 1.
	ecdsaRecoveryIDOffset = 27
)

var _ core.AuthService = (*authService)(nil)

type authService struct {
	userRepo          core.UserRepository
	nonceRepo         core.NonceRepository
	oauthProviderRepo core.UserOAuthProviderRepository
	walletRepo        core.UserWalletRepository
	txer              core.Transactor
	ramClient         core.RaylsAccessManagerClient
	provisioner       core.ProvisioningService // optional; nil disables auto-provision (new users stay waiting_role_assignment)
	baseURL           string
	chainless         bool
	log               logger.Logger
}

func NewAuthService(
	userRepo core.UserRepository,
	nonceRepo core.NonceRepository,
	oauthProviderRepo core.UserOAuthProviderRepository,
	walletRepo core.UserWalletRepository,
	txer core.Transactor,
	ramClient core.RaylsAccessManagerClient,
	provisioner core.ProvisioningService,
	baseURL string,
	chainless bool,
	log logger.Logger,
) core.AuthService {
	return &authService{
		userRepo:          userRepo,
		nonceRepo:         nonceRepo,
		oauthProviderRepo: oauthProviderRepo,
		walletRepo:        walletRepo,
		txer:              txer,
		ramClient:         ramClient,
		provisioner:       provisioner,
		baseURL:           baseURL,
		chainless:         chainless,
		log:               log,
	}
}

// GenerateChallenge creates a SIWE challenge message and nonce for the given wallet address.
func (s *authService) GenerateChallenge(ctx context.Context, walletAddress string) (string, string, error) {
	nonceBytes := make([]byte, 32)
	if _, err := rand.Read(nonceBytes); err != nil {
		return "", "", core.NewInternalError("generate nonce", err)
	}
	nonceHex := hex.EncodeToString(nonceBytes)

	_ = s.nonceRepo.DeleteExpired(ctx)

	message := buildSIWEMessage(walletAddress, "rayls-ops-api", s.baseURL, nonceHex)

	nonceRecord := &domain.Nonce{
		WalletAddress: strings.ToLower(walletAddress),
		Nonce:         nonceHex,
		Message:       message,
		ExpiresAt:     timeNow().Add(defaultNonceExpiry),
	}
	if err := s.nonceRepo.Create(ctx, nonceRecord); err != nil {
		return "", "", core.NewInternalError("store nonce", err)
	}

	return message, nonceHex, nil
}

// VerifySIWE validates the SIWE signature and applies the v1.4 login decision tree.
// Returns (user, roles, nil) only when the user is eligible to receive a JWT.
// Auto-registers new wallets on first login (status=waiting_role_assignment).
func (s *authService) VerifySIWE(
	ctx context.Context,
	walletAddress, signature, nonce string,
) (*domain.User, []string, error) {
	lowerAddr := strings.ToLower(walletAddress)

	nonceRecord, err := s.nonceRepo.FindValidAndMarkUsed(ctx, lowerAddr, nonce)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, nil, core.NewUnauthorizedError("nonce not found, expired, or already used")
		}
		return nil, nil, core.NewInternalError("consume nonce", err)
	}

	recoveredAddr, err := recoverAddress(nonceRecord.Message, signature)
	if err != nil {
		return nil, nil, core.NewUnauthorizedError("signature verification failed")
	}

	if !strings.EqualFold(recoveredAddr, walletAddress) {
		return nil, nil, core.NewUnauthorizedError("recovered address does not match claimed address")
	}

	wallet, err := s.walletRepo.FindByRaylsAddress(ctx, walletAddress)
	if err != nil {
		if !errors.Is(err, core.ErrRecordNotFound) {
			return nil, nil, core.NewInternalError("find wallet by address", err)
		}
		// Auto-register: new wallet on first login.
		if regErr := s.registerSIWEUser(ctx, walletAddress); regErr != nil {
			return nil, nil, regErr
		}
		return nil, nil, &core.WalletRegisteredError{}
	}

	user, err := s.userRepo.FindByID(ctx, wallet.UserID)
	if err != nil {
		return nil, nil, core.NewInternalError("find user by id", err)
	}

	roles, loginErr := s.applyLoginDecisionTree(ctx, user, walletAddress)
	if loginErr != nil {
		return nil, nil, loginErr
	}
	return user, roles, nil
}

// FindOrCreateOAuthUser looks up an OAuth user and applies the v1.4 login decision tree.
// emailVerified enables the one-time bootstrap email fallback (SDD section 2c).
func (s *authService) FindOrCreateOAuthUser(
	ctx context.Context,
	provider domain.OAuthProvider,
	oauthID, name, email string,
	emailVerified bool,
) (*domain.User, []string, error) {
	existing, err := s.oauthProviderRepo.FindByProviderAndID(ctx, provider, oauthID)
	if err == nil {
		user, err := s.userRepo.FindByID(ctx, existing.UserID)
		if err != nil {
			return nil, nil, core.NewInternalError("find user by id", err)
		}
		wallet, _ := s.walletRepo.FindByUserID(ctx, user.ID)
		walletAddr := ""
		if wallet != nil {
			walletAddr = wallet.RaylsAddress
		}
		roles, loginErr := s.applyLoginDecisionTree(ctx, user, walletAddr)
		if loginErr != nil {
			return nil, nil, loginErr
		}
		return user, roles, nil
	}

	if !errors.Is(err, core.ErrRecordNotFound) {
		return nil, nil, core.NewInternalError("find OAuth provider", err)
	}

	// Email fallback: link bootstrap admin's OAuth identity to pre-created user.
	if emailVerified && email != "" {
		existingUser, findErr := s.userRepo.FindByEmail(ctx, email)
		if findErr == nil {
			// If the user has a provider record for this provider but with no OAuthID, update it.
			existingLink, linkErr := s.oauthProviderRepo.FindByProviderAndUserID(ctx, provider, existingUser.ID)
			if linkErr == nil && existingLink.OAuthID == "" {
				if updateErr := s.oauthProviderRepo.UpdateOAuthID(ctx, existingLink.ID, oauthID); updateErr != nil {
					return nil, nil, core.NewInternalError("update oauth id", updateErr)
				}
				wallet, _ := s.walletRepo.FindByUserID(ctx, existingUser.ID)
				walletAddr := ""
				if wallet != nil {
					walletAddr = wallet.RaylsAddress
				}
				roles, loginErr := s.applyLoginDecisionTree(ctx, existingUser, walletAddr)
				if loginErr != nil {
					return nil, nil, loginErr
				}
				return existingUser, roles, nil
			} else if linkErr != nil && !errors.Is(linkErr, core.ErrRecordNotFound) {
				return nil, nil, core.NewInternalError("find oauth provider by user", linkErr)
			}

			count, countErr := s.oauthProviderRepo.CountByUserID(ctx, existingUser.ID)
			if countErr != nil {
				return nil, nil, core.NewInternalError("count oauth providers", countErr)
			}
			if count == 0 {
				link := &domain.UserOAuthProvider{
					UserID:   existingUser.ID,
					Provider: provider,
					OAuthID:  oauthID,
					Email:    email,
				}
				if linkErr := s.oauthProviderRepo.Create(
					ctx,
					link,
				); linkErr != nil &&
					!errors.Is(linkErr, core.ErrDuplicateOAuthLink) {
					return nil, nil, core.NewInternalError("link oauth provider", linkErr)
				}
				wallet, _ := s.walletRepo.FindByUserID(ctx, existingUser.ID)
				walletAddr := ""
				if wallet != nil {
					walletAddr = wallet.RaylsAddress
				}
				roles, loginErr := s.applyLoginDecisionTree(ctx, existingUser, walletAddr)
				if loginErr != nil {
					return nil, nil, loginErr
				}
				return existingUser, roles, nil
			}
			// User exists with other providers already linked. Refuse rather than silently
			// linking a new identity to an established account: this path is reached with
			// emailVerified=true, which the email sign-up endpoint asserts without actually
			// verifying anything, so linking here would let anyone claim an existing account
			// by typing its address. Report which provider owns the account so the caller can
			// point the user at the login method that works.
			return existingUser, nil, &core.EmailAlreadyLinkedError{
				Provider: s.linkedProviderName(ctx, existingUser.ID),
			}
		} else if !errors.Is(findErr, core.ErrRecordNotFound) {
			return nil, nil, core.NewInternalError("find user by email", findErr)
		}
		// No user with this email — fall through to new-user creation below.
	}

	// New user: auto-register with waiting_role_assignment.
	user := &domain.User{
		Name:     name,
		Email:    email,
		IsActive: true,
		Status:   domain.UserStatusWaitingRoleAssignment,
	}

	if err := s.txer.WithTransaction(ctx, func(txCtx context.Context) error {
		if err := s.userRepo.Create(txCtx, user); err != nil {
			if errors.Is(err, core.ErrDuplicateEmail) {
				return core.NewValidationError("email", "email already registered")
			}
			return core.NewInternalError("create OAuth user", err)
		}

		oauthProvider := &domain.UserOAuthProvider{
			UserID:   user.ID,
			Provider: provider,
			OAuthID:  oauthID,
			Email:    email,
		}
		if err := s.oauthProviderRepo.Create(txCtx, oauthProvider); err != nil {
			if errors.Is(err, core.ErrDuplicateOAuthLink) {
				return core.ErrDuplicateOAuthLink
			}
			return core.NewInternalError("create OAuth provider link", err)
		}
		return nil
	}); err != nil {
		// Race condition: retry lookup after duplicate OAuth link on concurrent creation.
		if errors.Is(err, core.ErrDuplicateOAuthLink) {
			if raceExisting, findErr := s.oauthProviderRepo.FindByProviderAndID(
				ctx,
				provider,
				oauthID,
			); findErr == nil {
				raceUser, findErr := s.userRepo.FindByID(ctx, raceExisting.UserID)
				if findErr != nil {
					return nil, nil, core.NewInternalError("find user after race", findErr)
				}
				wallet, _ := s.walletRepo.FindByUserID(ctx, raceUser.ID)
				walletAddr := ""
				if wallet != nil {
					walletAddr = wallet.RaylsAddress
				}
				roles, loginErr := s.applyLoginDecisionTree(ctx, raceUser, walletAddr)
				if loginErr != nil {
					return nil, nil, loginErr
				}
				return raceUser, roles, nil
			}
		}
		return nil, nil, err
	}

	s.log.Info("OAuth user auto-registered", "provider", provider, "email", email)

	// New users start at waiting_role_assignment; the decision tree auto-provisions them
	// (custody wallet + role_assigned) when a provisioner is wired, else returns approval-pending.
	roles, loginErr := s.applyLoginDecisionTree(ctx, user, "")
	if loginErr != nil {
		return nil, nil, loginErr
	}
	return user, roles, nil
}

// applyLoginDecisionTree implements the v1.4 login decision tree (SDD section 2b).
// Returns the wallet's on-chain roles on success; callers must pass them to IssueToken.
// walletAddress may be empty for users with status=waiting_role_assignment (GetRoles not reached).
func (s *authService) applyLoginDecisionTree(
	ctx context.Context,
	user *domain.User,
	walletAddress string,
) ([]string, error) {
	if !user.IsActive {
		return nil, &core.AccountSuspendedError{}
	}

	if s.provisioner == nil {
		// Nothing can activate a waiting account, so it stays pending. An already-activated
		// user is unaffected.
		if user.Status == domain.UserStatusWaitingRoleAssignment {
			return nil, &core.RoleAssignmentPendingError{}
		}
	} else {
		// Provision on EVERY login, not only for waiting users. It is idempotent (an
		// existing wallet is reused, an activated status is left alone), and running it
		// unconditionally is what heals accounts that were activated before this service
		// minted wallets — they sit at role_assigned with no wallet, and gating on
		// "waiting" would leave them wallet-less forever, failing on every chain with
		// "user has no custody wallet".
		if err := s.provisioner.Provision(ctx, user); err != nil {
			return nil, core.NewInternalError("provision user", err)
		}
		wallet, _ := s.walletRepo.FindByUserID(ctx, user.ID)
		if wallet != nil {
			walletAddress = wallet.RaylsAddress
		}
	}

	// Chain-less (CHAINLESS=true — the RayUp dev flow, where the chain is created later
	// from the UI): there is no AccessManager to read roles from, so the on-chain check
	// below would find none and reject every login as approval-pending.
	//
	// Grant the operator role rather than an empty set: the role-gated routes (/api/tokens
	// and friends) require it, so returning nothing would let the user log in and then 403
	// on every request — a chain-less instance would be authenticated but unusable. This is
	// as much as we can verify with no chain: the user is authenticated and already
	// role_assigned in our own DB. Dev-only by construction — CHAINLESS is opt-in, and once
	// a chain exists it is off and the real AccessManager check applies again.
	if s.chainless {
		s.log.Warn("chain-less login: granting operator role without on-chain verification",
			"userID", user.ID, "role", domain.RolePrivacyNodeOperator)
		return []string{domain.RolePrivacyNodeOperator}, nil
	}

	// status = role_assigned: verify on-chain roles.
	roles, err := s.ramClient.GetRoles(ctx, walletAddress)
	if err != nil {
		s.log.Error("GetRoles RPC failed at login", "userID", user.ID, "address", walletAddress, "error", err)
		return nil, &core.ServiceUnavailableError{}
	}

	if len(roles) == 0 {
		return nil, &core.RoleAssignmentPendingError{}
	}

	return roles, nil
}

// linkedProviderName reports which login method userID's account already uses, for the
// EMAIL_ALREADY_LINKED message. The repository has no "list providers for a user" query, so it
// probes the known providers in turn; a user reaching this path has at least one link, and the
// first hit is the one to name.
//
// Cosmetic by design: it only sharpens an error message that is already being returned, so a
// lookup failure yields "" (the caller renders a generic message) rather than masking the real
// error with an internal one.
func (s *authService) linkedProviderName(ctx context.Context, userID uuid.UUID) string {
	for _, provider := range []domain.OAuthProvider{
		domain.OAuthProviderGoogle,
		domain.OAuthProviderMicrosoft,
		domain.OAuthProviderSIWE,
		domain.OAuthProviderEmail,
	} {
		if _, err := s.oauthProviderRepo.FindByProviderAndUserID(ctx, provider, userID); err == nil {
			return provider.String()
		}
	}
	return ""
}

// registerSIWEUser auto-creates a new user with an empty name/email and links the SIWE wallet.
// Used on first SIWE login (replaces the explicit POST /auth/register/siwe flow).
func (s *authService) registerSIWEUser(ctx context.Context, walletAddress string) error {
	user := &domain.User{
		IsActive: true,
		Status:   domain.UserStatusWaitingRoleAssignment,
	}

	if err := s.txer.WithTransaction(ctx, func(txCtx context.Context) error {
		if err := s.userRepo.Create(txCtx, user); err != nil {
			return core.NewInternalError("create SIWE user", err)
		}

		wallet := &domain.UserWallet{
			UserID:          user.ID,
			RaylsAddress:    walletAddress,
			CustodyProvider: domain.CustodyProviderSelf,
			IsActive:        true,
		}
		if err := s.walletRepo.Create(txCtx, wallet); err != nil {
			if errors.Is(err, core.ErrDuplicateWalletAddress) {
				return core.NewValidationError("walletAddress", "wallet address already registered")
			}
			return core.NewInternalError("create user wallet", err)
		}

		siweLink := &domain.UserOAuthProvider{
			UserID:        user.ID,
			Provider:      domain.OAuthProviderSIWE,
			OAuthID:       strings.ToLower(walletAddress),
			WalletAddress: walletAddress,
		}
		if err := s.oauthProviderRepo.Create(txCtx, siweLink); err != nil {
			return core.NewInternalError("create SIWE provider link", err)
		}
		return nil
	}); err != nil {
		return err
	}

	s.log.Info("SIWE user auto-registered", "address", walletAddress)
	return nil
}

func buildSIWEMessage(address, domain, uri, nonce string) string {
	return fmt.Sprintf(`%s wants you to sign in with your Ethereum account:
%s

Sign in to %s

URI: %s
Version: 1
Nonce: %s
Issued At: %s`, domain, address, domain, uri, nonce, timeNow().UTC().Format("2006-01-02T15:04:05Z"))
}

// recoverAddress is a package-level variable to allow injection in tests.
var recoverAddress = defaultRecoverAddress

func defaultRecoverAddress(message, signatureHex string) (string, error) {
	sig, err := hex.DecodeString(strings.TrimPrefix(signatureHex, "0x"))
	if err != nil {
		return "", fmt.Errorf("failed to decode signature: %w", err)
	}

	if len(sig) != ecdsaSignatureLen {
		return "", fmt.Errorf("invalid signature length: expected %d bytes, got %d", ecdsaSignatureLen, len(sig))
	}

	if sig[ecdsaSignatureLen-1] >= ecdsaRecoveryIDOffset {
		sig[ecdsaSignatureLen-1] -= ecdsaRecoveryIDOffset
	}

	prefixedMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(prefixedMsg))

	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return "", fmt.Errorf("failed to recover public key: %w", err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return common.HexToAddress(recoveredAddr.Hex()).Hex(), nil
}

var timeNow = defaultTimeNow

func defaultTimeNow() time.Time {
	return time.Now().UTC()
}
