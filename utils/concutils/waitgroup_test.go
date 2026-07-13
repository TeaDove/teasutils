package concutils

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWaitGroupWithLimit_RunsAllAndCapsConcurrency(t *testing.T) {
	t.Parallel()

	const (
		limit = 3
		tasks = 30
	)

	var (
		done     atomic.Int64
		inFlight atomic.Int64
		maxSeen  atomic.Int64
	)

	wg := NewSemWaitGroup(limit)
	for range tasks {
		wg.Go(func() {
			cur := inFlight.Add(1)

			for {
				old := maxSeen.Load()
				if cur <= old || maxSeen.CompareAndSwap(old, cur) {
					break
				}
			}

			time.Sleep(time.Millisecond)
			done.Add(1)
			inFlight.Add(-1)
		})
	}

	wg.Wait()

	assert.Equal(t, int64(tasks), done.Load(), "every task must run")
	assert.LessOrEqual(t, maxSeen.Load(), int64(limit), "concurrency must never exceed limit")
}

func TestNewSemWaitGroup_PanicsOnNonPositiveLimit(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() { NewSemWaitGroup(0) })
	assert.Panics(t, func() { NewSemWaitGroup(-1) })
}
