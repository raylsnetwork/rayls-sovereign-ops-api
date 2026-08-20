package repositories

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/adapters/repositories/models"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

// ---------------------------------------------------------------------------
// AccessManagerRoleRepository
// ---------------------------------------------------------------------------

type accessManagerRoleRepository struct{ db *gorm.DB }

var _ core.AccessManagerRoleRepository = (*accessManagerRoleRepository)(nil)

func NewAccessManagerRoleRepository(db *gorm.DB) core.AccessManagerRoleRepository {
	return &accessManagerRoleRepository{db: db}
}

func (r *accessManagerRoleRepository) Upsert(ctx context.Context, role *domain.AccessManagerRole) error {
	role.UpdatedAt = time.Now().UTC()
	m := models.AMRoleFromDomain(role)
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "role_id"}},
			DoUpdates: clause.AssignmentColumns(
				[]string{"name", "label", "admin_role_id", "guardian_role_id", "grant_delay_secs", "updated_at"},
			),
		}).
		Create(m).Error
}

func (r *accessManagerRoleRepository) FindByID(ctx context.Context, roleID uint64) (*domain.AccessManagerRole, error) {
	var m models.AMRole
	if err := r.db.WithContext(ctx).Where("role_id = ?", roleID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

func (r *accessManagerRoleRepository) List(ctx context.Context) ([]*domain.AccessManagerRole, error) {
	var rows []models.AMRole
	if err := r.db.WithContext(ctx).Order("role_id asc").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*domain.AccessManagerRole, len(rows))
	for i := range rows {
		out[i] = rows[i].ToDomain()
	}
	return out, nil
}

// ---------------------------------------------------------------------------
// AccessManagerRoleMemberRepository
// ---------------------------------------------------------------------------

type accessManagerRoleMemberRepository struct{ db *gorm.DB }

var _ core.AccessManagerRoleMemberRepository = (*accessManagerRoleMemberRepository)(nil)

func NewAccessManagerRoleMemberRepository(db *gorm.DB) core.AccessManagerRoleMemberRepository {
	return &accessManagerRoleMemberRepository{db: db}
}

func (r *accessManagerRoleMemberRepository) Upsert(ctx context.Context, member *domain.AccessManagerRoleMember) error {
	member.UpdatedAt = time.Now().UTC()
	m := models.AMRoleMemberFromDomain(member)
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "role_id"}, {Name: "account"}},
			DoUpdates: clause.AssignmentColumns([]string{"execution_delay", "active_since", "is_active", "updated_at"}),
		}).
		Create(m).Error
}

func (r *accessManagerRoleMemberRepository) Revoke(ctx context.Context, roleID uint64, account string) error {
	return r.db.WithContext(ctx).
		Model(&models.AMRoleMember{}).
		Where("role_id = ? AND account = ?", roleID, account).
		Updates(map[string]any{"is_active": false, "updated_at": time.Now().UTC()}).Error
}

func (r *accessManagerRoleMemberRepository) ListByRole(
	ctx context.Context,
	roleID uint64,
	activeOnly bool,
) ([]*domain.AccessManagerRoleMember, error) {
	q := r.db.WithContext(ctx).Where("role_id = ?", roleID)
	if activeOnly {
		q = q.Where("is_active = true")
	}
	var rows []models.AMRoleMember
	if err := q.Order("account asc").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*domain.AccessManagerRoleMember, len(rows))
	for i := range rows {
		out[i] = rows[i].ToDomain()
	}
	return out, nil
}

func (r *accessManagerRoleMemberRepository) FindActiveAccountWithAllRoles(
	ctx context.Context,
	roleIDs []uint64,
) (string, error) {
	if len(roleIDs) == 0 {
		return "", core.ErrRecordNotFound
	}
	var account string
	err := r.db.WithContext(ctx).
		Model(&models.AMRoleMember{}).
		Where("role_id IN ? AND is_active = true", roleIDs).
		Group("account").
		Having("COUNT(DISTINCT role_id) = ?", len(roleIDs)).
		Order("account asc").
		Limit(1).
		Pluck("account", &account).Error
	if err != nil {
		return "", err
	}
	if account == "" {
		return "", core.ErrRecordNotFound
	}
	return account, nil
}

func (r *accessManagerRoleMemberRepository) ListByAccount(
	ctx context.Context,
	account string,
) ([]*domain.AccessManagerRoleMember, error) {
	var rows []models.AMRoleMember
	if err := r.db.WithContext(ctx).Where("account = ?", account).Order("role_id asc").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*domain.AccessManagerRoleMember, len(rows))
	for i := range rows {
		out[i] = rows[i].ToDomain()
	}
	return out, nil
}

// ---------------------------------------------------------------------------
// AccessManagerManagedContractRepository
// ---------------------------------------------------------------------------

type accessManagerManagedContractRepository struct{ db *gorm.DB }

var _ core.AccessManagerManagedContractRepository = (*accessManagerManagedContractRepository)(nil)

func NewAccessManagerManagedContractRepository(db *gorm.DB) core.AccessManagerManagedContractRepository {
	return &accessManagerManagedContractRepository{db: db}
}

func (r *accessManagerManagedContractRepository) Upsert(
	ctx context.Context,
	c *domain.AccessManagerManagedContract,
) error {
	c.UpdatedAt = time.Now().UTC()
	if c.RegisteredAt.IsZero() {
		c.RegisteredAt = c.UpdatedAt
	}
	m := models.AMManagedContractFromDomain(c)
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "contract_address"}},
			DoUpdates: clause.AssignmentColumns([]string{"deployer", "is_paused", "updated_at"}),
		}).
		Create(m).Error
}

func (r *accessManagerManagedContractRepository) SetPaused(
	ctx context.Context,
	contractAddress string,
	paused bool,
) error {
	return r.db.WithContext(ctx).
		Model(&models.AMManagedContract{}).
		Where("contract_address = ?", contractAddress).
		Updates(map[string]any{"is_paused": paused, "updated_at": time.Now().UTC()}).Error
}

func (r *accessManagerManagedContractRepository) List(
	ctx context.Context,
) ([]*domain.AccessManagerManagedContract, error) {
	var rows []models.AMManagedContract
	if err := r.db.WithContext(ctx).Order("registered_at asc").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*domain.AccessManagerManagedContract, len(rows))
	for i := range rows {
		out[i] = rows[i].ToDomain()
	}
	return out, nil
}

func (r *accessManagerManagedContractRepository) FindByAddress(
	ctx context.Context,
	contractAddress string,
) (*domain.AccessManagerManagedContract, error) {
	var m models.AMManagedContract
	if err := r.db.WithContext(ctx).Where("contract_address = ?", contractAddress).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

// ---------------------------------------------------------------------------
// AccessManagerFunctionPermissionRepository
// ---------------------------------------------------------------------------

type accessManagerFunctionPermissionRepository struct{ db *gorm.DB }

var _ core.AccessManagerFunctionPermissionRepository = (*accessManagerFunctionPermissionRepository)(nil)

func NewAccessManagerFunctionPermissionRepository(db *gorm.DB) core.AccessManagerFunctionPermissionRepository {
	return &accessManagerFunctionPermissionRepository{db: db}
}

func (r *accessManagerFunctionPermissionRepository) Upsert(
	ctx context.Context,
	perm *domain.AccessManagerFunctionPermission,
) error {
	perm.UpdatedAt = time.Now().UTC()
	m := models.AMFunctionPermissionFromDomain(perm)
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "contract_address"}, {Name: "selector"}, {Name: "role_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"is_active", "updated_at"}),
		}).
		Create(m).Error
}

func (r *accessManagerFunctionPermissionRepository) Remove(
	ctx context.Context,
	contractAddress, selector string,
	roleID uint64,
) error {
	return r.db.WithContext(ctx).
		Model(&models.AMFunctionPermission{}).
		Where("contract_address = ? AND selector = ? AND role_id = ?", contractAddress, selector, roleID).
		Updates(map[string]any{"is_active": false, "updated_at": time.Now().UTC()}).Error
}

func (r *accessManagerFunctionPermissionRepository) ListByContract(
	ctx context.Context,
	contractAddress string,
) ([]*domain.AccessManagerFunctionPermission, error) {
	var rows []models.AMFunctionPermission
	if err := r.db.WithContext(ctx).
		Where("contract_address = ?", contractAddress).
		Order("selector asc, role_id asc").
		Find(&rows).
		Error; err != nil {
		return nil, err
	}
	out := make([]*domain.AccessManagerFunctionPermission, len(rows))
	for i := range rows {
		out[i] = rows[i].ToDomain()
	}
	return out, nil
}

// ---------------------------------------------------------------------------
// AccessManagerScheduledOperationRepository
// ---------------------------------------------------------------------------

type accessManagerScheduledOperationRepository struct{ db *gorm.DB }

var _ core.AccessManagerScheduledOperationRepository = (*accessManagerScheduledOperationRepository)(nil)

func NewAccessManagerScheduledOperationRepository(db *gorm.DB) core.AccessManagerScheduledOperationRepository {
	return &accessManagerScheduledOperationRepository{db: db}
}

func (r *accessManagerScheduledOperationRepository) Upsert(
	ctx context.Context,
	op *domain.AccessManagerScheduledOperation,
) error {
	op.UpdatedAt = time.Now().UTC()
	m := models.AMScheduledOperationFromDomain(op)
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "operation_id"}},
			DoUpdates: clause.AssignmentColumns(
				[]string{"caller", "managed_contract", "execute_after", "status", "updated_at"},
			),
		}).
		Create(m).Error
}

func (r *accessManagerScheduledOperationRepository) UpdateStatus(
	ctx context.Context,
	operationID, status string,
) error {
	return r.db.WithContext(ctx).
		Model(&models.AMScheduledOperation{}).
		Where("operation_id = ?", operationID).
		Updates(map[string]any{"status": status, "updated_at": time.Now().UTC()}).Error
}

func (r *accessManagerScheduledOperationRepository) List(
	ctx context.Context,
	status string,
) ([]*domain.AccessManagerScheduledOperation, error) {
	q := r.db.WithContext(ctx)
	if status != "" {
		q = q.Where("status = ?", status)
	}
	var rows []models.AMScheduledOperation
	if err := q.Order("updated_at desc").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*domain.AccessManagerScheduledOperation, len(rows))
	for i := range rows {
		out[i] = rows[i].ToDomain()
	}
	return out, nil
}

// ---------------------------------------------------------------------------
// AccessManagerContractScopedRoleMemberRepository
// ---------------------------------------------------------------------------

type accessManagerContractScopedRoleMemberRepository struct{ db *gorm.DB }

var _ core.AccessManagerContractScopedRoleMemberRepository = (*accessManagerContractScopedRoleMemberRepository)(nil)

func NewAccessManagerContractScopedRoleMemberRepository(
	db *gorm.DB,
) core.AccessManagerContractScopedRoleMemberRepository {
	return &accessManagerContractScopedRoleMemberRepository{db: db}
}

func (r *accessManagerContractScopedRoleMemberRepository) Upsert(
	ctx context.Context,
	m *domain.AccessManagerContractScopedRoleMember,
) error {
	m.UpdatedAt = time.Now().UTC()
	row := models.AMContractScopedRoleMemberFromDomain(m)
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "role_id"}, {Name: "account"}, {Name: "contract_address"}},
			DoUpdates: clause.AssignmentColumns([]string{"execution_delay", "active_since", "is_active", "updated_at"}),
		}).
		Create(row).Error
}

func (r *accessManagerContractScopedRoleMemberRepository) Revoke(
	ctx context.Context,
	roleID uint64,
	account, contractAddress string,
) error {
	return r.db.WithContext(ctx).
		Model(&models.AMContractScopedRoleMember{}).
		Where("role_id = ? AND account = ? AND contract_address = ?", roleID, account, contractAddress).
		Updates(map[string]any{"is_active": false, "updated_at": time.Now().UTC()}).Error
}

func (r *accessManagerContractScopedRoleMemberRepository) ListByAccount(
	ctx context.Context,
	account string,
) ([]*domain.AccessManagerContractScopedRoleMember, error) {
	var rows []models.AMContractScopedRoleMember
	if err := r.db.WithContext(ctx).Where("account = ?", account).Order("role_id asc").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*domain.AccessManagerContractScopedRoleMember, len(rows))
	for i := range rows {
		out[i] = rows[i].ToDomain()
	}
	return out, nil
}

// ---------------------------------------------------------------------------
// AccessManagerEventLogRepository
// ---------------------------------------------------------------------------

type accessManagerEventLogRepository struct{ db *gorm.DB }

var _ core.AccessManagerEventLogRepository = (*accessManagerEventLogRepository)(nil)

func NewAccessManagerEventLogRepository(db *gorm.DB) core.AccessManagerEventLogRepository {
	return &accessManagerEventLogRepository{db: db}
}

func (r *accessManagerEventLogRepository) Insert(ctx context.Context, entry *domain.AccessManagerEventLog) error {
	m := models.AMEventLogFromDomain(entry)
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "tx_hash"}, {Name: "log_index"}},
			DoNothing: true,
		}).
		Create(m).Error
}

func (r *accessManagerEventLogRepository) List(
	ctx context.Context,
	eventName string,
	page, limit int,
) ([]*domain.AccessManagerEventLog, int64, error) {
	q := r.db.WithContext(ctx).Model(&models.AMEventLog{})
	if eventName != "" {
		q = q.Where("event_name = ?", eventName)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if limit <= 0 {
		limit = 20
	}
	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}
	var rows []models.AMEventLog
	if err := q.Order("block_number desc, log_index desc").Limit(limit).Offset(offset).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	out := make([]*domain.AccessManagerEventLog, len(rows))
	for i := range rows {
		out[i] = rows[i].ToDomain()
	}
	return out, total, nil
}
