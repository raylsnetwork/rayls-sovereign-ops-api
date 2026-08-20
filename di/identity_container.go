package di

import (
	"fmt"

	authlib "github.com/go-pkgz/auth/v2"
	"golang.org/x/oauth2"
	"gorm.io/gorm"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/adapters/custody"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/adapters/handlers"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/adapters/repositories"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/auth"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/config"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/infrastructure/database"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

// IdentityContainer holds the dependencies of the SHARED IDENTITY SERVICE.
//
// One identity service serves every chain. It answers exactly one question — "who is
// this person?" — and knows nothing about any chain: no RPC client, no AccessManager,
// no indexers, no tokens. Its JWT therefore carries identity only (id, name, email,
// auth_method, is_admin); roles and signing wallets are per-chain facts that each
// chain's ops-api resolves for itself against its own database.
//
// Deliberately a separate, much smaller graph than Container: sharing that one would
// drag the whole chain-side dependency tree (and its config requirements) into a
// service that must boot with no chain at all.
type IdentityContainer struct {
	Config *config.Config
	DB     *gorm.DB
	Logger logger.Logger

	PkgzAuth     *authlib.Service
	TokenWrapper *auth.TokenWrapper

	UserRepo           core.UserRepository
	NonceRepo          core.NonceRepository
	OAuthProviderRepo  core.UserOAuthProviderRepository
	WalletRepo         core.UserWalletRepository
	TokenBlacklistRepo core.TokenBlacklistRepository
	SignupDetailsRepo  core.UserSignupDetailsRepository

	AuthService      core.AuthService
	BootstrapService core.BootstrapService

	SIWEProvider *auth.SIWEProvider
	OAuthHandler *auth.OAuthHandler

	HealthHandler    *handlers.HealthHandler
	AuthHandler      *handlers.AuthHandler
	BootstrapHandler *handlers.BootstrapHandler
	UserHandler      *handlers.UserHandler
}

// NewIdentity builds the identity service's dependency graph.
func NewIdentity(configPath string) (*IdentityContainer, error) {
	conf, err := config.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	logger.InitializeLogger(conf)
	log := logger.NewLogger()

	// Its own schema: identity tables only. Connect() would apply the ops-api migrations,
	// creating chain tables nothing here maintains and leaving the version sequence at a
	// number the identity migrations do not have.
	db, err := database.ConnectWithMigrations(conf.Database.ConnectionString, "migrations-identity")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	userRepo := repositories.NewUserRepository(db)
	nonceRepo := repositories.NewNonceRepository(db)
	oauthProviderRepo := repositories.NewUserOAuthProviderRepository(db)
	walletRepo := repositories.NewUserWalletRepository(db)
	tokenBlacklistRepo := repositories.NewTokenBlacklistRepository(db)
	signupDetailsRepo := repositories.NewUserSignupDetailsRepository(db)
	txer := database.NewTransactor(db)

	// Custody mints the wallet a bootstrapped admin signs with. Wallet creation is
	// chain-agnostic (only SignAndTransact names a chain — see adapters/custody/rayls_hsm.go),
	// so it is usable here; the wallets it mints only acquire meaning once a chain funds
	// them and grants them roles, which is per-chain work. The empty RPC URL below reflects
	// that: this container is chain-less and never signs, so it has no chain to point at.
	var custodyService core.CustodyService
	custodyProviderType := domain.CustodyProviderRaylsHSM
	if conf.Custody.RaylsHSMURL != "" {
		if conf.Custody.RaylsHSMAPIKey == "" {
			log.Warn("CUSTODY_RAYLS_HSM_API_KEY is not set — HSM requests will be unauthenticated")
		}
		custodyService = custody.NewRaylsHSM(
			conf.Custody.RaylsHSMURL,
			conf.Custody.RaylsHSMAPIKey,
			conf.Custody.RaylsHSMPassword,
			"",
		)
	} else {
		// SelfCustody cannot mint (SIWE users bring their own address), so /admin/bootstrap
		// will fail without an HSM. SIWE and OAuth login still work — only the bootstrapped
		// admin needs a minted wallet.
		log.Warn(
			"No custody service configured (CUSTODY_RAYLS_HSM_URL) — /admin/bootstrap will fail; SIWE/OAuth login is unaffected",
		)
		custodyService = custody.NewSelfCustody()
		custodyProviderType = domain.CustodyProviderSelf
	}

	// No funder: topping a wallet up with gas is chain-specific, and the identity
	// service has no chain. The per-chain ops-api funds wallets on its own chain.
	bootstrapService := services.NewBootstrapService(
		userRepo,
		walletRepo,
		oauthProviderRepo,
		txer,
		custodyService,
		nil,
		log,
	)

	// Activates newly registered accounts AND mints their one custody wallet. Identity owns
	// wallets: an EVM keypair works on every chain, so minting here is what gives a user the
	// same address everywhere — and is the only way a chain can resolve a wallet for someone
	// who has never touched it. Funding and role grants stay per-chain (this service has no
	// chain). Every login path — Google, Microsoft, SIWE, standalone email — goes through
	// this, so all of them produce an account that works identically.
	provisioningService := services.NewIdentityProvisioningService(
		userRepo, walletRepo, custodyService, custodyProviderType, log,
	)

	// authService with NO chain:
	//   ramClient   = nil — there is no AccessManager to read roles from
	//   chainless   = true — skips the on-chain role check in the login decision tree
	//
	// The chainless branch returns PRIVACY_NODE_OPERATOR so a dev instance stays usable;
	// that role is dropped when the token is minted (walletRepo/roles omitted below),
	// because authorization is the per-chain ops-api's job.
	authService := services.NewAuthService(
		userRepo, nonceRepo, oauthProviderRepo, walletRepo, txer,
		nil, provisioningService, conf.Auth.BaseURL, true, log,
	)

	pkgzAuth := auth.NewPkgzAuthService(conf.Auth, conf.Auth.BaseURL)

	// Identity-only: no roles (a per-chain fact this token cannot assert everywhere). The
	// custody wallet IS included — one EVM keypair per user, valid on every chain.
	tokenWrapper := auth.NewIdentityTokenWrapper(
		pkgzAuth.TokenService(), conf.Auth.JWTSecret, walletRepo, conf.Auth.AdminEmail,
	)

	siweProvider := auth.NewSIWEProvider(authService, tokenWrapper, log)
	pkgzAuth.AddCustomHandler(siweProvider)

	// One registered OAuth redirect_uri fronts every chain: providers only redirect to
	// pre-registered URIs, and chains are created on demand. OAUTH_REDIRECT_BASE points
	// at the playground, which relays the callback on using the URL carried in the state.
	redirectBase := conf.Auth.OAuthRedirectBase
	if redirectBase == "" {
		redirectBase = conf.Auth.BaseURL
	}

	var googleCfg, microsoftCfg *oauth2.Config
	if conf.Auth.GoogleClientID != "" {
		log.Info("Google OAuth configured")
		googleCfg = auth.NewGoogleOAuthConfig(conf.Auth.GoogleClientID, conf.Auth.GoogleClientSecret, redirectBase)
	}
	if conf.Auth.MicrosoftClientID != "" {
		log.Info("Microsoft OAuth configured")
		microsoftCfg = auth.NewMicrosoftOAuthConfig(
			conf.Auth.MicrosoftClientID,
			conf.Auth.MicrosoftClientSecret,
			redirectBase,
		)
	}
	if conf.Auth.BootstrapToken == "" {
		log.Warn("BOOTSTRAP_TOKEN is not set — POST /admin/bootstrap is unauthenticated. " +
			"On a fresh database the first caller becomes the administrator of every chain; " +
			"set a token or keep the endpoint unreachable from outside.")
	}
	if conf.Auth.EmailSignupEnabled {
		log.Warn("EMAIL_SIGNUP_ENABLED=true — /auth/signup issues a session from an unverified " +
			"email address. Local development only: anyone who can reach it can log in as anyone.")
	}
	oauthHandler := auth.NewOAuthHandler(
		authService, tokenWrapper, googleCfg, microsoftCfg,
		conf.Auth.PostLoginRedirectURL, conf.Auth.BaseURL, conf.Auth.EmailSignupEnabled,
		signupDetailsRepo, log,
	)

	// ramClient and walletRepo nil → Refresh re-issues an identity-only token instead of
	// looking up on-chain roles (see AuthHandler.Refresh).
	authHandler := handlers.NewAuthHandler(authService, tokenWrapper, tokenBlacklistRepo, nil, userRepo, nil, log)

	return &IdentityContainer{
		Config: conf,
		DB:     db,
		Logger: log,

		PkgzAuth:     pkgzAuth,
		TokenWrapper: tokenWrapper,

		UserRepo:           userRepo,
		NonceRepo:          nonceRepo,
		OAuthProviderRepo:  oauthProviderRepo,
		WalletRepo:         walletRepo,
		TokenBlacklistRepo: tokenBlacklistRepo,
		SignupDetailsRepo:  signupDetailsRepo,

		AuthService:      authService,
		BootstrapService: bootstrapService,

		SIWEProvider: siweProvider,
		OAuthHandler: oauthHandler,

		HealthHandler:    handlers.NewHealthHandler(log),
		AuthHandler:      authHandler,
		BootstrapHandler: handlers.NewBootstrapHandler(bootstrapService, conf.Auth.BootstrapToken, log),
		UserHandler:      handlers.NewUserHandler(userRepo, walletRepo, log),
	}, nil
}

// Close releases the identity service's resources.
func (c *IdentityContainer) Close() error {
	if c.DB != nil {
		sqlDB, err := c.DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
