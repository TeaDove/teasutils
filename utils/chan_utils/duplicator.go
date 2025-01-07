package chan_utils

import "sync"

// Both channels are never closed, cap is the same as in parent chan.
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
