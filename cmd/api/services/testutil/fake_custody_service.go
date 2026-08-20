package testutil

import (
	"context"

	"github.com/google/uuid"
)

// FakeCustodyService is an in-memory custody service for use in tests.
type FakeCustodyService struct {
	Address    string
	ExternalID string
	Err        error
	// Calls counts mints — the real HSM generates a fresh random key per call, so a test that
	// asserts on orphan behaviour needs to know how many keys were minted.
	Calls int
	// OnCreate runs at mint time, before the address is returned. Lets a test observe what the
	// database looked like at the moment of the irreversible side effect.
	OnCreate func()
}

func NewFakeCustodyService(address, externalID string) *FakeCustodyService {
	return &FakeCustodyService{Address: address, ExternalID: externalID}
}

func (f *FakeCustodyService) CreateWallet(_ context.Context, _ uuid.UUID) (string, string, error) {
	f.Calls++
	if f.OnCreate != nil {
		f.OnCreate()
	}
	return f.Address, f.ExternalID, f.Err
}

func (f *FakeCustodyService) SignAndTransact(_ context.Context, _ []byte, _ string, _ string) (string, error) {
	return "", f.Err
}
