package testutil

import "context"

// FakeTransactor executes fn directly without a real DB transaction.
type FakeTransactor struct {
	Err error // if set, returned instead of executing fn
}

func (t *FakeTransactor) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	if t.Err != nil {
		return t.Err
	}
	return fn(ctx)
}
