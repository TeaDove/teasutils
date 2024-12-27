package chan_utils

import "sync"

// Duplicator
// Duplicates data to 2 chains.
func Duplicator[T any](channel <-chan T, wg *sync.WaitGroup) (chan T, chan T) {
	a := make(chan T, cap(channel))
	b := make(chan T, cap(channel))

	wg.Add(1)

	go func() {
		defer wg.Done()

		var v T
		for v = range channel {
			a <- v
			b <- v
		}
	}()

	return a, b
}
