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

func newProvisioningSvc(
	userRepo *testutil.FakeUserRepository,
	walletRepo *testutil.FakeUserWalletRepository,
	custody core.CustodyService,
) core.ProvisioningService {
	return NewProvisioningService(
		userRepo,
		walletRepo,
		custody,
		domain.CustodyProviderRaylsHSM,
		nil,
		nil,
		&testutil.StubLogger{},
	)
}

// fakeRoleGranter records the addresses it was asked to grant roles to (and the last role set),
// and can be configured to fail so callers can assert the grant is best-effort.
type fakeRoleGranter struct {
	granted   []string
	lastRoles []string
	err       error
}

func (g *fakeRoleGranter) GrantRoles(_ context.Context, address string, roles []string) error {
	g.granted = append(g.granted, address)
	g.lastRoles = roles
	return g.err
}

// ── Provision ─────────────────────────────────────────────────────────────────

func TestProvisioningService_Provision_AlreadyProvisionedDoesNotRecreateWallet(t *testing.T) {
	// A role_assigned user with a wallet is not re-created via custody on subsequent logins
	user := &domain.User{Status: domain.UserStatusRoleAssigned}
	user.ID = uuid.New()

	walletRepo := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{UserID: user.ID, RaylsAddress: "0xexisting", IsActive: true}},
	}
	custody := &testutil.FakeCustodyService{Err: errors.New("should not be called")}
	svc := newProvisioningSvc(testutil.NewFakeUserRepository(), walletRepo, custody)

	err := svc.Provision(context.Background(), user)

	require.NoError(t, err)
	require.Len(t, walletRepo.Wallets, 1)
}

func TestProvisioningService_Provision_RegrantsFactoryDeployerOnEveryLogin(t *testing.T) {
	// The FACTORY_DEPLOYER grant is re-checked on every login, self-healing wallets provisioned
	// before auto-granting existed (the granter no-ops when the role is already held).
	user := &domain.User{Status: domain.UserStatusRoleAssigned}
	user.ID = uuid.New()

	walletRepo := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{UserID: user.ID, RaylsAddress: "0xexisting", IsActive: true}},
	}
	custody := &testutil.FakeCustodyService{Err: errors.New("should not be called")}
	granter := &fakeRoleGranter{}

	svc := NewProvisioningService(
		testutil.NewFakeUserRepository(),
		walletRepo,
		custody,
		domain.CustodyProviderRaylsHSM,
		nil,
		granter,
		&testutil.StubLogger{},
	)

	err := svc.Provision(context.Background(), user)

	require.NoError(t, err)
	require.Equal(t, []string{"0xexisting"}, granter.granted)
}

func TestProvisioningService_Provision_UsesExistingWallet(t *testing.T) {
	// If the user already has a wallet, custody is not called and the status is advanced
	user := &domain.User{Status: domain.UserStatusWaitingRoleAssignment}
	user.ID = uuid.New()

	existingWallet := domain.UserWallet{
		UserID:       user.ID,
		RaylsAddress: "0xexisting",
		IsActive:     true,
	}
	walletRepo := &testutil.FakeUserWalletRepository{Wallets: []domain.UserWallet{existingWallet}}
	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{*user}
	custody := &testutil.FakeCustodyService{Err: errors.New("should not be called")}

	svc := newProvisioningSvc(userRepo, walletRepo, custody)

	err := svc.Provision(context.Background(), user)

	require.NoError(t, err)
	assert.Equal(t, domain.UserStatusRoleAssigned, user.Status)
}

func TestProvisioningService_Provision_CreatesCustodyWalletWhenMissing(t *testing.T) {
	// If the user has no wallet, custody is called and the new wallet is persisted
	user := &domain.User{Status: domain.UserStatusWaitingRoleAssignment}
	user.ID = uuid.New()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{*user}
	walletRepo := &testutil.FakeUserWalletRepository{}
	custody := testutil.NewFakeCustodyService("0xcustody", "ext-123")

	svc := newProvisioningSvc(userRepo, walletRepo, custody)

	err := svc.Provision(context.Background(), user)

	require.NoError(t, err)
	require.Len(t, walletRepo.Wallets, 1)
	assert.Equal(t, "0xcustody", walletRepo.Wallets[0].RaylsAddress)
	assert.Equal(t, "ext-123", walletRepo.Wallets[0].CustodyExternalID)
	assert.Equal(t, domain.CustodyProviderRaylsHSM, walletRepo.Wallets[0].CustodyProvider)
	assert.Equal(t, domain.UserStatusRoleAssigned, user.Status)
}

func TestProvisioningService_Provision_UpdatesUserStatusInRepository(t *testing.T) {
	// On success, the user status is updated to role_assigned in the repository
	user := &domain.User{Status: domain.UserStatusWaitingRoleAssignment}
	user.ID = uuid.New()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{*user}
	walletRepo := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{UserID: user.ID, RaylsAddress: "0xaddr", IsActive: true}},
	}

	svc := newProvisioningSvc(userRepo, walletRepo, testutil.NewFakeCustodyService("", ""))

	err := svc.Provision(context.Background(), user)

	require.NoError(t, err)
	assert.Equal(t, domain.UserStatusRoleAssigned, user.Status)
	// Verify the DB was updated too
	assert.Equal(t, domain.UserStatusRoleAssigned, userRepo.Users[0].Status)
}

func TestProvisioningService_Provision_GrantsDeployAndOperatorRolesToNewWallet(t *testing.T) {
	// When a role granter is wired, a newly-created custody wallet is granted FACTORY_DEPLOYER + PRIVACY_NODE_OPERATOR
	user := &domain.User{Status: domain.UserStatusWaitingRoleAssignment}
	user.ID = uuid.New()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{*user}
	walletRepo := &testutil.FakeUserWalletRepository{}
	custody := testutil.NewFakeCustodyService("0xcustody", "ext-123")
	granter := &fakeRoleGranter{}

	svc := NewProvisioningService(
		userRepo,
		walletRepo,
		custody,
		domain.CustodyProviderRaylsHSM,
		nil,
		granter,
		&testutil.StubLogger{},
	)

	err := svc.Provision(context.Background(), user)

	require.NoError(t, err)
	require.Equal(t, []string{"0xcustody"}, granter.granted)
	assert.Equal(t, []string{domain.RoleFactoryDeployer, domain.RolePrivacyNodeOperator}, granter.lastRoles)
	assert.Equal(t, domain.UserStatusRoleAssigned, user.Status)
}

func TestProvisioningService_Provision_GrantFailureDoesNotBlockProvisioning(t *testing.T) {
	// A FACTORY_DEPLOYER grant failure is best-effort: the wallet is still created and the user advances
	user := &domain.User{Status: domain.UserStatusWaitingRoleAssignment}
	user.ID = uuid.New()

	userRepo := testutil.NewFakeUserRepository()
	userRepo.Users = []domain.User{*user}
	walletRepo := &testutil.FakeUserWalletRepository{}
	custody := testutil.NewFakeCustodyService("0xcustody", "ext-123")
	granter := &fakeRoleGranter{err: errors.New("grantor lacks admin authority")}

	svc := NewProvisioningService(
		userRepo,
		walletRepo,
		custody,
		domain.CustodyProviderRaylsHSM,
		nil,
		granter,
		&testutil.StubLogger{},
	)

	err := svc.Provision(context.Background(), user)

	require.NoError(t, err)
	require.Len(t, walletRepo.Wallets, 1)
	assert.Equal(t, domain.UserStatusRoleAssigned, user.Status)
}

func TestProvisioningService_Provision_ReturnsCustodyErrorWhenWalletCreationFails(t *testing.T) {
	// If custody wallet creation fails, the error is propagated and no wallet is persisted
	user := &domain.User{Status: domain.UserStatusWaitingRoleAssignment}
	user.ID = uuid.New()

	walletRepo := &testutil.FakeUserWalletRepository{}
	custody := &testutil.FakeCustodyService{Err: errors.New("HSM unreachable")}

	svc := newProvisioningSvc(testutil.NewFakeUserRepository(), walletRepo, custody)

	err := svc.Provision(context.Background(), user)

	require.Error(t, err)
	assert.Empty(t, walletRepo.Wallets)
	assert.Equal(t, domain.UserStatusWaitingRoleAssignment, user.Status)
}
