package services

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

func newAuthSvc(
	userRepo *testutil.FakeUserRepository,
	oauthRepo *testutil.FakeUserOAuthProviderRepository,
	walletRepo *testutil.FakeUserWalletRepository,
	ramClient *testutil.FakeRaylsAccessManagerClient,
) core.AuthService {
	return NewAuthService(
		userRepo,
		&testutil.FakeNonceRepository{},
		oauthRepo,
		walletRepo,
		&testutil.FakeTransactor{},
		ramClient,
		nil,
		"http://localhost:8080",
		false,
		&testutil.StubLogger{},
	)
}

func newAuthSvcWithNonce(
	nonceRepo *testutil.FakeNonceRepository,
	userRepo *testutil.FakeUserRepository,
	walletRepo *testutil.FakeUserWalletRepository,
	ramClient *testutil.FakeRaylsAccessManagerClient,
) core.AuthService {
	return NewAuthService(
		userRepo,
		nonceRepo,
		&testutil.FakeUserOAuthProviderRepository{},
		walletRepo,
		&testutil.FakeTransactor{},
		ramClient,
		nil,
		"http://localhost:8080",
		false,
		&testutil.StubLogger{},
	)
}

// newAuthSvcChainlessWithProvisioner builds a chain-less auth service wired with a
// provisioner, so new/waiting users are auto-provisioned and granted the operator role.
func newAuthSvcChainlessWithProvisioner(
	userRepo *testutil.FakeUserRepository,
	oauthRepo *testutil.FakeUserOAuthProviderRepository,
	walletRepo *testutil.FakeUserWalletRepository,
	provisioner core.ProvisioningService,
) core.AuthService {
	return NewAuthService(
		userRepo,
		&testutil.FakeNonceRepository{},
		oauthRepo,
		walletRepo,
		&testutil.FakeTransactor{},
		&testutil.FakeRaylsAccessManagerClient{},
		provisioner,
		"http://localhost:8080",
		true, // chainless: grant operator from role_assigned status alone
		&testutil.StubLogger{},
	)
}

// newAuthSvcChainless builds a chain-less auth service with no provisioner, exposing the RAM
// client so a test can assert the chain-less path never reaches it.
func newAuthSvcChainless(
	userRepo *testutil.FakeUserRepository,
	oauthRepo *testutil.FakeUserOAuthProviderRepository,
	walletRepo *testutil.FakeUserWalletRepository,
	ramClient *testutil.FakeRaylsAccessManagerClient,
) core.AuthService {
	return NewAuthService(
		userRepo,
		&testutil.FakeNonceRepository{},
		oauthRepo,
		walletRepo,
		&testutil.FakeTransactor{},
		ramClient,
		nil,
		"http://localhost:8080",
		true, // chainless: grant operator from role_assigned status alone
		&testutil.StubLogger{},
	)
}

// ── FindOrCreateOAuthUser ─────────────────────────────────────────────────────

func TestAuthService_FindOrCreateOAuthUser_ReturnsDuplicateEmailValidationError(t *testing.T) {
	// When user creation fails with a duplicate email, a validation error is returned

	userRepo := testutil.NewFakeUserRepository()
	userRepo.CreateErr = core.ErrDuplicateEmail
	svc := newAuthSvc(
		userRepo,
		&testutil.FakeUserOAuthProviderRepository{},
		&testutil.FakeUserWalletRepository{},
		&testutil.FakeRaylsAccessManagerClient{},
	)

	_, _, err := svc.FindOrCreateOAuthUser(
		context.Background(),
		domain.OAuthProviderGoogle,
		"oauth-999",
		"Dave",
		"dave@example.com",
		false,
	)

	var validationErr *core.ValidationError
	require.Error(t, err)
	require.True(t, errors.As(err, &validationErr))
	assert.Equal(t, "email", validationErr.Field)
}

func TestAuthService_FindOrCreateOAuthUser_ReturnsRoleAssignmentPendingForExistingUserWaitingRole(t *testing.T) {
	// An existing user with status=waiting_role_assignment gets RoleAssignmentPendingError

	existingUser := domain.User{
		Name:     "Alice",
		IsActive: true,
		Status:   domain.UserStatusWaitingRoleAssignment,
	}
	existingUser.ID = uuid.New()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{existingUser}

	oauthRepo := &testutil.FakeUserOAuthProviderRepository{
		Providers: []domain.UserOAuthProvider{{
			UserID:   existingUser.ID,
			Provider: domain.OAuthProviderGoogle,
			OAuthID:  "oauth-123",
		}},
	}

	svc := newAuthSvc(
		userRepo,
		oauthRepo,
		&testutil.FakeUserWalletRepository{},
		&testutil.FakeRaylsAccessManagerClient{},
	)

	_, _, err := svc.FindOrCreateOAuthUser(
		context.Background(),
		domain.OAuthProviderGoogle,
		"oauth-123",
		"Alice",
		"",
		false,
	)

	var pending *core.RoleAssignmentPendingError
	require.Error(t, err)
	require.True(t, errors.As(err, &pending))
}

func TestAuthService_FindOrCreateOAuthUser_AutoProvisionsNewUserWhenProvisionerWired(t *testing.T) {
	// With a provisioner, a brand-new user is auto-provisioned and logs in (no approval pending)

	userRepo := testutil.NewFakeUserRepository()
	provisioner := testutil.NewFakeProvisioningService()
	svc := newAuthSvcChainlessWithProvisioner(
		userRepo,
		&testutil.FakeUserOAuthProviderRepository{},
		&testutil.FakeUserWalletRepository{},
		provisioner,
	)

	user, roles, err := svc.FindOrCreateOAuthUser(
		context.Background(),
		domain.OAuthProviderEmail,
		"new@example.com",
		"New",
		"new@example.com",
		true,
	)

	require.NoError(t, err)
	require.Len(t, provisioner.Calls, 1)
	assert.Equal(t, domain.UserStatusRoleAssigned, user.Status)
	assert.Equal(t, []string{domain.RolePrivacyNodeOperator}, roles)
}

func TestAuthService_FindOrCreateOAuthUser_AutoProvisionsExistingWaitingUser(t *testing.T) {
	// An existing waiting_role_assignment user is auto-provisioned on login when a provisioner is wired

	existingUser := domain.User{
		Name:     "Alice",
		Email:    "alice@example.com",
		IsActive: true,
		Status:   domain.UserStatusWaitingRoleAssignment,
	}
	existingUser.ID = uuid.New()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{existingUser}
	oauthRepo := &testutil.FakeUserOAuthProviderRepository{
		Providers: []domain.UserOAuthProvider{
			{UserID: existingUser.ID, Provider: domain.OAuthProviderEmail, OAuthID: "alice@example.com"},
		},
	}
	provisioner := testutil.NewFakeProvisioningService()
	svc := newAuthSvcChainlessWithProvisioner(userRepo, oauthRepo, &testutil.FakeUserWalletRepository{}, provisioner)

	_, roles, err := svc.FindOrCreateOAuthUser(
		context.Background(),
		domain.OAuthProviderEmail,
		"alice@example.com",
		"Alice",
		"alice@example.com",
		true,
	)

	require.NoError(t, err)
	require.Len(t, provisioner.Calls, 1)
	assert.Equal(t, []string{domain.RolePrivacyNodeOperator}, roles)
}

func TestAuthService_FindOrCreateOAuthUser_ReturnsUserWhenRoleAssignedWithRoles(t *testing.T) {
	// An existing user with status=role_assigned and on-chain roles returns the user

	existingUser := domain.User{
		Name:     "Alice",
		IsActive: true,
		Status:   domain.UserStatusRoleAssigned,
	}
	existingUser.ID = uuid.New()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{existingUser}

	walletAddr := "0xabc123def456789012345678901234567890abcd"
	walletRepo := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{UserID: existingUser.ID, RaylsAddress: walletAddr, IsActive: true}},
	}

	oauthRepo := &testutil.FakeUserOAuthProviderRepository{
		Providers: []domain.UserOAuthProvider{{
			UserID:   existingUser.ID,
			Provider: domain.OAuthProviderGoogle,
			OAuthID:  "oauth-123",
		}},
	}

	ramClient := &testutil.FakeRaylsAccessManagerClient{Roles: []string{"PRIVACY_NODE_OPERATOR"}}
	svc := newAuthSvc(userRepo, oauthRepo, walletRepo, ramClient)

	user, roles, err := svc.FindOrCreateOAuthUser(
		context.Background(),
		domain.OAuthProviderGoogle,
		"oauth-123",
		"Alice",
		"",
		false,
	)

	require.NoError(t, err)
	assert.Equal(t, existingUser.ID, user.ID)
	assert.Equal(t, []string{"PRIVACY_NODE_OPERATOR"}, roles)
}

func TestAuthService_FindOrCreateOAuthUser_ChainlessRoleAssignedBypassesAccessManager(t *testing.T) {
	// Chain-less: a role_assigned user is granted the operator role from its stored status alone,
	// without consulting the AccessManager. The fake is loaded with a DIFFERENT role so a read
	// would be visible in the result, and CallCount proves the RPC never happened.

	existingUser := domain.User{Name: "Alice", IsActive: true, Status: domain.UserStatusRoleAssigned}
	existingUser.ID = uuid.New()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{existingUser}

	walletRepo := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{
			UserID:       existingUser.ID,
			RaylsAddress: "0xabc123def456789012345678901234567890abcd",
			IsActive:     true,
		}},
	}
	oauthRepo := &testutil.FakeUserOAuthProviderRepository{
		Providers: []domain.UserOAuthProvider{{
			UserID:   existingUser.ID,
			Provider: domain.OAuthProviderGoogle,
			OAuthID:  "oauth-123",
		}},
	}

	ramClient := &testutil.FakeRaylsAccessManagerClient{Roles: []string{"SOME_OTHER_ROLE"}}
	svc := newAuthSvcChainless(userRepo, oauthRepo, walletRepo, ramClient)

	user, roles, err := svc.FindOrCreateOAuthUser(
		context.Background(),
		domain.OAuthProviderGoogle,
		"oauth-123",
		"Alice",
		"",
		false,
	)

	require.NoError(t, err)
	assert.Equal(t, existingUser.ID, user.ID)
	assert.Equal(t, []string{domain.RolePrivacyNodeOperator}, roles)
	assert.Zero(t, ramClient.CallCount, "the chain-less path must not consult the AccessManager")
}

func TestAuthService_FindOrCreateOAuthUser_ChainlessIgnoresAccessManagerFailure(t *testing.T) {
	// The bypass is unconditional: a chain-less instance logs in even when the AccessManager
	// would error, which is the whole point (there is no chain to ask).

	existingUser := domain.User{Name: "Alice", IsActive: true, Status: domain.UserStatusRoleAssigned}
	existingUser.ID = uuid.New()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{existingUser}

	walletRepo := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{
			UserID:       existingUser.ID,
			RaylsAddress: "0xabc123def456789012345678901234567890abcd",
			IsActive:     true,
		}},
	}
	oauthRepo := &testutil.FakeUserOAuthProviderRepository{
		Providers: []domain.UserOAuthProvider{{
			UserID:   existingUser.ID,
			Provider: domain.OAuthProviderGoogle,
			OAuthID:  "oauth-123",
		}},
	}

	ramClient := &testutil.FakeRaylsAccessManagerClient{Err: errors.New("no chain to dial")}
	svc := newAuthSvcChainless(userRepo, oauthRepo, walletRepo, ramClient)

	_, roles, err := svc.FindOrCreateOAuthUser(
		context.Background(),
		domain.OAuthProviderGoogle,
		"oauth-123",
		"Alice",
		"",
		false,
	)

	require.NoError(t, err)
	assert.Equal(t, []string{domain.RolePrivacyNodeOperator}, roles)
	assert.Zero(t, ramClient.CallCount)
}

func TestAuthService_FindOrCreateOAuthUser_ChainedRoleAssignedDoesConsultAccessManager(t *testing.T) {
	// The counterpart to the bypass test: with a chain, the same user DOES hit the
	// AccessManager, so the CallCount assertion above is meaningful rather than vacuous.

	existingUser := domain.User{Name: "Alice", IsActive: true, Status: domain.UserStatusRoleAssigned}
	existingUser.ID = uuid.New()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{existingUser}

	walletRepo := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{
			UserID:       existingUser.ID,
			RaylsAddress: "0xabc123def456789012345678901234567890abcd",
			IsActive:     true,
		}},
	}
	oauthRepo := &testutil.FakeUserOAuthProviderRepository{
		Providers: []domain.UserOAuthProvider{{
			UserID:   existingUser.ID,
			Provider: domain.OAuthProviderGoogle,
			OAuthID:  "oauth-123",
		}},
	}

	ramClient := &testutil.FakeRaylsAccessManagerClient{Roles: []string{"SOME_OTHER_ROLE"}}
	svc := newAuthSvc(userRepo, oauthRepo, walletRepo, ramClient)

	_, roles, err := svc.FindOrCreateOAuthUser(
		context.Background(),
		domain.OAuthProviderGoogle,
		"oauth-123",
		"Alice",
		"",
		false,
	)

	require.NoError(t, err)
	assert.Equal(t, []string{"SOME_OTHER_ROLE"}, roles, "the chained path returns what the AccessManager says")
	assert.Equal(t, 1, ramClient.CallCount)
}

func TestAuthService_FindOrCreateOAuthUser_ReturnsAccountSuspendedWhenDeactivated(t *testing.T) {
	// A deactivated user returns AccountSuspendedError

	deactivatedUser := domain.User{IsActive: false, Status: domain.UserStatusRoleAssigned}
	deactivatedUser.ID = uuid.New()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{deactivatedUser}

	oauthRepo := &testutil.FakeUserOAuthProviderRepository{
		Providers: []domain.UserOAuthProvider{{
			UserID:   deactivatedUser.ID,
			Provider: domain.OAuthProviderGoogle,
			OAuthID:  "oauth-456",
		}},
	}

	svc := newAuthSvc(
		userRepo,
		oauthRepo,
		&testutil.FakeUserWalletRepository{},
		&testutil.FakeRaylsAccessManagerClient{},
	)

	_, _, err := svc.FindOrCreateOAuthUser(
		context.Background(),
		domain.OAuthProviderGoogle,
		"oauth-456",
		"Bob",
		"",
		false,
	)

	var suspended *core.AccountSuspendedError
	require.Error(t, err)
	require.True(t, errors.As(err, &suspended))
}

func TestAuthService_FindOrCreateOAuthUser_AutoRegistersNewUserAndReturnsPending(t *testing.T) {
	// A new OAuth user is auto-registered with waiting_role_assignment and returns RoleAssignmentPendingError

	userRepo := testutil.NewFakeUserRepository()
	oauthRepo := &testutil.FakeUserOAuthProviderRepository{}
	svc := newAuthSvc(
		userRepo,
		oauthRepo,
		&testutil.FakeUserWalletRepository{},
		&testutil.FakeRaylsAccessManagerClient{},
	)

	_, _, err := svc.FindOrCreateOAuthUser(
		context.Background(),
		domain.OAuthProviderGoogle,
		"oauth-789",
		"Carol",
		"carol@example.com",
		false,
	)

	var pending *core.RoleAssignmentPendingError
	require.Error(t, err)
	require.True(t, errors.As(err, &pending))

	require.Len(t, userRepo.Users, 1)
	assert.Equal(t, domain.UserStatusWaitingRoleAssignment, userRepo.Users[0].Status)
	require.Len(t, oauthRepo.Providers, 1)
	assert.Equal(t, "oauth-789", oauthRepo.Providers[0].OAuthID)
}

func TestAuthService_FindOrCreateOAuthUser_EmailFallbackLinksBootstrapAdmin(t *testing.T) {
	// With emailVerified=true and zero existing providers, admin is linked by email and decision tree applied

	adminUser := domain.User{
		Email:    "admin@example.com",
		IsActive: true,
		Status:   domain.UserStatusRoleAssigned,
	}
	adminUser.ID = uuid.New()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{adminUser}

	walletAddr := "0xabc123def456789012345678901234567890abcd"
	walletRepo := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{UserID: adminUser.ID, RaylsAddress: walletAddr, IsActive: true}},
	}

	oauthRepo := &testutil.FakeUserOAuthProviderRepository{}
	ramClient := &testutil.FakeRaylsAccessManagerClient{Roles: []string{"PRIVACY_NODE_OPERATOR"}}

	svc := newAuthSvc(userRepo, oauthRepo, walletRepo, ramClient)

	user, _, err := svc.FindOrCreateOAuthUser(
		context.Background(),
		domain.OAuthProviderGoogle,
		"google-sub-001",
		"Admin",
		"admin@example.com",
		true,
	)

	require.NoError(t, err)
	assert.Equal(t, adminUser.ID, user.ID)
	require.Len(t, oauthRepo.Providers, 1, "OAuth link should be created")
	assert.Equal(t, "google-sub-001", oauthRepo.Providers[0].OAuthID)
}

func TestAuthService_FindOrCreateOAuthUser_EmailSignupOnGoogleAccountReturnsEmailAlreadyLinked(t *testing.T) {
	// Email sign-up reusing a Google account's address is refused with EmailAlreadyLinkedError,
	// naming google as the provider to use instead

	googleUser := domain.User{
		Email:    "alice@example.com",
		IsActive: true,
		Status:   domain.UserStatusRoleAssigned,
	}
	googleUser.ID = uuid.New()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{googleUser}

	oauthRepo := &testutil.FakeUserOAuthProviderRepository{
		Providers: []domain.UserOAuthProvider{{
			UserID:   googleUser.ID,
			Provider: domain.OAuthProviderGoogle,
			OAuthID:  "google-sub-777",
			Email:    "alice@example.com",
		}},
	}

	svc := newAuthSvc(
		userRepo,
		oauthRepo,
		&testutil.FakeUserWalletRepository{},
		&testutil.FakeRaylsAccessManagerClient{},
	)

	_, _, err := svc.FindOrCreateOAuthUser(
		context.Background(),
		domain.OAuthProviderEmail,
		"alice@example.com",
		"alice",
		"alice@example.com",
		true,
	)

	var linked *core.EmailAlreadyLinkedError
	require.Error(t, err)
	require.True(t, errors.As(err, &linked))
	assert.Equal(t, "google", linked.Provider)
	assert.Contains(t, linked.Message(), "google")
	assert.Len(t, userRepo.Users, 1, "no second account should be created for the same email")
	assert.Len(t, oauthRepo.Providers, 1, "the email identity must not be linked to the Google account")
}

func TestAuthService_FindOrCreateOAuthUser_EmailFallbackSkippedWhenNotVerified(t *testing.T) {
	// With emailVerified=false the email fallback is skipped — a new user is created instead

	adminUser := domain.User{
		Email:    "admin@example.com",
		IsActive: true,
		Status:   domain.UserStatusRoleAssigned,
	}
	adminUser.ID = uuid.New()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{adminUser}
	oauthRepo := &testutil.FakeUserOAuthProviderRepository{}

	svc := newAuthSvc(
		userRepo,
		oauthRepo,
		&testutil.FakeUserWalletRepository{},
		&testutil.FakeRaylsAccessManagerClient{},
	)

	_, _, err := svc.FindOrCreateOAuthUser(
		context.Background(),
		domain.OAuthProviderGoogle,
		"google-sub-002",
		"Admin",
		"admin@example.com",
		false,
	)

	var pending *core.RoleAssignmentPendingError
	require.Error(t, err)
	require.True(t, errors.As(err, &pending))
	assert.Len(t, userRepo.Users, 2, "new user should be created, not linked")
}

// ── VerifySIWE ───────────────────────────────────────────────────────────────

func TestAuthService_VerifySIWE_NonceNotFound_ReturnsUnauthorized(t *testing.T) {
	// An unknown or expired nonce returns an unauthorized error

	svc := newAuthSvcWithNonce(
		&testutil.FakeNonceRepository{},
		testutil.NewFakeUserRepository(),
		&testutil.FakeUserWalletRepository{},
		&testutil.FakeRaylsAccessManagerClient{},
	)

	_, _, err := svc.VerifySIWE(
		context.Background(),
		"0xabc123def456789012345678901234567890abcd",
		"0xsig",
		"bad-nonce",
	)

	var unauthErr *core.UnauthorizedError
	require.Error(t, err)
	require.True(t, errors.As(err, &unauthErr))
}

func TestAuthService_VerifySIWE_NonceLookupError_ReturnsInternalError(t *testing.T) {
	// A DB error during nonce lookup returns an internal error

	nonceRepo := &testutil.FakeNonceRepository{FindErr: errors.New("db timeout")}
	svc := newAuthSvcWithNonce(
		nonceRepo,
		testutil.NewFakeUserRepository(),
		&testutil.FakeUserWalletRepository{},
		&testutil.FakeRaylsAccessManagerClient{},
	)

	_, _, err := svc.VerifySIWE(
		context.Background(),
		"0xabc123def456789012345678901234567890abcd",
		"0xsig",
		"some-nonce",
	)

	var internalErr *core.InternalError
	require.Error(t, err)
	require.True(t, errors.As(err, &internalErr))
}

func TestAuthService_VerifySIWE_InvalidSignature_ReturnsUnauthorized(t *testing.T) {
	// A signature that fails recovery returns an unauthorized error

	const addr = "0xabc123def456789012345678901234567890abcd"
	nonceRepo := &testutil.FakeNonceRepository{
		Nonce: &domain.Nonce{WalletAddress: addr, Message: "test message"},
	}
	old := recoverAddress
	recoverAddress = func(_, _ string) (string, error) { return "", errors.New("bad sig") }
	defer func() { recoverAddress = old }()

	svc := newAuthSvcWithNonce(
		nonceRepo,
		testutil.NewFakeUserRepository(),
		&testutil.FakeUserWalletRepository{},
		&testutil.FakeRaylsAccessManagerClient{},
	)

	_, _, err := svc.VerifySIWE(context.Background(), addr, "0xbadsig", "nonce")

	var unauthErr *core.UnauthorizedError
	require.Error(t, err)
	require.True(t, errors.As(err, &unauthErr))
}

func TestAuthService_VerifySIWE_WalletNotFound_AutoRegistersAndReturnsWalletRegistered(t *testing.T) {
	// A valid signature for an unknown wallet auto-registers the user and returns WalletRegisteredError

	const addr = "0xabc123def456789012345678901234567890abcd"
	nonceRepo := &testutil.FakeNonceRepository{
		Nonce: &domain.Nonce{WalletAddress: addr, Message: "test message"},
	}
	old := recoverAddress
	recoverAddress = func(_, _ string) (string, error) { return addr, nil }
	defer func() { recoverAddress = old }()

	userRepo := testutil.NewFakeUserRepository()
	walletRepo := &testutil.FakeUserWalletRepository{}
	svc := newAuthSvcWithNonce(nonceRepo, userRepo, walletRepo, &testutil.FakeRaylsAccessManagerClient{})

	_, _, err := svc.VerifySIWE(context.Background(), addr, "0xsig", "nonce")

	var registered *core.WalletRegisteredError
	require.Error(t, err)
	require.True(t, errors.As(err, &registered))
	assert.Len(t, userRepo.Users, 1, "new user should be auto-registered")
	assert.Len(t, walletRepo.Wallets, 1, "wallet should be created")
}

func TestAuthService_VerifySIWE_WalletLookupError_ReturnsInternalError(t *testing.T) {
	// A DB error during wallet lookup returns an internal error

	const addr = "0xabc123def456789012345678901234567890abcd"
	nonceRepo := &testutil.FakeNonceRepository{
		Nonce: &domain.Nonce{WalletAddress: addr, Message: "test message"},
	}
	old := recoverAddress
	recoverAddress = func(_, _ string) (string, error) { return addr, nil }
	defer func() { recoverAddress = old }()

	walletRepo := &testutil.FakeUserWalletRepository{FindByRaylsAddressErr: errors.New("db error")}
	svc := newAuthSvcWithNonce(
		nonceRepo,
		testutil.NewFakeUserRepository(),
		walletRepo,
		&testutil.FakeRaylsAccessManagerClient{},
	)

	_, _, err := svc.VerifySIWE(context.Background(), addr, "0xsig", "nonce")

	var internalErr *core.InternalError
	require.Error(t, err)
	require.True(t, errors.As(err, &internalErr))
}

func TestAuthService_VerifySIWE_ExistingUserWithRoles_ReturnsUser(t *testing.T) {
	// A valid signature for a role_assigned user with on-chain roles returns the user

	const addr = "0xabc123def456789012345678901234567890abcd"
	existingUser := domain.User{Name: "Alice", IsActive: true, Status: domain.UserStatusRoleAssigned}
	existingUser.ID = uuid.New()

	nonceRepo := &testutil.FakeNonceRepository{
		Nonce: &domain.Nonce{WalletAddress: addr, Message: "test message"},
	}
	old := recoverAddress
	recoverAddress = func(_, _ string) (string, error) { return addr, nil }
	defer func() { recoverAddress = old }()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{existingUser}
	walletRepo := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{RaylsAddress: addr, UserID: existingUser.ID, IsActive: true}},
	}
	ramClient := &testutil.FakeRaylsAccessManagerClient{Roles: []string{"PRIVACY_NODE_OPERATOR"}}
	svc := newAuthSvcWithNonce(nonceRepo, userRepo, walletRepo, ramClient)

	user, roles, err := svc.VerifySIWE(context.Background(), addr, "0xsig", "nonce")

	require.NoError(t, err)
	assert.Equal(t, []string{"PRIVACY_NODE_OPERATOR"}, roles)
	assert.Equal(t, existingUser.ID, user.ID)
}

func TestAuthService_VerifySIWE_WaitingRoleAssignment_ReturnsPending(t *testing.T) {
	// A user with status=waiting_role_assignment returns RoleAssignmentPendingError

	const addr = "0xabc123def456789012345678901234567890abcd"
	waitingUser := domain.User{IsActive: true, Status: domain.UserStatusWaitingRoleAssignment}
	waitingUser.ID = uuid.New()

	nonceRepo := &testutil.FakeNonceRepository{
		Nonce: &domain.Nonce{WalletAddress: addr, Message: "test message"},
	}
	old := recoverAddress
	recoverAddress = func(_, _ string) (string, error) { return addr, nil }
	defer func() { recoverAddress = old }()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{waitingUser}
	walletRepo := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{RaylsAddress: addr, UserID: waitingUser.ID, IsActive: true}},
	}
	svc := newAuthSvcWithNonce(nonceRepo, userRepo, walletRepo, &testutil.FakeRaylsAccessManagerClient{})

	_, _, err := svc.VerifySIWE(context.Background(), addr, "0xsig", "nonce")

	var pending *core.RoleAssignmentPendingError
	require.Error(t, err)
	require.True(t, errors.As(err, &pending))
}

func TestAuthService_VerifySIWE_DeactivatedUser_ReturnsAccountSuspended(t *testing.T) {
	// A valid signature for a deactivated user returns AccountSuspendedError

	const addr = "0xabc123def456789012345678901234567890abcd"
	deactivatedUser := domain.User{IsActive: false, Status: domain.UserStatusRoleAssigned}
	deactivatedUser.ID = uuid.New()

	nonceRepo := &testutil.FakeNonceRepository{
		Nonce: &domain.Nonce{WalletAddress: addr, Message: "test message"},
	}
	old := recoverAddress
	recoverAddress = func(_, _ string) (string, error) { return addr, nil }
	defer func() { recoverAddress = old }()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{deactivatedUser}
	walletRepo := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{RaylsAddress: addr, UserID: deactivatedUser.ID, IsActive: true}},
	}
	svc := newAuthSvcWithNonce(nonceRepo, userRepo, walletRepo, &testutil.FakeRaylsAccessManagerClient{})

	_, _, err := svc.VerifySIWE(context.Background(), addr, "0xsig", "nonce")

	var suspended *core.AccountSuspendedError
	require.Error(t, err)
	require.True(t, errors.As(err, &suspended))
}

func TestAuthService_VerifySIWE_GetRolesFails_ReturnsServiceUnavailable(t *testing.T) {
	// When GetRoles RPC fails, ServiceUnavailableError is returned (not ROLE_ASSIGNMENT_PENDING)

	const addr = "0xabc123def456789012345678901234567890abcd"
	existingUser := domain.User{IsActive: true, Status: domain.UserStatusRoleAssigned}
	existingUser.ID = uuid.New()

	nonceRepo := &testutil.FakeNonceRepository{
		Nonce: &domain.Nonce{WalletAddress: addr, Message: "test message"},
	}
	old := recoverAddress
	recoverAddress = func(_, _ string) (string, error) { return addr, nil }
	defer func() { recoverAddress = old }()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{existingUser}
	walletRepo := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{RaylsAddress: addr, UserID: existingUser.ID, IsActive: true}},
	}
	ramClient := &testutil.FakeRaylsAccessManagerClient{Err: errors.New("rpc timeout")}
	svc := newAuthSvcWithNonce(nonceRepo, userRepo, walletRepo, ramClient)

	_, _, err := svc.VerifySIWE(context.Background(), addr, "0xsig", "nonce")

	var svcUnavail *core.ServiceUnavailableError
	require.Error(t, err)
	require.True(t, errors.As(err, &svcUnavail))
}

func TestAuthService_VerifySIWE_RoleAssignedButNoRoles_ReturnsPending(t *testing.T) {
	// status=role_assigned but GetRoles returns empty → RoleAssignmentPendingError

	const addr = "0xabc123def456789012345678901234567890abcd"
	existingUser := domain.User{IsActive: true, Status: domain.UserStatusRoleAssigned}
	existingUser.ID = uuid.New()

	nonceRepo := &testutil.FakeNonceRepository{
		Nonce: &domain.Nonce{WalletAddress: addr, Message: "test message"},
	}
	old := recoverAddress
	recoverAddress = func(_, _ string) (string, error) { return addr, nil }
	defer func() { recoverAddress = old }()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{existingUser}
	walletRepo := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{RaylsAddress: addr, UserID: existingUser.ID, IsActive: true}},
	}
	ramClient := &testutil.FakeRaylsAccessManagerClient{Roles: []string{}}
	svc := newAuthSvcWithNonce(nonceRepo, userRepo, walletRepo, ramClient)

	_, _, err := svc.VerifySIWE(context.Background(), addr, "0xsig", "nonce")

	var pending *core.RoleAssignmentPendingError
	require.Error(t, err)
	require.True(t, errors.As(err, &pending))
}
