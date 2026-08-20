package di

import (
	"context"
	"fmt"
	"math/big"

	authlib "github.com/go-pkgz/auth/v2"
	"golang.org/x/oauth2"
	"gorm.io/gorm"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/adapters/blockchain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/adapters/custody"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/adapters/handlers"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/adapters/repositories"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/auth"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/indexer"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/messaging"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/sse"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/config"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/infrastructure/database"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

// Container holds all application dependencies.
type Container struct {
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

	RaylsAccessManagerClient core.RaylsAccessManagerClient
	BootstrapService         core.BootstrapService
	AuthService              core.AuthService

	SIWEProvider *auth.SIWEProvider
	OAuthHandler *auth.OAuthHandler

	TokenRepo         core.TokenRepository
	WalletBalanceRepo core.WalletBalanceRepository

	// Indexer infrastructure — initialised by #36 (mode-conditional startup).
	// Declared here so Container.Close() can drain it regardless of how it was wired.
	NATSManager                  *messaging.Manager
	BlockscoutBackfiller         *indexer.BlockscoutBackfiller
	BlockscoutListener           *indexer.BlockscoutListener
	BlockscoutBalancesBackfiller *indexer.BlockscoutBalancesBackfiller
	BlockscoutBalancesListener   *indexer.BlockscoutBalancesListener

	// Access Manager indexer — nil when blockchain config is incomplete.
	AccessManagerBackfiller       *indexer.AccessManagerBackfiller
	AccessManagerListener         *indexer.AccessManagerListener
	AccessManagerEventHandler     *indexer.AccessManagerEventHandler
	AccessManagerRoleRepo         core.AccessManagerRoleRepository
	AccessManagerRoleMemberRepo   core.AccessManagerRoleMemberRepository
	AccessManagerContractRepo     core.AccessManagerManagedContractRepository
	AccessManagerFnPermRepo       core.AccessManagerFunctionPermissionRepository
	AccessManagerScheduledOpRepo  core.AccessManagerScheduledOperationRepository
	AccessManagerScopedMemberRepo core.AccessManagerContractScopedRoleMemberRepository
	AccessManagerEventLogRepo     core.AccessManagerEventLogRepository

	// OperatorSignerResolver — nil when the Access Manager is not wired (on-chain config incomplete).
	OperatorSignerResolver core.OperatorSignerResolver

	// ChainRoleService resolves the caller's roles ON THIS CHAIN, per request, so a
	// shared identity token (which carries none) can still be authorized. Nil when the
	// Access Manager is not wired; RequireRole then falls back to the JWT claim.
	ChainRoleService *services.ChainRoleService

	// Onboarding (address pairs via RNUserGovernance) — nil when blockchain config is incomplete.
	UserGovernanceService core.UserGovernanceService
	OnboardingService     core.OnboardingService

	TokenRegistryHandler       *handlers.TokenRegistryHandler
	OnboardingHandler          *handlers.OnboardingHandler
	HealthHandler              *handlers.HealthHandler
	AuthHandler                *handlers.AuthHandler
	BootstrapHandler           *handlers.BootstrapHandler
	UserHandler                *handlers.UserHandler
	TokenHandler               *handlers.TokenHandler
	TokenDeployHandler         *handlers.TokenDeployHandler
	TokenStreamHandler         *handlers.TokenStreamHandler
	TokenPermissionHandler     *handlers.TokenPermissionHandler
	TokenActionHandler         *handlers.TokenActionHandler
	AccessManagerHandler       *handlers.AccessManagerHandler
	WalletBalanceHandler       *handlers.WalletBalanceHandler
	WalletBalanceStreamHandler *handlers.WalletBalanceStreamHandler

	TokenStreamHub         *sse.Hub
	WalletBalanceStreamHub *sse.Hub
}

// custodyRPCURL is the chain RPC endpoint sent to custody with every signing request.
//
// Normally the same endpoint this service uses. It differs when custody runs outside our
// network: a SHARED custody (one HSM holding each user's single wallet, which every chain
// must sign through) is not in the cluster, so the in-cluster Service DNS name we dial
// ourselves does not resolve for it and signing fails with "Name or service not known".
// CUSTODY_RPC_URL carries the host-reachable form for that case.
func custodyRPCURL(conf *config.Config) string {
	if conf.Custody.RPCURL != "" {
		return conf.Custody.RPCURL
	}
	return conf.Blockchain.RPCURL
}

// New builds the full dependency graph.
func New(configPath string) (*Container, error) {
	conf, err := config.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	logger.InitializeLogger(conf)
	log := logger.NewLogger()

	db, err := database.Connect(conf.Database.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Accounts (users + custody wallets) come from the SHARED identity database when one is
	// configured; chain-scoped data (am_*, tokens, balances, cursors) always stays in this
	// instance's own database. One custody wallet serves every chain — it is an EVM keypair,
	// so only its on-chain state (balance, roles) differs per chain.
	//
	// Identity owns this schema, so we open the connection WITHOUT running migrations: the
	// identity service applies migrations-identity, and golang-migrate tracks one version
	// sequence per database — a second migrator would fight it.
	accountDB := db
	if conf.Database.IdentityConnectionString != "" {
		identityDB, identityErr := database.Open(conf.Database.IdentityConnectionString)
		if identityErr != nil {
			return nil, fmt.Errorf("failed to connect to identity database: %w", identityErr)
		}
		accountDB = identityDB
		log.Info("accounts resolved from the shared identity database")
	}

	if triggerErr := database.ApplyBlockscoutTrigger(
		context.Background(),
		conf.Blockscout.DBConnString,
		log,
	); triggerErr != nil {
		log.Warn("Failed to apply Blockscout trigger", "error", triggerErr)
	}
	if triggerErr := database.ApplyBlockscoutBalancesTrigger(
		context.Background(),
		conf.Blockscout.DBConnString,
		log,
	); triggerErr != nil {
		log.Warn("Failed to apply Blockscout balances trigger", "error", triggerErr)
	}

	var natsManager *messaging.Manager
	var publisher core.EventPublisher
	var livePublisher core.LiveEventPublisher
	if conf.NATS.URL != "" {
		mgr, natsErr := messaging.NewManager(context.Background(), conf.NATS.URL, messaging.TLSConfig{
			CAFile:   conf.NATS.TLSCAFile,
			CertFile: conf.NATS.TLSCertFile,
			KeyFile:  conf.NATS.TLSKeyFile,
		})
		if natsErr != nil {
			log.Warn("Failed to connect to NATS — token events will not be published", "error", natsErr)
		} else {
			amStreamName := indexer.InstanceJSName(conf.InstanceName, indexer.StreamAccessManager)
			amStreamSubject := indexer.InstanceSubject(conf.InstanceName, "ops.access_manager.>")
			if streamErr := mgr.EnsureStream(
				context.Background(),
				amStreamName,
				[]string{amStreamSubject},
			); streamErr != nil {
				log.Warn("Failed to ensure access manager stream", "name", amStreamName, "error", streamErr)
			}
			tokenStreamName := indexer.InstanceJSName(conf.InstanceName, indexer.StreamTokens)
			tokenStreamSubject := indexer.InstanceSubject(conf.InstanceName, "ops.tokens.>")
			if streamErr := mgr.EnsureStream(
				context.Background(),
				tokenStreamName,
				[]string{tokenStreamSubject},
			); streamErr != nil {
				log.Warn("Failed to ensure tokens stream", "name", tokenStreamName, "error", streamErr)
			}
			walletBalanceStreamName := indexer.InstanceJSName(conf.InstanceName, indexer.StreamWalletBalances)
			walletBalanceStreamSubject := indexer.InstanceSubject(conf.InstanceName, "ops.wallet_balances.>")
			if streamErr := mgr.EnsureStream(
				context.Background(),
				walletBalanceStreamName,
				[]string{walletBalanceStreamSubject},
			); streamErr != nil {
				log.Warn("Failed to ensure wallet balances stream", "name", walletBalanceStreamName, "error", streamErr)
			}
			natsManager = mgr
			publisher = mgr.NewPublisher()
			livePublisher = mgr
		}
	}

	tokenRepo := repositories.NewTokenRepository(db)
	indexerStateRepo := repositories.NewIndexerStateRepository(db)
	walletBalanceRepo := repositories.NewWalletBalanceRepository(db)

	// Account repositories read from accountDB — the shared identity database when one is
	// configured, otherwise this instance's own. The transactor follows them: every caller
	// that uses it (bootstrap, auth, onboarding) writes accounts, and a transaction opened
	// on the chain database could not cover them.
	userRepo := repositories.NewUserRepository(accountDB)
	nonceRepo := repositories.NewNonceRepository(accountDB)
	oauthProviderRepo := repositories.NewUserOAuthProviderRepository(accountDB)
	walletRepo := repositories.NewUserWalletRepository(accountDB)
	tokenBlacklistRepo := repositories.NewTokenBlacklistRepository(accountDB)
	txer := database.NewTransactor(accountDB)

	var blockscoutBackfiller *indexer.BlockscoutBackfiller
	var blockscoutListener *indexer.BlockscoutListener
	var blockscoutBalancesBackfiller *indexer.BlockscoutBalancesBackfiller
	var blockscoutBalancesListener *indexer.BlockscoutBalancesListener
	if conf.Blockscout.DBConnString != "" {
		// Guard against ingesting another network's tokens: the backfill/listener stamp
		// every Blockscout token with conf.Blockscout.ChainID as its issuer_id without
		// checking that the chain behind our RPC is actually that chain. If this instance
		// is wired to the wrong environment's infrastructure (e.g. a fresh testnet ops-api
		// still pointing at the devnet Blockscout/RPC), that would relabel foreign tokens
		// as ours. Fail fast on a chain-ID mismatch instead.
		if err := indexer.VerifyChainID(
			context.Background(),
			conf.Blockchain.RPCURL,
			conf.Blockscout.ChainID,
			log,
		); err != nil {
			return nil, fmt.Errorf("blockscout chain-ID verification failed: %w", err)
		}
		blockscoutBackfiller = indexer.NewBlockscoutBackfiller(
			conf.Blockscout.DBConnString, tokenRepo, indexerStateRepo,
			conf.Blockscout.BackfillBatch, log, conf.Blockscout.ChainID,
		)
		blockscoutListener = indexer.NewBlockscoutListener(
			conf.Blockscout.DBConnString,
			tokenRepo,
			publisher,
			livePublisher,
			log,
			conf.InstanceName,
			conf.Blockscout.ChainID,
		)
		blockscoutBalancesBackfiller = indexer.NewBlockscoutBalancesBackfiller(
			conf.Blockscout.DBConnString, walletBalanceRepo, walletRepo, indexerStateRepo,
			conf.Blockscout.BackfillBatch, log,
		)
		blockscoutBalancesListener = indexer.NewBlockscoutBalancesListener(
			conf.Blockscout.DBConnString, walletBalanceRepo, walletRepo,
			publisher, livePublisher, log, conf.InstanceName,
		)
	}

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
			custodyRPCURL(conf),
		)
	} else {
		custodyService = custody.NewSelfCustody()
		custodyProviderType = domain.CustodyProviderSelf
		if conf.Auth.GoogleClientID != "" || conf.Auth.MicrosoftClientID != "" {
			log.Warn("CUSTODY_RAYLS_HSM_URL is not set — bootstrap endpoint requires HSM custody")
		}
	}

	var ramClient core.RaylsAccessManagerClient
	if conf.Blockchain.RPCURL != "" && conf.Blockchain.DeploymentProxyRegistryAddr != "" {
		ramSvc, ramErr := blockchain.NewRaylsAccessManagerService(
			conf.Blockchain.RPCURL,
			conf.Blockchain.DeploymentProxyRegistryAddr,
		)
		if ramErr != nil {
			return nil, fmt.Errorf("failed to init RaylsAccessManager service: %w", ramErr)
		}
		ramClient = ramSvc
	} else {
		log.Warn(
			"blockchain config incomplete — role lookup disabled (set BLOCKCHAIN_RPC_URL, DEPLOYMENT_PROXY_REGISTRY_ADDR)",
		)
		ramClient = &noopRaylsAccessManagerClient{}
	}

	// Wallet funder — DEV-ONLY, wired only when a faucet key is configured (non-gasless
	// chains, e.g. the public-chain dev testnet). nil on gasless Privacy Nodes / production,
	// in which case new custody wallets are simply not auto-funded. Reuses the RAM RPC client.
	var walletFunder core.WalletFunder
	if conf.Blockchain.FaucetPrivateKey != "" {
		if ramSvc, ok := ramClient.(*blockchain.RaylsAccessManagerService); ok {
			funder, funderErr := blockchain.NewFaucetFunder(
				ramSvc.EthClient(), conf.Blockchain.FaucetPrivateKey, conf.Blockchain.FaucetFundAmountWei,
			)
			if funderErr != nil {
				return nil, fmt.Errorf("failed to init wallet faucet funder: %w", funderErr)
			}
			walletFunder = funder
			log.Info("wallet faucet funder enabled — new custody wallets will be auto-funded for gas")
		} else {
			log.Warn("BLOCKCHAIN_FAUCET_PRIVATE_KEY set but blockchain RPC is not configured — wallet funding disabled")
		}
	}

	// Role granter — DEV-ONLY, wired only when a grantor key is configured (defaults to the faucet
	// key). Grants FACTORY_DEPLOYER + PRIVACY_NODE_OPERATOR to each new custody wallet at provision
	// time so users can deploy tokens and pass the login role check without an out-of-band grant.
	// nil when no grantor key is set OR the blockchain is not configured, in which case role
	// assignment stays out of band. Reuses the RAM RPC client.
	var roleGranter core.RoleGranter
	grantorKey := conf.Blockchain.RoleGrantorPrivateKey
	if grantorKey == "" {
		grantorKey = conf.Blockchain.FaucetPrivateKey
	}
	if grantorKey != "" {
		if ramSvc, ok := ramClient.(*blockchain.RaylsAccessManagerService); ok {
			granter, granterErr := blockchain.NewFactoryDeployerGranter(
				ramSvc.EthClient(), conf.Blockchain.DeploymentProxyRegistryAddr, grantorKey,
			)
			if granterErr != nil {
				return nil, fmt.Errorf("failed to init factory-deployer role granter: %w", granterErr)
			}
			roleGranter = granter
			log.Info(
				"role granter enabled — new custody wallets will be auto-granted FACTORY_DEPLOYER + PRIVACY_NODE_OPERATOR",
			)
		} else {
			log.Warn("role grantor key set but blockchain RPC is not configured — auto role grant disabled")
		}
	}

	// Token deploy via RNContractFactory — wired only when blockchain is configured
	// AND the RNContractFactory is registered in the DeploymentProxyRegistry. Reuses
	// the RaylsAccessManager RPC client to avoid a second dial.
	//
	// RNContractFactory is a Privacy Node concept; the public-chain deployment does not
	// deploy or register it. On such targets the factory init fails with a zero address —
	// that is not fatal: we just leave the deploy route unregistered (api.go nil-guards
	// TokenDeployHandler), same as the noop fallbacks above for an incomplete blockchain.
	var tokenDeployHandler *handlers.TokenDeployHandler
	if ramSvc, ok := ramClient.(*blockchain.RaylsAccessManagerService); ok {
		factorySvc, factoryErr := blockchain.NewRaylsContractFactoryService(
			ramSvc.EthClient(), conf.Blockchain.DeploymentProxyRegistryAddr, custodyService,
		)
		if factoryErr != nil {
			log.Warn("RNContractFactory unavailable — token deploy disabled", "error", factoryErr)
		} else {
			tokenDeployHandler = handlers.NewTokenDeployHandler(factorySvc, walletRepo, tokenRepo, log)
		}
	}

	// Access Manager indexer — wired only when blockchain is fully configured.
	var amBackfiller *indexer.AccessManagerBackfiller
	var amListener *indexer.AccessManagerListener
	var amEventHandler *indexer.AccessManagerEventHandler
	var amRoleRepo core.AccessManagerRoleRepository
	var amMemberRepo core.AccessManagerRoleMemberRepository
	var amContractRepo core.AccessManagerManagedContractRepository
	var amFnPermRepo core.AccessManagerFunctionPermissionRepository
	var amScheduledOpRepo core.AccessManagerScheduledOperationRepository
	var amScopedMemberRepo core.AccessManagerContractScopedRoleMemberRepository
	var amEventLogRepo core.AccessManagerEventLogRepository
	if conf.Blockchain.RPCURL != "" && conf.Blockchain.DeploymentProxyRegistryAddr != "" {
		ramSvc, ok := ramClient.(*blockchain.RaylsAccessManagerService)
		if ok {
			amRoleRepo = repositories.NewAccessManagerRoleRepository(db)
			amMemberRepo = repositories.NewAccessManagerRoleMemberRepository(db)
			amContractRepo = repositories.NewAccessManagerManagedContractRepository(db)
			amFnPermRepo = repositories.NewAccessManagerFunctionPermissionRepository(db)
			amScheduledOpRepo = repositories.NewAccessManagerScheduledOperationRepository(db)
			amScopedMemberRepo = repositories.NewAccessManagerContractScopedRoleMemberRepository(db)
			amEventLogRepo = repositories.NewAccessManagerEventLogRepository(db)

			amEventHandler = indexer.NewAccessManagerEventHandler(
				amRoleRepo, amMemberRepo, amContractRepo,
				amFnPermRepo, amScheduledOpRepo, amScopedMemberRepo,
				amEventLogRepo, log,
			)
			amListener = indexer.NewAccessManagerListener(
				ramSvc.EthClient(), ramSvc.Contract(), ramSvc.ContractAddress(),
				indexerStateRepo, publisher, conf, log,
			)
			amBackfiller = indexer.NewAccessManagerBackfiller(
				ramSvc.EthClient(), ramSvc.Contract(), ramSvc.ContractAddress(),
				indexerStateRepo, publisher, conf, log,
			)
		}
	}

	var amHandler *handlers.AccessManagerHandler
	var tokenPermissionHandler *handlers.TokenPermissionHandler
	var tokenActionHandler *handlers.TokenActionHandler
	var tokenRegistryHandler *handlers.TokenRegistryHandler
	var operatorSignerResolver core.OperatorSignerResolver
	var chainRoleService *services.ChainRoleService
	var userGovernanceService core.UserGovernanceService
	var onboardingService core.OnboardingService
	var onboardingHandler *handlers.OnboardingHandler
	if amRoleRepo != nil {
		amHandler = handlers.NewAccessManagerHandler(
			amRoleRepo, amMemberRepo, amContractRepo,
			amFnPermRepo, amScheduledOpRepo, amEventLogRepo, log,
		)

		// Operator-signer resolution — resolves the HSM operator wallet for governance writes.
		operatorSignerResolver = services.NewOperatorSignerResolver(
			amRoleRepo, amMemberRepo, walletRepo, log,
		)

		// Per-request role resolution for THIS chain. Needed once identity is shared: a
		// token minted by the identity service carries no roles (a grant on one chain says
		// nothing about another), so authorization is answered here, from the am_* tables
		// this instance's indexer maintains. Nil when the Access Manager is not wired, and
		// RequireRole then falls back to the JWT claim as before.
		//
		// The wallet it resolves roles for is the user's ONE identity custody wallet; this
		// service does not mint per-chain wallets.
		chainRoleService = services.NewChainRoleService(
			walletRepo, userRepo, amMemberRepo, amRoleRepo, log,
		)
		// tokens/actions are supplied below, once the on-chain token service exists: the
		// stablecoin pause capability cannot be derived from the am_* tables (pause() is gated
		// on the contract's own `pauser` address, not an AM role), so it needs a live read.
		// Left nil when the Access Manager is not wired — CanPause then stays false.
		tokenPermSvc := services.NewTokenPermissionService(
			amMemberRepo,
			amScopedMemberRepo,
			amFnPermRepo,
			amContractRepo,
			tokenRepo,
			nil,
		)
		tokenPermissionHandler = handlers.NewTokenPermissionHandler(tokenPermSvc, walletRepo, log)

		// Mint/burn + onboarding — reuse the RaylsAccessManager RPC client; gated on the same AM availability.
		if ramSvc, ok := ramClient.(*blockchain.RaylsAccessManagerService); ok {
			// Token registry — TokenRegistry catalog writes signed by the resolved operator wallet.
			// Built first because the teleport service uses its Exists read for the preflight.
			tokenRegistryAdapter, tokenRegistryErr := blockchain.NewRaylsTokenRegistryService(
				ramSvc.EthClient(), conf.Blockchain.DeploymentProxyRegistryAddr, custodyService,
			)
			if tokenRegistryErr != nil {
				return nil, fmt.Errorf("failed to init TokenRegistry service: %w", tokenRegistryErr)
			}
			tokenRegistrySvc := services.NewTokenRegistryService(operatorSignerResolver, tokenRegistryAdapter, log)
			tokenRegistryHandler = handlers.NewTokenRegistryHandler(tokenRegistrySvc, log)

			// Auto-register + authorize deployed tokens so the owner can mint/transfer right away.
			if tokenDeployHandler != nil {
				tokenDeployHandler.SetTokenRegistry(tokenRegistrySvc)
			}

			// Mint/burn (pure on-chain token writes) + teleport on-chain client. The TeleportService
			// owns the destination (public) chain and passes it to the client as destinationChainId;
			// the teleport business rules (eligibility + registry Exists + balance/ownership
			// preflight) live there too.
			tokenActionSvc, tokenSvcErr := blockchain.NewRaylsTokenService(ramSvc.EthClient(), custodyService)
			if tokenSvcErr != nil {
				return nil, fmt.Errorf("failed to init token action service: %w", tokenSvcErr)
			}
			// Now that a chain client exists, let the permission service answer the stablecoin
			// pause capability (a live `pauser`/`paused` read — see resolvePause).
			tokenPermSvc.SetTokenActionService(tokenActionSvc)
			// Teleport is optional: it needs the public (destination) chain id. Private-chain-only
			// deployments leave PUBLIC_CHAIN_CHAIN_ID unset — mint/burn stay available and the
			// teleport endpoint rejects requests with 501 (guarded in the handler).
			var teleportSvc core.TeleportService
			if conf.Blockchain.PublicChainChainID != 0 {
				teleportSvc = services.NewTeleportService(
					tokenActionSvc,
					tokenRegistrySvc,
					new(big.Int).SetUint64(conf.Blockchain.PublicChainChainID),
					log,
				)
			} else {
				log.Warn("PUBLIC_CHAIN_CHAIN_ID not set — teleport disabled (mint/burn remain available)")
			}
			tokenActionHandler = handlers.NewTokenActionHandler(
				tokenActionSvc,
				teleportSvc,
				tokenPermSvc,
				tokenRepo,
				walletRepo,
				log,
			)

			// Onboarding — RNUserGovernance writes signed by the resolved operator wallet.
			userGovSvc, userGovErr := blockchain.NewRaylsUserGovernanceService(
				ramSvc.EthClient(), conf.Blockchain.DeploymentProxyRegistryAddr, custodyService,
			)
			if userGovErr != nil {
				return nil, fmt.Errorf("failed to init RNUserGovernance service: %w", userGovErr)
			}
			userGovernanceService = userGovSvc
			onboardingService = services.NewOnboardingService(
				custodyService,
				walletRepo,
				userRepo,
				operatorSignerResolver,
				userGovSvc,
				txer,
				log,
			)
			onboardingHandler = handlers.NewOnboardingHandler(onboardingService, log)
		}
	}

	bootstrapService := services.NewBootstrapService(
		userRepo,
		walletRepo,
		oauthProviderRepo,
		txer,
		custodyService,
		walletFunder,
		log,
	)

	// Auto-provision on login: every new signup gets a custody wallet + role_assigned so they
	// can use the app immediately (no operator approval). Reuses the custody/funder chosen above.
	provisioningService := services.NewProvisioningService(
		userRepo,
		walletRepo,
		custodyService,
		custodyProviderType,
		walletFunder,
		roleGranter,
		log,
	)
	authService := services.NewAuthService(
		userRepo,
		nonceRepo,
		oauthProviderRepo,
		walletRepo,
		txer,
		ramClient,
		provisioningService,
		conf.Auth.BaseURL,
		conf.Chainless,
		log,
	)

	pkgzAuth := auth.NewPkgzAuthService(conf.Auth, conf.Auth.BaseURL)
	tokenWrapper := auth.NewTokenWrapper(pkgzAuth.TokenService(), conf.Auth.JWTSecret, walletRepo, conf.Auth.AdminEmail)

	siweProvider := auth.NewSIWEProvider(authService, tokenWrapper, log)
	pkgzAuth.AddCustomHandler(siweProvider)

	// OAuth redirect_uri base: a fixed relay host when set (so one URI serves many
	// per-instance hosts), else this instance's own BASE_URL.
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

	oauthHandler := auth.NewOAuthHandler(
		authService,
		tokenWrapper,
		googleCfg,
		microsoftCfg,
		conf.Auth.PostLoginRedirectURL,
		conf.Auth.BaseURL,
		conf.Auth.EmailSignupEnabled,
		// No signup-details repo: sign-up answers belong to identity, and this database
		// dropped the identity tables in migration 000009.
		nil,
		log,
	)

	healthHandler := handlers.NewHealthHandler(log)
	authHandler := handlers.NewAuthHandler(
		authService,
		tokenWrapper,
		tokenBlacklistRepo,
		walletRepo,
		userRepo,
		ramClient,
		log,
	)
	bootstrapHandler := handlers.NewBootstrapHandler(bootstrapService, conf.Auth.BootstrapToken, log)
	userHandler := handlers.NewUserHandler(userRepo, walletRepo, log)
	tokenHandler := handlers.NewTokenHandler(tokenRepo, log)

	tokenStreamHub := sse.NewHub()
	tokenStreamHandler := handlers.NewTokenStreamHandler(tokenStreamHub, log)

	walletBalanceStreamHub := sse.NewHub()
	walletBalanceSvc := services.NewWalletBalanceService(walletBalanceRepo, walletRepo, tokenRepo, log)
	walletBalanceHandler := handlers.NewWalletBalanceHandler(walletBalanceSvc, log)
	walletBalanceStreamHandler := handlers.NewWalletBalanceStreamHandler(walletBalanceStreamHub, walletRepo, log)

	return &Container{
		Config:                       conf,
		DB:                           db,
		Logger:                       log,
		PkgzAuth:                     pkgzAuth,
		TokenWrapper:                 tokenWrapper,
		UserRepo:                     userRepo,
		NonceRepo:                    nonceRepo,
		OAuthProviderRepo:            oauthProviderRepo,
		WalletRepo:                   walletRepo,
		TokenBlacklistRepo:           tokenBlacklistRepo,
		RaylsAccessManagerClient:     ramClient,
		BootstrapService:             bootstrapService,
		AuthService:                  authService,
		SIWEProvider:                 siweProvider,
		OAuthHandler:                 oauthHandler,
		TokenRepo:                    tokenRepo,
		WalletBalanceRepo:            walletBalanceRepo,
		TokenStreamHandler:           tokenStreamHandler,
		TokenStreamHub:               tokenStreamHub,
		WalletBalanceHandler:         walletBalanceHandler,
		WalletBalanceStreamHandler:   walletBalanceStreamHandler,
		WalletBalanceStreamHub:       walletBalanceStreamHub,
		NATSManager:                  natsManager,
		BlockscoutBackfiller:         blockscoutBackfiller,
		BlockscoutListener:           blockscoutListener,
		BlockscoutBalancesBackfiller: blockscoutBalancesBackfiller,
		BlockscoutBalancesListener:   blockscoutBalancesListener,

		AccessManagerBackfiller:       amBackfiller,
		AccessManagerListener:         amListener,
		AccessManagerEventHandler:     amEventHandler,
		AccessManagerRoleRepo:         amRoleRepo,
		AccessManagerRoleMemberRepo:   amMemberRepo,
		AccessManagerContractRepo:     amContractRepo,
		AccessManagerFnPermRepo:       amFnPermRepo,
		AccessManagerScheduledOpRepo:  amScheduledOpRepo,
		AccessManagerScopedMemberRepo: amScopedMemberRepo,
		AccessManagerEventLogRepo:     amEventLogRepo,

		OperatorSignerResolver: operatorSignerResolver,
		ChainRoleService:       chainRoleService,
		UserGovernanceService:  userGovernanceService,
		OnboardingService:      onboardingService,

		HealthHandler:          healthHandler,
		AuthHandler:            authHandler,
		BootstrapHandler:       bootstrapHandler,
		UserHandler:            userHandler,
		TokenHandler:           tokenHandler,
		TokenDeployHandler:     tokenDeployHandler,
		TokenPermissionHandler: tokenPermissionHandler,
		TokenActionHandler:     tokenActionHandler,
		TokenRegistryHandler:   tokenRegistryHandler,
		AccessManagerHandler:   amHandler,
		OnboardingHandler:      onboardingHandler,
	}, nil
}

// Close releases DB resources and drains the NATS connection if wired.
func (c *Container) Close() error {
	if c.NATSManager != nil {
		c.NATSManager.Close()
	}
	return database.Disconnect(c.DB)
}

type noopRaylsAccessManagerClient struct{}

func (n *noopRaylsAccessManagerClient) GetRoles(_ context.Context, _ string) ([]string, error) {
	return nil, nil
}
