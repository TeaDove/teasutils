package concutils

import (
	"runtime"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/teadove/teasutils/utils/rutils"
)

func TestExecuteGroup(t *testing.T) {
	t.Parallel()

	exGroup := NewExecuteGroup[rutils.Pair[int, int]](5)
	for i := range 10 {
		exGroup.Go(func() (rutils.Pair[int, int], error) {
			return rutils.NewPair(i, i), nil
		})
	}

	results, err := exGroup.IntoSlice()
	require.NoError(t, err)
	assert.Len(t, results, 10)

	for _, result := range results {
		assert.Equal(t, result.Second, result.First)
	}
}

func TestExecuteGroup_IntoSlice_ReturnsError(t *testing.T) {
	t.Parallel()

	sentinel := errors.New("boom")

	exGroup := NewExecuteGroup[int](3)
	for i := range 10 {
		exGroup.Go(func() (int, error) {
			if i == 7 {
				return 0, sentinel
			}

			return i, nil
		})
	}

	results, err := exGroup.IntoSlice()
	require.Error(t, err)
	require.ErrorIs(t, err, sentinel)
	assert.Nil(t, results)
}

// TestExecuteGroup_IntoSlice_NoLeakOnError guards against workers blocking
// forever on send when IntoSlice returns early on the first error.
func TestExecuteGroup_IntoSlice_NoLeakOnError(t *testing.T) {
	t.Parallel()

	before := runtime.NumGoroutine()

	// Many tasks, small limit: with an unbuffered results channel most
	// workers are still pending when the error short-circuits IntoSlice.
	exGroup := NewExecuteGroup[int](2)
	for i := range 100 {
		exGroup.Go(func() (int, error) {
			if i == 0 {
				return 0, errors.New("boom")
			}

			return i, nil
		})
	}

	_, err := exGroup.IntoSlice()
	require.Error(t, err)

	assert.Eventually(t, func() bool {
		return runtime.NumGoroutine() <= before+2
	}, time.Second, 10*time.Millisecond, "worker goroutines must not leak")
}
