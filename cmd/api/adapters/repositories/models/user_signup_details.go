package models

import (
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

type UserSignupDetails struct {
	Model
	UserID     uuid.UUID `gorm:"column:user_id;not null;uniqueIndex"`
	Company    string    `gorm:"not null;default:''"`
	Employees  string    `gorm:"not null;default:''"`
	HeardAbout string    `gorm:"column:heard_about;not null;default:''"`
	Goals      string    `gorm:"not null;default:''"`
}

func (UserSignupDetails) TableName() string { return "user_signup_details" }

func (d *UserSignupDetails) ToDomain() *domain.UserSignupDetails {
	return &domain.UserSignupDetails{
		Model:      d.ToDomainModel(),
		UserID:     d.UserID,
		Company:    d.Company,
		Employees:  d.Employees,
		HeardAbout: d.HeardAbout,
		Goals:      d.Goals,
	}
}

func UserSignupDetailsFromDomain(d *domain.UserSignupDetails) *UserSignupDetails {
	return &UserSignupDetails{
		Model:      ModelFromDomain(d.Model),
		UserID:     d.UserID,
		Company:    d.Company,
		Employees:  d.Employees,
		HeardAbout: d.HeardAbout,
		Goals:      d.Goals,
	}
}
