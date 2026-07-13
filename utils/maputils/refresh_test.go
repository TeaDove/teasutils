package maputils

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRefreshingAtomicMap_InitError(t *testing.T) {
	t.Parallel()

	_, err := NewRefreshingAtomicMap(t.Context(), time.Hour, func() (map[string]string, error) {
		return nil, errors.New("boom")
	})
	require.Error(t, err)
}

func TestRefreshingAtomicMap_RefreshUpdatesValues(t *testing.T) {
	t.Parallel()

	var gen atomic.Int64

	m, err := NewRefreshingAtomicMap(t.Context(), time.Hour, func() (map[string]int, error) {
		v := int(gen.Add(1))

		return map[string]int{"k": v}, nil
	})
	require.NoError(t, err)

	got, ok := m.Get("k")
	require.True(t, ok)
	assert.Equal(t, 1, got)

	require.NoError(t, m.Refresh(t.Context()))

	got, ok = m.Get("k")
	require.True(t, ok)
	assert.Equal(t, 2, got, "Get must observe the refreshed value")
}

func TestRefreshingAtomicMap_AutoRefreshTicksAndStops(t *testing.T) {
	t.Parallel()

	var calls atomic.Int64

	ctx, cancel := context.WithCancel(t.Context())

	_, err := NewRefreshingAtomicMap(ctx, 20*time.Millisecond, func() (map[string]int, error) {
		calls.Add(1)

		return map[string]int{}, nil
	})
	require.NoError(t, err) // one call for the initial refresh

	assert.Eventually(t, func() bool {
		return calls.Load() >= 3
	}, time.Second, 5*time.Millisecond, "auto-refresh must keep ticking")

	cancel()

	time.Sleep(60 * time.Millisecond)

	stopped := calls.Load()

	time.Sleep(60 * time.Millisecond)
	assert.Equal(t, stopped, calls.Load(), "auto-refresh must stop after ctx is cancelled")
}
