package domain

import (
	"encoding/hex"
	"strings"

	"github.com/google/uuid"
)

type CustodyProviderType int

const (
	CustodyProviderUnknown  CustodyProviderType = iota // 0 — reserved/unset
	CustodyProviderRaylsHSM                            // 1
	// CustodyProviderSelf represents a self-custodied wallet (e.g. SIWE/MetaMask).
	// Used as a placeholder until custody integration (Phase 7) is implemented.
	CustodyProviderSelf // 2
)

func (c CustodyProviderType) IsValid() bool {
	switch c {
	case CustodyProviderRaylsHSM, CustodyProviderSelf:
		return true
	}
	return false
}

// WalletChain discriminates which chain a wallet operates on. It is a discriminator only:
// a user accumulates one private-chain and one public-chain wallet per onboarding call, so it
// is NOT unique per (user_id, chain). Global wallet-address uniqueness is the only invariant.
type WalletChain int

const (
	WalletChainPrivate WalletChain = 1 // privacy node; existing default
	WalletChainPublic  WalletChain = 2 // public chain
)

func (c WalletChain) IsValid() bool {
	switch c {
	case WalletChainPrivate, WalletChainPublic:
		return true
	}
	return false
}

type UserWallet struct {
	Model
	UserID            uuid.UUID           `json:"userId"`
	RaylsAddress      string              `json:"raylsAddress"`
	CustodyProvider   CustodyProviderType `json:"custodyProvider"`
	CustodyExternalID string              `json:"custodyExternalId"`
	Chain             WalletChain         `json:"chain"`
	IsActive          bool                `json:"isActive"`
}

// pendingAddressPrefix marks a wallet row that exists only to record the *intent* to mint.
// The HSM has no user id, no idempotency key and no delete API, so a crash between minting
// and persisting used to orphan the key permanently and — because the retry keyed on a row
// that was never written — mint a fresh orphan on every attempt. Writing the intent first
// makes the database authoritative before the side effect: a stranded row names the user,
// so recovery becomes a lookup instead of a guess.
//
// The placeholder occupies rayls_address (NOT NULL, unique on LOWER(rayls_address)) until
// the real address is known. It is per-attempt unique, and is_active=false keeps the row
// invisible to every existing query — none of which needs a migration to stay correct.
const pendingAddressPrefix = "pending:"

// PendingAddressPattern matches pending placeholders in a SQL LIKE. A real EVM address always
// starts "0x", so it can never match.
const PendingAddressPattern = pendingAddressPrefix + "%"

// PendingAddress builds the placeholder address for an in-flight mint. It is derived from a
// per-attempt UUID so two concurrent mints for the same user cannot collide on the unique
// index. Fits VARCHAR(42): 8 + 32 = 40 characters.
func PendingAddress(attempt uuid.UUID) string {
	return pendingAddressPrefix + hex.EncodeToString(attempt[:])
}

// IsPending reports whether the wallet is an unfulfilled mint intent rather than a real
// wallet. Such a row never has a usable address and must never be returned as a signer.
func (w *UserWallet) IsPending() bool {
	return strings.HasPrefix(w.RaylsAddress, pendingAddressPrefix)
}
