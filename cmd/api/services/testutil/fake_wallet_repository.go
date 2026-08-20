package testutil

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

// FakeUserWalletRepository is an in-memory UserWalletRepository for unit tests.
type FakeUserWalletRepository struct {
	Wallets               []domain.UserWallet
	FindByRaylsAddressErr error // if set, returned by FindByRaylsAddress instead of searching
	CreateErr             error // if set, returned by Create instead of appending
	CompletePendingErr    error // if set, returned by CompletePending instead of updating
}

func (r *FakeUserWalletRepository) Create(_ context.Context, wallet *domain.UserWallet) error {
	if r.CreateErr != nil {
		return r.CreateErr
	}
	if wallet.ID == uuid.Nil {
		wallet.ID = uuid.New()
	}
	r.Wallets = append(r.Wallets, *wallet)
	return nil
}

// CompletePending mirrors the repository: it only ever touches a row that is still pending,
// so a retry cannot overwrite a completed wallet.
func (r *FakeUserWalletRepository) CompletePending(
	_ context.Context,
	id uuid.UUID,
	address, externalID string,
) error {
	if r.CompletePendingErr != nil {
		return r.CompletePendingErr
	}
	for i := range r.Wallets {
		w := &r.Wallets[i]
		if w.ID != id || !w.IsPending() {
			continue
		}
		for j := range r.Wallets {
			if j != i && strings.EqualFold(r.Wallets[j].RaylsAddress, address) {
				return core.ErrDuplicateWalletAddress
			}
		}
		w.RaylsAddress = address
		w.CustodyExternalID = externalID
		w.IsActive = true
		return nil
	}
	return core.ErrRecordNotFound
}

func (r *FakeUserWalletRepository) DeletePending(_ context.Context, id uuid.UUID) error {
	for i := range r.Wallets {
		if r.Wallets[i].ID == id && r.Wallets[i].IsPending() {
			r.Wallets = append(r.Wallets[:i], r.Wallets[i+1:]...)
			return nil
		}
	}
	return nil
}

func (r *FakeUserWalletRepository) FindPendingByUserID(
	_ context.Context,
	userID uuid.UUID,
) ([]domain.UserWallet, error) {
	var pending []domain.UserWallet
	for i := range r.Wallets {
		if r.Wallets[i].UserID == userID && r.Wallets[i].IsPending() {
			pending = append(pending, r.Wallets[i])
		}
	}
	return pending, nil
}

func (r *FakeUserWalletRepository) FindByUserID(_ context.Context, userID uuid.UUID) (*domain.UserWallet, error) {
	for i := range r.Wallets {
		// Pending intents are is_active=false in the real schema, so they are invisible here
		// too — a stranded intent must never be returned as if it were a usable wallet.
		if r.Wallets[i].UserID == userID && r.Wallets[i].IsActive && !r.Wallets[i].IsPending() {
			return &r.Wallets[i], nil
		}
	}
	return nil, core.ErrRecordNotFound
}

func (r *FakeUserWalletRepository) GetSignerWalletForChain(
	_ context.Context,
	userID uuid.UUID,
	chain domain.WalletChain,
) (*domain.UserWallet, error) {
	for i := range r.Wallets {
		w := &r.Wallets[i]
		if w.UserID != userID || !w.IsActive || w.CustodyProvider != domain.CustodyProviderRaylsHSM {
			continue
		}
		// Mirror the DB default (column defaults to 1) for fixtures that leave Chain unset.
		walletChain := w.Chain
		if walletChain == 0 {
			walletChain = domain.WalletChainPrivate
		}
		if walletChain == chain {
			return w, nil
		}
	}
	return nil, core.ErrRecordNotFound
}

func (r *FakeUserWalletRepository) GetSignerWalletByAddress(
	_ context.Context,
	userID uuid.UUID,
	address string,
) (*domain.UserWallet, error) {
	for i := range r.Wallets {
		w := &r.Wallets[i]
		if w.UserID == userID && w.IsActive && w.CustodyProvider == domain.CustodyProviderRaylsHSM &&
			strings.EqualFold(w.RaylsAddress, address) {
			return w, nil
		}
	}
	return nil, core.ErrRecordNotFound
}

func (r *FakeUserWalletRepository) FindByRaylsAddress(_ context.Context, address string) (*domain.UserWallet, error) {
	if r.FindByRaylsAddressErr != nil {
		return nil, r.FindByRaylsAddressErr
	}
	for i := range r.Wallets {
		if strings.EqualFold(r.Wallets[i].RaylsAddress, address) {
			return &r.Wallets[i], nil
		}
	}
	return nil, core.ErrRecordNotFound
}
