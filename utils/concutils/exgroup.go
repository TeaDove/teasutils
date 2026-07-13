package concutils

import (
	"sync"

	"github.com/teadove/teasutils/utils/rutils"
)

// ExecuteGroup runs fallible functions concurrently with a bounded degree
// of parallelism and streams their results back as rutils.Result values.
// The zero value is not usable; create one with NewExecuteGroup.
type ExecuteGroup[T any] struct {
	wg      sync.WaitGroup
	sem     sem
	results chan rutils.Result[T]
}

// NewExecuteGroup returns an ExecuteGroup that runs at most limit
// functions concurrently. It panics if limit <= 0.
func NewExecuteGroup[T any](limit int) *ExecuteGroup[T] {
	return &ExecuteGroup[T]{results: make(chan rutils.Result[T]), sem: newSem(limit)}
}

// Go schedules fn to run in its own goroutine, blocking the caller until a
// concurrency slot is free. The (value, error) returned by fn is delivered
// on the Results channel, so every Go call must eventually be matched by a
// receive (via Results or IntoSlice) or the worker goroutine will block.
// Do not call Go after Results/IntoSlice has been called.
func (r *ExecuteGroup[T]) Go(fn func() (T, error)) {
	r.sem.lock()
	r.wg.Go(func() {
		result := rutils.NewResult(fn())

		r.sem.unlock()

		r.results <- result
	})
}

// Results returns a channel that yields one result per scheduled function
// and is closed once every worker has finished. The caller must drain it
// fully, otherwise unfinished workers leak on their send.
//
// Contract: Results must be called at most once per ExecuteGroup. Calling
// it again panics ("close of closed channel") when the second closer runs.
func (r *ExecuteGroup[T]) Results() <-chan rutils.Result[T] {
	go func() {
		r.wg.Wait()
		close(r.results)
	}()

	return r.results
}

// IntoSlice collects every result, returning the values in completion order.
// It returns the first error encountered, discarding all values in that case;
// remaining results are drained in the background so no worker leaks.
//
// Contract: like Results, IntoSlice must be called at most once per group.
func (r *ExecuteGroup[T]) IntoSlice() ([]T, error) {
	const defaultCap = 5

	results := make([]T, 0, defaultCap)

	resultsChan := r.Results()
	for result := range resultsChan {
		if result.Err != nil {
			// Keep draining so pending workers can finish their send and
			// the closer goroutine (wg.Wait) can complete.
			go func() {
				for range resultsChan { //nolint:revive // intentional drain
				}
			}()

			return nil, result.Err
		}

		results = append(results, result.Ok)
	}

	return results, nil
}
