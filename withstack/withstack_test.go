package withstack

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrap_PreservesOriginalInUnwrapChain(t *testing.T) {
	// Wrapping an error preserves the original so errors.Is still recognizes it.
	orig := errors.New("boom")

	w := Wrap(orig)

	require.NotEmpty(t, w.Error())
	assert.True(t, errors.Is(w, orig))
}

func TestWrap_NilReturnsNil(t *testing.T) {
	// Wrapping a nil error returns a nil wrapper rather than a non-nil empty one.
	assert.Nil(t, Wrap(nil))
}
