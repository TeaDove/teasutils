package concutils

import (
	"sync"
)

// WaitGroupWithLimit is a sync.WaitGroup that caps the number of
// simultaneously running goroutines. The zero value is not usable;
// create one with NewSemWaitGroup.
type WaitGroupWithLimit struct {
	sem sem
	wg  sync.WaitGroup
}

// NewSemWaitGroup returns a WaitGroupWithLimit that runs at most limit
// goroutines concurrently. It panics if limit <= 0.
func NewSemWaitGroup(limit int) *WaitGroupWithLimit {
	return &WaitGroupWithLimit{sem: newSem(limit)}
}

// Go schedules fn to run in its own goroutine, blocking the caller until
// a concurrency slot is free (at most limit are in flight at once).
// A panic in fn propagates on its own goroutine and is not recovered.
func (r *WaitGroupWithLimit) Go(fn func()) {
	r.sem.lock()
	r.wg.Go(func() {
		defer r.sem.unlock()

		fn()
	})
}

// Wait blocks until all scheduled functions have returned.
func (r *WaitGroupWithLimit) Wait() {
	r.wg.Wait()
}
