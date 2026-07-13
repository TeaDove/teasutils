package mustutils

import (
	"context"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithRecover_Success(t *testing.T) {
	t.Parallel()

	require.NoError(t, WithRecover(func() error { return nil })())
}

func TestWithRecover_ReturnsFnError(t *testing.T) {
	t.Parallel()

	sentinel := errors.New("boom")
	err := WithRecover(func() error { return sentinel })()
	assert.ErrorIs(t, err, sentinel)
}

func TestWithRecover_RecoversPanic(t *testing.T) {
	t.Parallel()

	err := WithRecover(func() error { panic("bad") })()
	require.Error(t, err)
	require.ErrorContains(t, err, "recovered")
	require.ErrorContains(t, err, "panicked")
}

// TestDoInBackground_RunsDetachedFromParentCancel verifies fn still runs even
// though the parent context is already cancelled (DoInBackground detaches it).
func TestDoInBackground_RunsDetachedFromParentCancel(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	done := make(chan struct{})

	DoInBackground(ctx, func(_ context.Context) error {
		close(done)

		return nil
	})

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("fn was not executed")
	}
}

// TestDoInBackground_RecoversPanic ensures a panic in fn does not crash the
// process; the goroutine recovers and logs instead.
func TestDoInBackground_RecoversPanic(t *testing.T) {
	t.Parallel()

	started := make(chan struct{})

	DoInBackground(t.Context(), func(_ context.Context) error {
		close(started)
		panic("bad")
	})

	<-started
	// Give the deferred recover time to run; a crash here would fail the test.
	time.Sleep(50 * time.Millisecond)
}
