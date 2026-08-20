package testutil

import "context"

// RollbackTransactor is a FakeTransactor that also undoes writes when the closure fails,
// which plain FakeTransactor cannot do — it runs fn against the same in-memory slices
// whether or not fn succeeds, so a test using it cannot tell "rolled back" from "committed".
//
// Snapshot/Restore are supplied by the caller because the fakes hold unrelated slice types.
// The usual shape is a closure per repository capturing its own backing slice:
//
//	tx := &RollbackTransactor{
//	    Snapshot: func() { savedUsers = append([]domain.User(nil), users.Users...) },
//	    Restore:  func() { users.Users = savedUsers },
//	}
type RollbackTransactor struct {
	Snapshot func()
	Restore  func()
	Rollback bool // set to true when a failing closure caused Restore to run
}

func (t *RollbackTransactor) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	if t.Snapshot != nil {
		t.Snapshot()
	}

	if err := fn(ctx); err != nil {
		t.Rollback = true
		if t.Restore != nil {
			t.Restore()
		}
		return err
	}
	return nil
}
