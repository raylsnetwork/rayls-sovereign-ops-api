package models

import "github.com/raylsnetwork/rayls-sovereign-ops-api/domain"

type Token struct {
	Model
	Name            string  `gorm:"type:text"`
	Symbol          string  `gorm:"type:text"`
	ResourceID      *string `gorm:"uniqueIndex"`
	MetadataURL     string  `gorm:"type:text"`
	ErcStandard     int16
	Decimals        int16
	IssuerID        string `gorm:"type:text"`
	Status          int16
	ContractAddress string `gorm:"uniqueIndex;size:42;not null"`
	TokenClass      string `gorm:"size:50;not null;default:unknown"`
	TotalSupply     string `gorm:"type:text;not null;default:''"`
	HolderCount     int    `gorm:"not null;default:0"`
}

func (Token) TableName() string { return "tokens" }

func (t *Token) ToDomain() *domain.Token {
	return &domain.Token{
		Model:           t.ToDomainModel(),
		Name:            t.Name,
		Symbol:          t.Symbol,
		ResourceID:      t.ResourceID,
		MetadataURL:     t.MetadataURL,
		ErcStandard:     domain.ErcStandard(t.ErcStandard),
		Decimals:        uint8(t.Decimals),
		IssuerID:        t.IssuerID,
		Status:          domain.TokenStatus(t.Status),
		ContractAddress: t.ContractAddress,
		TokenClass:      t.TokenClass,
		TotalSupply:     t.TotalSupply,
		HolderCount:     t.HolderCount,
	}
}

func TokenFromDomain(d *domain.Token) *Token {
	return &Token{
		Model:           ModelFromDomain(d.Model),
		Name:            d.Name,
		Symbol:          d.Symbol,
		ResourceID:      d.ResourceID,
		MetadataURL:     d.MetadataURL,
		ErcStandard:     int16(d.ErcStandard),
		Decimals:        int16(d.Decimals),
		IssuerID:        d.IssuerID,
		Status:          int16(d.Status),
		ContractAddress: d.ContractAddress,
		TokenClass:      d.TokenClass,
		TotalSupply:     d.TotalSupply,
		HolderCount:     d.HolderCount,
	}
}
