package custody

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
)

var _ core.CustodyService = (*SelfCustody)(nil)

// SelfCustody is a no-op implementation for users who bring their own wallet (SIWE).
// It should never be called in practice — SIWE users have their address stored in
// user_wallets at registration time, so provisioning skips wallet creation entirely.
type SelfCustody struct{}

func NewSelfCustody() *SelfCustody { return &SelfCustody{} }

func (s *SelfCustody) CreateWallet(_ context.Context, userID uuid.UUID) (string, string, error) {
	return "", "", fmt.Errorf("self-custody does not generate wallets: user %s must provide their own address", userID)
}

func (s *SelfCustody) SignAndTransact(_ context.Context, _ []byte, _ string, _ string) (string, error) {
	return "", fmt.Errorf("self-custody does not support transaction signing")
}
