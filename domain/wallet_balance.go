package domain

type WalletBalance struct {
	Model
	WalletAddress string `json:"walletAddress"`
	TokenAddress  string `json:"tokenAddress"`
	Balance       string `json:"balance"`
	BlockNumber   uint64 `json:"blockNumber"`
}
