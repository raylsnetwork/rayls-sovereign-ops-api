package domain

import "time"

type TokenBlacklist struct {
	ImmutableModel
	JTI       string    `json:"jti"`
	ExpiresAt time.Time `json:"expiresAt"`
}
