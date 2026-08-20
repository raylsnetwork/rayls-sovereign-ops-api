package domain

import "time"

type Nonce struct {
	Model
	WalletAddress string    `json:"walletAddress"`
	Nonce         string    `json:"nonce"`
	Message       string    `json:"message"`
	ExpiresAt     time.Time `json:"expiresAt"`
	Used          bool      `json:"used"`
}
