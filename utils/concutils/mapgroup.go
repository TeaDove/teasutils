package concutils

import (
	"errors"
	"sync"
)

// MapGroup runs functions concurrently with a bounded degree of parallelism
// and gathers their (key, value) results into a map. The zero value is not
// usable; create one with NewMapGroup.
type MapGroup[K comparable, T any] struct {
	mapping map[K]T
	sem     sem
	wg      sync.WaitGroup
	mu      sync.Mutex
	err     error
}

// NewMapGroup returns a MapGroup whose result map is pre-sized to mapCap and
// which runs at most limit functions concurrently. It panics if limit <= 0.
func NewMapGroup[K comparable, T any](mapCap int, limit int) *MapGroup[K, T] {
	return &MapGroup[K, T]{mapping: make(map[K]T, mapCap), sem: newSem(limit)}
}

// Go schedules fn to run in its own goroutine, blocking the caller until a
// concurrency slot is free. On success the returned key/value is stored; on
// error the value is dropped and the error is accumulated (see Collect).
// A later write to the same key overwrites the earlier one.
func (r *MapGroup[K, T]) Go(fn func() (K, T, error)) {
	r.sem.lock()

	r.wg.Go(func() {
		defer r.sem.unlock()

		k, t, err := fn()

		r.mu.Lock()
		defer r.mu.Unlock()

		if err != nil {
			r.err = errors.Join(r.err, err)

			return
		}

		r.mapping[k] = t
	})
}

// Collect blocks until all scheduled functions have returned, then reports
// the accumulated map together with the joined error of every failed call
// (nil if all succeeded). The returned map is the group's own map: treat it
// as owned by the caller and do not schedule more work after Collect.
func (r *MapGroup[K, T]) Collect() (map[K]T, error) {
	r.wg.Wait()

	return r.mapping, r.err
}
