package core

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

// ============================================================================
// SECONDARY PORTS (Infrastructure Needs)
// These are what the Core needs - dependencies injected from outside
// ============================================================================

// ErrRecordNotFound should be returned by repositories when a queried record doesn't exist
var ErrRecordNotFound = errors.New("record not found")

// ErrDuplicateEmail should be returned by repositories when an email uniqueness constraint is violated
var ErrDuplicateEmail = errors.New("email already registered")

// ErrDuplicateOAuthLink should be returned when an OAuth provider+ID pair already exists
var ErrDuplicateOAuthLink = errors.New("OAuth provider link already exists")

// ErrDuplicateWalletAddress should be returned when a rayls_address uniqueness constraint is violated
var ErrDuplicateWalletAddress = errors.New("wallet address already registered")

// ErrWalletNotFound is returned when a queried wallet address is not present in user_wallets.
var ErrWalletNotFound = errors.New("wallet not found")

// ErrTxReverted is returned by on-chain services when a transaction receipt reports
// status 0 (reverted). Handlers match it with errors.Is to return a 422 without
// depending on the (custody/go-ethereum) error message text.
var ErrTxReverted = errors.New("transaction reverted on-chain")

// UserRepository defines persistence operations for users
type UserRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Count(ctx context.Context) (int64, error)
	// SetOnChainUserID persists keccak256(user.ID) on the user row for reverse-lookup. Idempotent.
	SetOnChainUserID(ctx context.Context, userID uuid.UUID, onChainUserID []byte) error
	// FindByOnChainUserIDs reverse-maps on-chain keccak256 hashes to user UUIDs in one query,
	// keyed by the lowercase hex encoding of the hash. Unmatched hashes are absent from the map.
	FindByOnChainUserIDs(ctx context.Context, hashes [][]byte) (map[string]uuid.UUID, error)
}

// UserSignupDetailsRepository persists the profile answers collected by the standalone
// email sign-up form. One row per user; absent for users who signed in via OAuth or SIWE.
type UserSignupDetailsRepository interface {
	// Upsert writes the details for details.UserID, replacing any previous submission.
	Upsert(ctx context.Context, details *domain.UserSignupDetails) error
	// FindByUserID returns core.ErrRecordNotFound when the user never filled the form.
	FindByUserID(ctx context.Context, userID uuid.UUID) (*domain.UserSignupDetails, error)
}

// UserOAuthProviderRepository defines persistence operations for OAuth provider links
type UserOAuthProviderRepository interface {
	Create(ctx context.Context, provider *domain.UserOAuthProvider) error
	FindByProviderAndID(
		ctx context.Context,
		provider domain.OAuthProvider,
		oauthID string,
	) (*domain.UserOAuthProvider, error)
	FindByProviderAndUserID(
		ctx context.Context,
		provider domain.OAuthProvider,
		userID uuid.UUID,
	) (*domain.UserOAuthProvider, error)
	UpdateOAuthID(ctx context.Context, id uuid.UUID, oauthID string) error
	CountByUserID(ctx context.Context, userID uuid.UUID) (int64, error)
}

// UserWalletRepository defines persistence operations for user blockchain wallets
type UserWalletRepository interface {
	Create(ctx context.Context, wallet *domain.UserWallet) error
	// FindByUserID returns the user's identity/login wallet (earliest active wallet, provider-agnostic).
	FindByUserID(ctx context.Context, userID uuid.UUID) (*domain.UserWallet, error)
	// GetSignerWalletForChain returns the wallet used for on-chain signing on the given chain:
	// the earliest active HSM-custodied wallet. Custody can only sign HSM wallets.
	GetSignerWalletForChain(ctx context.Context, userID uuid.UUID, chain domain.WalletChain) (*domain.UserWallet, error)
	// GetSignerWalletByAddress returns the user's active HSM-custodied wallet whose address matches
	// (case-insensitive), or ErrRecordNotFound. Custody can only sign HSM wallets; chain is the
	// caller's to validate. Used to confirm a caller-supplied signer is one of their own signer wallets.
	GetSignerWalletByAddress(ctx context.Context, userID uuid.UUID, address string) (*domain.UserWallet, error)
	FindByRaylsAddress(ctx context.Context, address string) (*domain.UserWallet, error)
	// CompletePending fills in the real address/external id of a mint-intent row and activates
	// it. Returns ErrDuplicateWalletAddress if the address is already held by another row.
	CompletePending(ctx context.Context, id uuid.UUID, address, externalID string) error
	// DeletePending removes an unfulfilled mint-intent row. Used when the mint itself failed,
	// so no HSM key was ever created and the intent is meaningless.
	DeletePending(ctx context.Context, id uuid.UUID) error
	// FindPendingByUserID returns the user's unfulfilled mint intents, oldest first. A row here
	// means a previous attempt died between writing the intent and persisting the address — the
	// HSM may hold a key that nothing references.
	FindPendingByUserID(ctx context.Context, userID uuid.UUID) ([]domain.UserWallet, error)
}

// TokenBlacklistRepository defines persistence operations for blacklisted refresh tokens
type TokenBlacklistRepository interface {
	Create(ctx context.Context, entry *domain.TokenBlacklist) error
	ExistsByJTI(ctx context.Context, jti string) (bool, error)
	DeleteExpired(ctx context.Context) error
}

// Transactor executes a function within a single database transaction.
// If fn returns an error the transaction is rolled back; otherwise it is committed.
type Transactor interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

// NonceRepository defines persistence operations for SIWE nonces
type NonceRepository interface {
	Create(ctx context.Context, nonce *domain.Nonce) error
	FindValidAndMarkUsed(ctx context.Context, addr, nonce string) (*domain.Nonce, error)
	DeleteExpired(ctx context.Context) error
}

// RaylsAccessManagerClient checks on-chain role assignments at login time.
type RaylsAccessManagerClient interface {
	// GetRoles returns the role labels assigned to the given wallet address.
	// Returns an empty slice (not an error) when the wallet has no roles.
	GetRoles(ctx context.Context, walletAddress string) ([]string, error)
}

// BootstrapService provisions the initial system administrator.
type BootstrapService interface {
	// Bootstrap creates the admin user with an HSM wallet. Returns the wallet address.
	// Returns BootstrapAlreadyCompletedError if any user already exists.
	Bootstrap(ctx context.Context, email string) (address string, err error)
}

// CustodyService creates and manages blockchain wallets via an external custody provider.
type CustodyService interface {
	// CreateWallet generates a new wallet for the given user and returns the address
	// and a provider-specific external ID for future reference.
	CreateWallet(ctx context.Context, userID uuid.UUID) (address, externalID string, err error)
	// SignAndTransact signs the RLP-encoded unsigned transaction with the key stored under
	// signerAddress and broadcasts it to the chain identified by chainID.
	// Returns the transaction hash on success.
	SignAndTransact(ctx context.Context, payload []byte, signerAddress, chainID string) (string, error)
}

// RoleService queries on-chain role assignments for a wallet address.
type RoleService interface {
	// GetUserRole returns the role labels for the given wallet address (e.g. "BANK_EMPLOYEE").
	// Returns an empty string when the wallet has no roles. userID may be ignored by implementations
	// that only need the wallet address (e.g. RaylsAccessManager).
	GetUserRole(ctx context.Context, userID uuid.UUID, walletAddress string) (string, error)
}

// OperatorSignerResolver resolves, at request time, the HSM operator wallet address that signs
// operator-authority governance writes. There is no configured operator address or env var; the
// operator is the first eligible member holding both the PRIVACY_NODE_OPERATOR and BANK_EMPLOYEE
// roles found in the database — the two roles that together gate every governance write.
type OperatorSignerResolver interface {
	// Resolve returns the RaylsAddress of the first active member holding both PRIVACY_NODE_OPERATOR
	// and BANK_EMPLOYEE whose UserWallet is HSM-custodied. Returns *NoOperatorSignerError when either
	// role is missing or no member holding both has an eligible HSM wallet.
	Resolve(ctx context.Context) (operatorAddress string, err error)
}

// OnChainAddressPair is the read model for a user's public/private wallet pair in RNUserGovernance.
type OnChainAddressPair struct {
	PublicChainAddress  string
	PrivateChainAddress string
	Status              domain.ApprovalStatus
	CreatedAt           time.Time
}

// OnChainPendingGroup is the admin-discovery read model: every pending pair for one on-chain user,
// keyed by the raw keccak256 hash (not a UUID), as returned by getAllPendingAddressPairs.
type OnChainPendingGroup struct {
	OnChainUserID [32]byte
	AddressPairs  []OnChainAddressPair
}

// PendingUserAddressPairs is the admin-discovery read model after the on-chain keccak256 hash has
// been reverse-mapped to an ops-api user UUID.
type PendingUserAddressPairs struct {
	UserID       uuid.UUID
	AddressPairs []OnChainAddressPair
}

// UserGovernanceService manages on-chain user identity and address pairs in RNUserGovernance.
// All writes are operator-authority writes signed by the resolved operator wallet.
type UserGovernanceService interface {
	// EnsureUser idempotently creates the on-chain user: it checks HasUser and only calls
	// CreateUser when the user does not already exist on-chain.
	EnsureUser(ctx context.Context, operatorAddress string, onChainUserID [32]byte) error
	AddAddressPair(
		ctx context.Context,
		operatorAddress string,
		onChainUserID [32]byte,
		publicAddr, privateAddr string,
	) (txHash string, err error)
	SetApprovalStatus(
		ctx context.Context,
		operatorAddress string,
		onChainUserID [32]byte,
		publicAddr, privateAddr string,
		status domain.ApprovalStatus,
	) (txHash string, err error)
	ListPending(ctx context.Context, onChainUserID [32]byte) ([]OnChainAddressPair, error)
	ListApproved(ctx context.Context, onChainUserID [32]byte) ([]OnChainAddressPair, error)
	// ListAllPending backs the admin discovery endpoint: every pending pair across all on-chain
	// users, grouped by the raw keccak256 hash (not a UUID).
	ListAllPending(ctx context.Context) ([]OnChainPendingGroup, error)
}

// OnboardingService orchestrates self-service onboarding: it creates a fresh HSM wallet pair,
// persists both wallets, and registers the pair on-chain via operator-authority governance writes.
type OnboardingService interface {
	// AddAddressPair creates a new HSM wallet pair for the user and registers it in RNUserGovernance,
	// returning the created pair in pending status. Each call creates a fresh pair.
	AddAddressPair(ctx context.Context, userID uuid.UUID) (*OnChainAddressPair, error)
	// ListMine returns the caller's on-chain address pairs. A nil status returns all pairs
	// (pending ∪ approved); otherwise it filters to that approval status.
	ListMine(ctx context.Context, userID uuid.UUID, status *domain.ApprovalStatus) ([]OnChainAddressPair, error)
	// ListAllPending returns every pending pair across all users, each owner resolved from its
	// on-chain keccak256 hash to an ops-api UUID. Hashes that do not resolve are skipped and logged.
	ListAllPending(ctx context.Context) ([]PendingUserAddressPairs, error)
	// SetApprovalStatus approves or rejects a pair for the user identified by userID (never the body).
	SetApprovalStatus(
		ctx context.Context,
		userID uuid.UUID,
		publicAddr, privateAddr string,
		status domain.ApprovalStatus,
	) error
}

// OnChainService interacts with on-chain contracts on behalf of the OSA wallet.
type OnChainService interface {
	// RegisterUser grants the configured role to walletAddress via RaylsAccessManager.grantRole,
	// waits for the transaction receipt, and returns an error if the receipt status is not successful.
	RegisterUser(ctx context.Context, userID uuid.UUID, walletAddress, operatorAddress string) error
}

// TokenDeploySpec describes a single protocol token to deploy via RNContractFactory.
// Which fields are required depends on ErcStandard: ERC20/Enygma use Name/Symbol/Decimals,
// ERC721 variants use URI/Name/Symbol, and ERC1155 variants use URI/Name.
type TokenDeploySpec struct {
	ErcStandard domain.ErcStandard
	Name        string
	Symbol      string
	URI         string
	Decimals    uint8
}

// TokenFunction is a single AM-gated contract function the caller is allowed to invoke.
type TokenFunction struct {
	Selector string `json:"selector"`       // 4-byte selector, e.g. "0x40c10f19"
	Name     string `json:"name,omitempty"` // friendly name (mint/burn/...), empty if unknown
}

// TokenPermissions describes what a given wallet can do on a token contract, derived from the
// Access Manager (role membership ∩ function permissions).
type TokenPermissions struct {
	ContractAddress string `json:"contractAddress"`
	WalletAddress   string `json:"walletAddress"`
	// IsPaused is the ACCESS MANAGER's pause on this managed contract (mirrored from its
	// ContractPauseUpdated events) — NOT the stablecoin's own `paused` state variable.
	IsPaused             bool `json:"isPaused"`
	CanMint              bool `json:"canMint"`
	CanBurn              bool `json:"canBurn"`
	CanSubmitTokenUpdate bool `json:"canSubmitTokenUpdate"`
	// CanPause reports whether this wallet is the stablecoin's `pauser`. Deliberately not
	// derived from AM roles: pause()/unpause() are gated on a msg.sender equality check
	// against the contract's own `pauser` address. False for non-stablecoins, which have no
	// pause function at all.
	CanPause bool `json:"canPause"`
	// IsTokenPaused is the stablecoin's OWN `paused` flag, read live from the contract — the
	// one that actually halts transfers. Distinct from IsPaused above; false for standards
	// with no pause function.
	IsTokenPaused bool            `json:"isTokenPaused"`
	Functions     []TokenFunction `json:"functions"`
}

// TokenPermissionService computes the functions a wallet may call on a token contract.
type TokenPermissionService interface {
	GetTokenPermissions(ctx context.Context, contractAddress, walletAddress string) (*TokenPermissions, error)
}

// MintInput holds the parsed parameters for a mint call. Which fields are used depends on the
// token standard: fungible (ERC20/Enygma) uses To+Amount; ERC721 uses To+TokenID; ERC1155 uses
// To+TokenID+Amount(+Data). Amount is already scaled to base units.
type MintInput struct {
	To      string
	Amount  *big.Int
	TokenID *big.Int
	Data    []byte
}

// BurnInput holds the parsed parameters for a burn call. Fungible (ERC20/Enygma) uses From+Amount;
// ERC721 uses TokenID; ERC1155 uses From+TokenID+Amount. Amount is already scaled to base units.
type BurnInput struct {
	From    string
	Amount  *big.Int
	TokenID *big.Int
}

// TeleportInput holds the parsed parameters for a teleportToPublicChain call. Which fields are used
// depends on the token standard: ERC20 uses To+Amount; ERC721 uses To+TokenID; ERC1155 uses
// To+TokenID+Amount(+Data). From is the origin on the private chain (the caller's signer wallet,
// whose position the preflight validates). Amount is in base units. To is the destination on the
// public chain.
type TeleportInput struct {
	From    string
	To      string
	Amount  *big.Int
	TokenID *big.Int
	Data    []byte
}

// TokenActionService signs and broadcasts mint/burn transactions to a token contract using the
// caller's custody wallet, dispatching the correct function per token standard.
type TokenActionService interface {
	Mint(
		ctx context.Context,
		signerAddress, tokenAddress string,
		standard domain.ErcStandard,
		in MintInput,
	) (txHash string, err error)
	Burn(
		ctx context.Context,
		signerAddress, tokenAddress string,
		standard domain.ErcStandard,
		in BurnInput,
	) (txHash string, err error)
	// SetPaused calls pause()/unpause() on a stablecoin. Unlike mint/burn these are gated on
	// the contract's own `pauser` address rather than an AccessManager role, so callers must
	// authorize with Pauser() instead of TokenPermissions.
	SetPaused(
		ctx context.Context,
		signerAddress, tokenAddress string,
		paused bool,
	) (txHash string, err error)
	// Pauser reads the contract's `pauser` address — the only account pause()/unpause() accept.
	Pauser(ctx context.Context, tokenAddress string) (address string, err error)
	// IsPaused reads the contract's own `paused` flag (distinct from the AccessManager-level
	// pause mirrored in am_managed_contracts).
	IsPaused(ctx context.Context, tokenAddress string) (paused bool, err error)
}

// TokenChainClient is the on-chain token client backing teleport: it reads a wallet's position and
// submits the teleportToPublicChain transaction. Pure on-chain — it applies no business rules
// (standard eligibility, registration, sufficiency); those live in TeleportService.
type TokenChainClient interface {
	ERC20Balance(ctx context.Context, tokenAddress, account string) (*big.Int, error)
	ERC721Owner(ctx context.Context, tokenAddress string, tokenID *big.Int) (ownerAddress string, err error)
	ERC1155Balance(ctx context.Context, tokenAddress, account string, tokenID *big.Int) (*big.Int, error)
	TeleportERC20(
		ctx context.Context,
		signerAddress, tokenAddress, to string,
		amount, destinationChainID *big.Int,
	) (txHash string, err error)
	TeleportERC721(
		ctx context.Context,
		signerAddress, tokenAddress, to string,
		tokenID, destinationChainID *big.Int,
	) (txHash string, err error)
	TeleportERC1155(
		ctx context.Context,
		signerAddress, tokenAddress, to string,
		tokenID, amount, destinationChainID *big.Int,
		data []byte,
	) (txHash string, err error)
}

// TeleportService moves an asset from the privacy chain to the public chain. It owns the business
// rules: only ERC20/ERC721/ERC1155 are eligible, and a mandatory preflight (token registered &
// active, caller balance/ownership) runs before the teleportToPublicChain transaction is signed.
// On-chain interaction is delegated to a TokenChainClient.
type TeleportService interface {
	Teleport(
		ctx context.Context,
		tokenAddress string,
		standard domain.ErcStandard,
		in TeleportInput,
	) (txHash string, err error)
}

// TokenDeployService deploys a protocol token through the RNContractFactory, signing the
// transaction with the given custody wallet (signerAddress).
type TokenDeployService interface {
	// Deploy builds the factory deploy calldata for spec, signs it via the custody wallet at
	// signerAddress, waits for the receipt, and returns the deployed token address and tx hash.
	Deploy(ctx context.Context, signerAddress string, spec TokenDeploySpec) (deployedAddr, txHash string, err error)
	// EstimateDeploy returns the real on-chain gas estimate for deploying spec from
	// signerAddress, without executing it (eth_estimateGas + eth_gasPrice).
	EstimateDeploy(ctx context.Context, signerAddress string, spec TokenDeploySpec) (TokenDeployEstimate, error)
	// ChainID returns the chain ID the factory is deployed on, used to stamp the token's issuer.
	ChainID() string
}

type RegisteredToken struct {
	Standard     domain.ErcStandard
	Name         string
	Symbol       string
	URI          string
	TokenAddress string
	Status       domain.PrivacyNodeStatus
	LastUpdated  time.Time
}

// RegisterTokenInput holds the parameters for registering an already-deployed token contract into
// the on-chain TokenRegistry. Registration is address-only: the contract reads the token's
// name/symbol/standard/supply on-chain, so TokenAddress (the private-chain contract address) is all
// the adapter needs.
type RegisterTokenInput struct {
	TokenAddress string
}

// TokenRegistryAdapter is the pure on-chain TokenRegistry (PNTokenCoreV1) catalog adapter. Writes are
// operator-authority writes signed via custody by the operator wallet passed in (operatorAddress);
// it does not resolve the operator itself.
type TokenRegistryAdapter interface {
	Register(ctx context.Context, operatorAddress string, in RegisterTokenInput) (txHash string, err error)
	SetStatus(
		ctx context.Context,
		operatorAddress, tokenAddress string,
		status domain.PrivacyNodeStatus,
	) (txHash string, err error)
	// Freeze/Unfreeze freeze or unfreeze a registered token at the given layer via the dedicated
	// TokenFreezeManager contract methods (freezeOnPrivacyNode / freezeOnPublicChain and their
	// unfreeze counterparts), rather than routing FROZEN through SetStatus.
	Freeze(
		ctx context.Context,
		operatorAddress, tokenAddress string,
		layer domain.FreezeLayer,
	) (txHash string, err error)
	Unfreeze(
		ctx context.Context,
		operatorAddress, tokenAddress string,
		layer domain.FreezeLayer,
	) (txHash string, err error)
	// Submit submits an AUTHORIZED token to the given target layer (submitToHub / submitToPublicChain).
	// It only initiates the flow; activation/deployment complete via cross-chain PNH / relayer callbacks.
	Submit(
		ctx context.Context,
		operatorAddress, tokenAddress string,
		target domain.SubmitTarget,
	) (txHash string, err error)

	List(ctx context.Context) ([]RegisteredToken, error)
	ListByStatus(ctx context.Context, status domain.PrivacyNodeStatus) ([]RegisteredToken, error)
	GetByAddress(ctx context.Context, tokenAddress string) (*RegisteredToken, error)
	GetBySymbol(ctx context.Context, symbol string) (*RegisteredToken, error)
	Exists(ctx context.Context, tokenAddress string) (bool, error)
}

// TokenRegistryService orchestrates the token-registry capability: it resolves the operator wallet
// and drives the TokenRegistryAdapter. Register reads the entry back after the write and returns the
// resulting RegisteredToken (initial status WAITING_APPROVAL).
type TokenRegistryService interface {
	Register(ctx context.Context, in RegisterTokenInput) (*RegisteredToken, error)
	SetStatus(ctx context.Context, tokenAddress string, status domain.PrivacyNodeStatus) (txHash string, err error)
	// Freeze/Unfreeze resolve the operator wallet and submit an operator-signed freeze/unfreeze at
	// the given layer.
	Freeze(ctx context.Context, tokenAddress string, layer domain.FreezeLayer) (txHash string, err error)
	Unfreeze(ctx context.Context, tokenAddress string, layer domain.FreezeLayer) (txHash string, err error)
	// Submit resolves the operator wallet and submits an operator-signed submitToHub /
	// submitToPublicChain for the given target.
	Submit(ctx context.Context, tokenAddress string, target domain.SubmitTarget) (txHash string, err error)

	List(ctx context.Context) ([]RegisteredToken, error)
	ListByStatus(ctx context.Context, status domain.PrivacyNodeStatus) ([]RegisteredToken, error)
	GetByAddress(ctx context.Context, tokenAddress string) (*RegisteredToken, error)
	GetBySymbol(ctx context.Context, symbol string) (*RegisteredToken, error)
	Exists(ctx context.Context, tokenAddress string) (bool, error)
}

// ProvisioningService handles post-login setup for new users: custody wallet
// creation and on-chain role assignment.
// TokenDeployEstimate is the real gas cost of a token deploy. All numeric fields are
// decimal strings so they survive JSON without precision loss (wei values exceed 2^53).
type TokenDeployEstimate struct {
	GasLimit    uint64 `json:"gasLimit"`
	GasPriceWei string `json:"gasPriceWei"`
	TotalFeeWei string `json:"totalFeeWei"` // gasLimit * gasPriceWei
}

// ProvisioningService handles post-login setup for new users: custody wallet creation.
// On-chain role assignment is performed out of band by the deploy/ops tooling.
type ProvisioningService interface {
	// Provision is idempotent — it is safe to call on every login.
	// It returns immediately if the user is already fully provisioned.
	Provision(ctx context.Context, user *domain.User) error
}

// WalletFunder tops up a freshly-created custody wallet so it can pay gas on a
// non-gasless chain (e.g. the public-chain dev testnet). Implementations are DEV-ONLY
// and wired only when a faucet key is configured; on gasless Privacy Nodes it is nil.
type WalletFunder interface {
	// FundWallet ensures address holds at least the configured target balance, transferring
	// only the shortfall. It is best-effort: callers log and continue on error rather than
	// failing wallet creation.
	FundWallet(ctx context.Context, address string) error
}

// RoleGranter grants a freshly-provisioned custody wallet the on-chain roles it needs
// (FACTORY_DEPLOYER to deploy tokens via RNContractFactory, PRIVACY_NODE_OPERATOR for the
// login role check) — the same roles a user would otherwise receive out of band. Each grant
// is a permissioned write signed by a privileged grantor EOA that holds admin authority over
// the role. Wired only when a grantor key is configured; nil otherwise (in which case role
// assignment stays fully out of band, as before).
type RoleGranter interface {
	// GrantRoles ensures address holds each named role on-chain. It is idempotent — a no-op
	// per role the account already holds. Best-effort at the call site: a failure is logged
	// and does not fail wallet creation / login.
	GrantRoles(ctx context.Context, address string, roles []string) error
}

// AuthService defines authentication business logic
type AuthService interface {
	GenerateChallenge(ctx context.Context, walletAddress string) (message, nonce string, err error)
	// VerifySIWE validates the SIWE signature and applies the v1.4 login decision tree.
	// Returns (user, roles, nil) only when the user is fully eligible to receive a JWT.
	// Returns a domain error (WalletRegisteredError, RoleAssignmentPendingError, AccountSuspendedError,
	// ServiceUnavailableError, UnauthorizedError) otherwise.
	VerifySIWE(
		ctx context.Context,
		walletAddress, signature, nonce string,
	) (user *domain.User, roles []string, err error)
	// FindOrCreateOAuthUser looks up an OAuth user and applies the v1.4 login decision tree.
	// emailVerified must be true for the email-based bootstrap fallback to activate.
	// Returns (user, roles, nil) only when the user is eligible to receive a JWT.
	FindOrCreateOAuthUser(
		ctx context.Context,
		provider domain.OAuthProvider,
		oauthID, name, email string,
		emailVerified bool,
	) (*domain.User, []string, error)
}

// ============================================================================
// INDEXER PORTS
// ============================================================================

// EventPublisher publishes a typed event to a NATS JetStream subject.
// Implemented by cmd/api/messaging/publisher.go.
type EventPublisher interface {
	Publish(ctx context.Context, subject string, msg any) error
}

// LiveEventPublisher publishes an ephemeral (non-persisted) NATS message for live fan-out,
// such as pushing token updates to SSE clients. Delivery is at-most-once.
type LiveEventPublisher interface {
	PublishLive(subject string, data []byte) error
}

// TokenFilter holds optional filters and pagination for TokenRepository.List.
type TokenFilter struct {
	TokenClass string
	Status     *uint8
	TokenType  string // ERC standard string e.g. "ERC-20", "ERC-721", "ERC-1155"
	Name       string // partial, case-insensitive
	Symbol     string // partial, case-insensitive
	IssuerID   string // exact match on the chain the token was deployed on (chainId as string)
	Page       int
	Limit      int
}

// TokenRepository persists and queries tokens discovered by the indexer.
type TokenRepository interface {
	Upsert(ctx context.Context, token *domain.Token) error
	UpdateSupplyAndHolders(ctx context.Context, address, totalSupply string, holderCount int) error
	FindByAddress(ctx context.Context, address string) (*domain.Token, error)
	List(ctx context.Context, filter TokenFilter) ([]*domain.Token, int64, error)
}

// TokenEventRepository records immutable on-chain status change events.
type TokenEventRepository interface {
	Insert(ctx context.Context, event *domain.TokenEvent) error
	ListByAddress(ctx context.Context, address string, page, limit int) ([]*domain.TokenEvent, int64, error)
}

// WalletBalanceRepository persists per-wallet, per-token balances synced from Blockscout.
// Upsert is keyed on (wallet_address, token_address) and MUST refuse to overwrite a stored
// row whose block_number is greater than the incoming one — older updates cannot regress
// a fresher state.
type WalletBalanceRepository interface {
	Upsert(ctx context.Context, balance *domain.WalletBalance) error
	ListByWallet(ctx context.Context, walletAddress string) ([]*domain.WalletBalance, error)
	GetByWalletAndToken(ctx context.Context, walletAddress, tokenAddress string) (*domain.WalletBalance, error)
}

// Cursor keys used with IndexerStateRepository.
const (
	CursorBlockscoutInsertedAt   = "blockscout_cursor_inserted_at"
	CursorBlockscoutAddress      = "blockscout_cursor_address"
	CursorBlockscoutBalancesID   = "blockscout_balances_cursor_id"
	CursorPNLastEventBlock       = "pn_last_event_block"
	CursorAccessManagerLastBlock = "access_manager_last_block"
)

// IndexerStateRepository is a key/value store for listener cursors.
// A missing key must return ErrRecordNotFound — callers treat this as zero/epoch.
type IndexerStateRepository interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string) error
}

// MessageRepository persists Hub contract events (Hub mode only).
type MessageRepository interface {
	Insert(ctx context.Context, msg *domain.Message) error
	List(ctx context.Context, page, limit int) ([]*domain.Message, int64, error)
}

// ============================================================================
// ACCESS MANAGER REPOSITORIES
// ============================================================================

type AccessManagerRoleRepository interface {
	Upsert(ctx context.Context, role *domain.AccessManagerRole) error
	FindByID(ctx context.Context, roleID uint64) (*domain.AccessManagerRole, error)
	List(ctx context.Context) ([]*domain.AccessManagerRole, error)
}

type AccessManagerRoleMemberRepository interface {
	Upsert(ctx context.Context, member *domain.AccessManagerRoleMember) error
	Revoke(ctx context.Context, roleID uint64, account string) error
	ListByRole(ctx context.Context, roleID uint64, activeOnly bool) ([]*domain.AccessManagerRoleMember, error)
	ListByAccount(ctx context.Context, account string) ([]*domain.AccessManagerRoleMember, error)
	// FindActiveAccountWithAllRoles returns one active-member account that holds EVERY role in roleIDs
	// (set intersection), the lowest-ordered account when several qualify. Returns ErrRecordNotFound
	// when no active account holds all of them (including when roleIDs is empty).
	FindActiveAccountWithAllRoles(ctx context.Context, roleIDs []uint64) (account string, err error)
}

type AccessManagerManagedContractRepository interface {
	Upsert(ctx context.Context, c *domain.AccessManagerManagedContract) error
	SetPaused(ctx context.Context, contractAddress string, paused bool) error
	List(ctx context.Context) ([]*domain.AccessManagerManagedContract, error)
	FindByAddress(ctx context.Context, contractAddress string) (*domain.AccessManagerManagedContract, error)
}

type AccessManagerFunctionPermissionRepository interface {
	Upsert(ctx context.Context, perm *domain.AccessManagerFunctionPermission) error
	Remove(ctx context.Context, contractAddress, selector string, roleID uint64) error
	ListByContract(ctx context.Context, contractAddress string) ([]*domain.AccessManagerFunctionPermission, error)
}

type AccessManagerScheduledOperationRepository interface {
	Upsert(ctx context.Context, op *domain.AccessManagerScheduledOperation) error
	UpdateStatus(ctx context.Context, operationID, status string) error
	List(ctx context.Context, status string) ([]*domain.AccessManagerScheduledOperation, error)
}

type AccessManagerContractScopedRoleMemberRepository interface {
	Upsert(ctx context.Context, m *domain.AccessManagerContractScopedRoleMember) error
	Revoke(ctx context.Context, roleID uint64, account, contractAddress string) error
	ListByAccount(ctx context.Context, account string) ([]*domain.AccessManagerContractScopedRoleMember, error)
}

type AccessManagerEventLogRepository interface {
	Insert(ctx context.Context, entry *domain.AccessManagerEventLog) error
	List(ctx context.Context, eventName string, page, limit int) ([]*domain.AccessManagerEventLog, int64, error)
}
