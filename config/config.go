package config

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Database struct {
	ConnectionString string `mapstructure:"DATABASE_CONNECTIONSTRING"`
	// IdentityConnectionString points at the SHARED identity database (ops_identity), which
	// owns users and their custody wallets. A custody wallet is an EVM keypair and works on
	// every chain, so there is exactly one per user — the per-chain ops-api reads it from
	// here rather than minting its own. Roles stay per-chain (resolved from am_*).
	//
	// Empty = single-service deployment: users live in this instance's own database, as
	// before. Set it and the account tables in this database are ignored entirely.
	IdentityConnectionString string `mapstructure:"IDENTITY_DB_CONN"`
}

type Auth struct {
	JWTSecret             string `mapstructure:"JWT_SECRET" validate:"required"`
	BaseURL               string `mapstructure:"BASE_URL" validate:"required"`
	SecureCookies         bool   `mapstructure:"SECURE_COOKIES"`
	GoogleClientID        string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret    string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	MicrosoftClientID     string `mapstructure:"MICROSOFT_CLIENT_ID"`
	MicrosoftClientSecret string `mapstructure:"MICROSOFT_CLIENT_SECRET"`
	// PostLoginRedirectURL is the frontend URL to redirect to after a successful OAuth login.
	// If empty, the callback returns JSON instead of redirecting.
	PostLoginRedirectURL string `mapstructure:"POST_LOGIN_REDIRECT_URL"`
	// OAuthRedirectBase is a FIXED host used as the OAuth redirect_uri base, so a
	// single redirect URI can be registered for many per-instance hosts. When set,
	// redirect_uri = OAUTH_REDIRECT_BASE + "/auth/<provider>/callback" and the
	// originating instance callback is carried in the OAuth state for a relay to
	// bounce back. Empty = use BASE_URL directly (single-host behaviour).
	OAuthRedirectBase string `mapstructure:"OAUTH_REDIRECT_BASE"`
	// CookieDomain sets the session cookie's Domain so it is shared across
	// subdomains (e.g. ".rayup.dappsforall.com"), letting the per-chain
	// playground (play-<slug>) forward the ops session to this api. Empty = host-only.
	CookieDomain string `mapstructure:"COOKIE_DOMAIN"`
	// AdminEmail identifies the single admin user by email. A login (Google, Microsoft,
	// or standalone email) whose email matches gets an `is_admin` claim in the JWT,
	// gating admin-only actions. Empty = no admin (every user is a regular user).
	AdminEmail string `mapstructure:"OPS_ADMIN_EMAIL"`
	// BootstrapToken protects POST /admin/bootstrap, which creates the FIRST admin. The
	// endpoint is otherwise unauthenticated and guarded only by "no users exist yet", so
	// on a fresh deployment whoever calls it first becomes the administrator. When set,
	// callers must present it as `Authorization: Bearer <token>`. Empty keeps the old
	// unauthenticated behaviour, which is only safe when the endpoint is unreachable
	// from outside (dev, or an infra-level network policy).
	BootstrapToken string `mapstructure:"BOOTSTRAP_TOKEN"`
	// EmailSignupEnabled turns on POST /auth/signup, which creates an account and issues
	// a session from an email address ALONE — there is no verification code yet, so it
	// trusts whatever address the caller submits. That means anyone who can reach the
	// endpoint can log in as anyone, including the admin, so it is OFF by default and
	// must stay off outside local development. Remove this flag once signup requires a
	// verified code (see EmailSignup).
	EmailSignupEnabled bool `mapstructure:"EMAIL_SIGNUP_ENABLED"`
}

// validateDeployed rejects auth settings that are safe on a laptop but not in a deployed
// environment. SecureCookies is the discriminator: it is only set where the service is
// served over HTTPS, which is exactly where these defaults stop being acceptable.
//
// These are refusals rather than warnings because both failures are silent and total. A
// warning is emitted once into a log nobody reads during a rollout, while the service comes
// up and serves an authentication bypass indefinitely. A pod that will not start is noticed
// immediately, and the operator is one env var away from a correct deployment.
func (a Auth) validateDeployed() error {
	if !a.SecureCookies {
		return nil
	}

	if a.EmailSignupEnabled {
		return fmt.Errorf(
			"EMAIL_SIGNUP_ENABLED must be false when SECURE_COOKIES=true: /auth/signup issues " +
				"a session from an UNVERIFIED email address, so anyone who can reach it can log " +
				"in as anyone, including the admin, on every chain sharing this identity database",
		)
	}

	if a.BootstrapToken == "" {
		return fmt.Errorf(
			"BOOTSTRAP_TOKEN is required when SECURE_COOKIES=true: POST /admin/bootstrap creates " +
				"the FIRST admin and is otherwise guarded only by \"no users exist yet\", so on a " +
				"fresh database the first caller becomes the administrator",
		)
	}

	return nil
}

// PrivacyNode holds configuration for connecting to and polling the Privacy Node.
type PrivacyNode struct {
	RPCURL            string `mapstructure:"PN_RPC_URL"`
	TokenRegistryAddr string `mapstructure:"TOKEN_REGISTRY_ADDRESS"`
	BlockRange        int    `mapstructure:"PN_BLOCK_RANGE"`
	PollInterval      string `mapstructure:"PN_POLL_INTERVAL"`
}

// Blockscout holds configuration for the direct Blockscout PostgreSQL connection
// used for LISTEN/NOTIFY and the initial token backfill query.
type Blockscout struct {
	DBConnString  string `mapstructure:"BLOCKSCOUT_DB_CONN"`
	BackfillBatch int    `mapstructure:"BLOCKSCOUT_BACKFILL_BATCH"`
	// ChainID is the chain this instance indexes. Stamped as issuerId on tokens the
	// indexer discovers (each instance indexes exactly one chain), so the token list can
	// be scoped per-chain. Empty leaves discovered tokens without an issuer.
	ChainID string `mapstructure:"BLOCKCHAIN_CHAIN_ID"`
}

// Hub holds configuration for Hub-mode operation.
// The Hub uses PrivacyNode.RPCURL for RPC — no separate RPC URL is needed.
type Hub struct {
	ContractAddress string `mapstructure:"HUB_CONTRACT_ADDRESS"`
	DecryptionKey   string `mapstructure:"HUB_DECRYPTION_KEY"`
}

// NATS holds configuration for the JetStream messaging layer.
type NATS struct {
	URL string `mapstructure:"NATS_URL"`
	// Optional mutual TLS. When CAFile/CertFile/KeyFile are all set, the client
	// connects over TLS with the provided cert as its identity. Leaving them
	// empty preserves the previous plain-text connect behaviour.
	TLSCAFile   string `mapstructure:"NATS_TLS_CA_FILE"`
	TLSCertFile string `mapstructure:"NATS_TLS_CERT_FILE"`
	TLSKeyFile  string `mapstructure:"NATS_TLS_KEY_FILE"`
}

type Custody struct {
	// RaylsHSMURL is the base URL of the Rayls HSM custody service.
	// Required when CUSTODY_PROVIDER=rayls_hsm.
	RaylsHSMURL    string `mapstructure:"CUSTODY_RAYLS_HSM_URL"`
	RaylsHSMAPIKey string `mapstructure:"CUSTODY_RAYLS_HSM_API_KEY"`
	// TODO: move to a secrets manager (vault, AWS Secrets Manager, etc.)
	RaylsHSMPassword string `mapstructure:"CUSTODY_RAYLS_HSM_PASSWORD"`
	// RPCURL is the chain's JSON-RPC endpoint as CUSTODY reaches it, sent with every
	// signing request. Usually identical to BLOCKCHAIN_RPC_URL — but not when custody runs
	// outside this ops-api's network: a SHARED custody (one HSM holding every user's single
	// wallet) is not in the cluster, so the in-cluster Service DNS name we use ourselves does
	// not resolve for it and signing fails with "Name or service not known". Set this to the
	// host-reachable form there. Empty = fall back to BLOCKCHAIN_RPC_URL.
	RPCURL string `mapstructure:"CUSTODY_RPC_URL"`
}

type Blockchain struct {
	// RPCURL is the JSON-RPC endpoint of the private hub (Besu).
	RPCURL string `mapstructure:"BLOCKCHAIN_RPC_URL"`
	// DeploymentProxyRegistryAddr is the deployed address of DeploymentProxyRegistry,
	// used to resolve RaylsAccessManager (and RNContractFactory) at startup.
	DeploymentProxyRegistryAddr string `mapstructure:"DEPLOYMENT_PROXY_REGISTRY_ADDR"`
	// StartingBlock is the block number to begin indexing from when no cursor exists in the DB.
	// Defaults to 0. Set to the AccessManager deployment block to skip pre-deployment history.
	StartingBlock uint64 `mapstructure:"BLOCKCHAIN_STARTING_BLOCK"`
	// BlockBatchSize is the maximum number of blocks fetched per polling tick.
	// Defaults to 100.
	BlockBatchSize int64 `mapstructure:"BLOCKCHAIN_BLOCK_BATCH_SIZE"`
	// PublicChainChainID is the destination chain ID packed into teleportToPublicChain calldata as
	// destinationChainId. The public chain is never called directly.
	PublicChainChainID uint64 `mapstructure:"PUBLIC_CHAIN_CHAIN_ID"`
	// FaucetPrivateKey funds newly-created user custody wallets so they can pay gas on a
	// non-gasless chain (e.g. the public-chain dev testnet). When empty, wallet funding is
	// disabled (the default — Privacy Nodes are gasless). DEV-ONLY: never set in production.
	FaucetPrivateKey string `mapstructure:"BLOCKCHAIN_FAUCET_PRIVATE_KEY"`
	// FaucetFundAmountWei is the target balance (wei) topped up to each new wallet.
	// Defaults to 5 ETH when a faucet key is set.
	FaucetFundAmountWei string `mapstructure:"BLOCKCHAIN_FAUCET_FUND_WEI"`
	// RoleGrantorPrivateKey signs the on-chain grantRole call that gives each newly-provisioned
	// custody wallet the FACTORY_DEPLOYER role, so users can deploy tokens without an out-of-band
	// grant. The key must hold admin authority over FACTORY_DEPLOYER. When empty, it defaults to
	// BLOCKCHAIN_FAUCET_PRIVATE_KEY; when both are empty, auto-granting is disabled and role
	// assignment stays out of band. DEV-ONLY: never set in production.
	RoleGrantorPrivateKey string `mapstructure:"BLOCKCHAIN_ROLE_GRANTOR_PRIVATE_KEY"`
}

type Config struct {
	Database    Database    `mapstructure:",squash"`
	Auth        Auth        `mapstructure:",squash"`
	Custody     Custody     `mapstructure:",squash"`
	Blockchain  Blockchain  `mapstructure:",squash"`
	PrivacyNode PrivacyNode `mapstructure:",squash"`
	Blockscout  Blockscout  `mapstructure:",squash"`
	Hub         Hub         `mapstructure:",squash"`
	NATS        NATS        `mapstructure:",squash"`
	Mode        string      `mapstructure:"RAYLS_MODE" validate:"omitempty,oneof=privacy_node hub"`
	Logging     string      `mapstructure:"LOGGING"`
	CorsUrls    string      `mapstructure:"CORSURLS"`
	// InstanceName isolates this ops-api deployment when multiple instances share
	// the same NATS server (e.g. dev runs per-participant). When non-empty it is
	// woven into JetStream stream names, durable consumer names, and the second
	// segment of "ops.*" subjects, so two instances never collide on queues or
	// publish into each other's streams. Leave empty for the legacy single-instance
	// behaviour.
	InstanceName string `mapstructure:"INSTANCE_NAME"`
	// Chainless starts the API with NO chain bound: PN_RPC_URL / BLOCKCHAIN_RPC_URL /
	// the registry may all be empty, and every on-chain feature self-disables (role
	// lookup, token deploy, the Blockscout indexers) rather than failing to boot. Auth
	// and admin bootstrap need no chain, so login still works.
	//
	// This exists for the RayUp dev flow (`start_dev.sh --rayup`), where the chain is
	// created later from the playground UI and the API must be up BEFORE it exists.
	// Opt-in on purpose: without it, an empty PN_RPC_URL in privacy_node mode stays a
	// hard error, so a genuinely misconfigured deployment still fails fast.
	Chainless bool `mapstructure:"CHAINLESS"`
	// IdentityPort is the port the SHARED IDENTITY SERVICE listens on (`identity`
	// subcommand). Defaults to 8090 so it can run alongside a per-chain ops-api on
	// 8080 during development. Unused by `run`/`worker`.
	IdentityPort string `mapstructure:"IDENTITY_PORT"`
}

func Load(envPath string) (*Config, error) {
	// Create a new Viper instance with ExperimentalBindStruct enabled
	// This automatically binds all mapstructure tags to environment variables
	v := viper.NewWithOptions(viper.ExperimentalBindStruct())
	v.AutomaticEnv()

	v.SetDefault("PN_BLOCK_RANGE", 100)
	v.SetDefault("PN_POLL_INTERVAL", "2s")
	v.SetDefault("BLOCKCHAIN_STARTING_BLOCK", 0)
	v.SetDefault("BLOCKCHAIN_BLOCK_BATCH_SIZE", 100)

	if envPath != "" {
		// If a specific .env file path is provided, read from it
		v.SetConfigFile(envPath)
		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("failed to read .env file: %w", err)
		}
	}
	// If no path provided, the config will work purely with OS environment variables
	// ExperimentalBindStruct() handles the binding automatically

	// Unmarshal configuration
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Validate configuration
	validate := validator.New()
	if err := validate.Struct(&config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	if _, err := time.ParseDuration(config.PrivacyNode.PollInterval); err != nil {
		return nil, fmt.Errorf(
			"PN_POLL_INTERVAL %q is not a valid duration (e.g. 2s, 500ms): %w",
			config.PrivacyNode.PollInterval,
			err,
		)
	}

	if config.Mode == "privacy_node" && config.PrivacyNode.RPCURL == "" && !config.Chainless {
		return nil, fmt.Errorf(
			"PN_RPC_URL is required when RAYLS_MODE=privacy_node (set CHAINLESS=true to start with no chain bound)",
		)
	}

	if err := config.Auth.validateDeployed(); err != nil {
		return nil, err
	}

	return &config, nil
}
