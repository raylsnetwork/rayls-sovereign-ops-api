package domain

// TokenEvent is an immutable audit record of an on-chain status change event.
// It embeds ImmutableModel instead of Model because token_events has no updated_at column.
// Payload serialization to JSONB is handled by the GORM model in adapters/repositories/models.
type TokenEvent struct {
	ImmutableModel
	ContractAddress string         `json:"contractAddress"`
	EventType       string         `json:"eventType"`
	BlockNumber     int64          `json:"blockNumber"`
	TxHash          string         `json:"txHash"`
	Payload         map[string]any `json:"payload"`
}
