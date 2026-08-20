package testutil

import (
	"context"
)

// FakeRaylsAccessManagerClient is an in-memory RaylsAccessManagerClient for unit tests.
type FakeRaylsAccessManagerClient struct {
	Roles []string
	Err   error
	// CallCount records how many times GetRoles was invoked, so a test can assert that a code
	// path bypassed the AccessManager entirely rather than inferring it from a return value
	// the fake would have produced anyway.
	CallCount int
}

func (f *FakeRaylsAccessManagerClient) GetRoles(_ context.Context, _ string) ([]string, error) {
	f.CallCount++
	if f.Err != nil {
		return nil, f.Err
	}
	return f.Roles, nil
}
