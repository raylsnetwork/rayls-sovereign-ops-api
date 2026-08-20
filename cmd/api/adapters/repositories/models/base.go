package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

type Model struct {
	ID        uuid.UUID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Model) BeforeCreate(_ *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	now := time.Now().UTC()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = now
	}
	return nil
}

func (m *Model) BeforeUpdate(_ *gorm.DB) error {
	m.UpdatedAt = time.Now().UTC()
	return nil
}

func (m *Model) ToDomainModel() domain.Model {
	return domain.Model{ID: m.ID, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt}
}

func ModelFromDomain(d domain.Model) Model {
	return Model{ID: d.ID, CreatedAt: d.CreatedAt, UpdatedAt: d.UpdatedAt}
}

type ImmutableModel struct {
	ID        uuid.UUID `gorm:"primarykey"`
	CreatedAt time.Time
}

func (m *ImmutableModel) BeforeCreate(_ *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now().UTC()
	}
	return nil
}

func (m *ImmutableModel) ToDomainImmutableModel() domain.ImmutableModel {
	return domain.ImmutableModel{ID: m.ID, CreatedAt: m.CreatedAt}
}

func ImmutableModelFromDomain(d domain.ImmutableModel) ImmutableModel {
	return ImmutableModel{ID: d.ID, CreatedAt: d.CreatedAt}
}
