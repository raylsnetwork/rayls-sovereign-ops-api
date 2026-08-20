package custody

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelfCustody_CreateWallet_AlwaysErrors(t *testing.T) {
	// Self-custody never generates wallets: SIWE users provide their own address at registration.
	s := NewSelfCustody()

	addr, priv, err := s.CreateWallet(context.Background(), uuid.New())

	require.Error(t, err)
	assert.Empty(t, addr)
	assert.Empty(t, priv)
}

func TestSelfCustody_SignAndTransact_AlwaysErrors(t *testing.T) {
	// Self-custody holds no key material, so it must refuse to sign rather than silently succeed.
	s := NewSelfCustody()

	txHash, err := s.SignAndTransact(context.Background(), []byte{0x01}, "0xsigner", "1")

	require.Error(t, err)
	assert.Empty(t, txHash)
}
