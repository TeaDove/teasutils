package concutils

// sem is a counting semaphore built on a buffered channel: a successful
// lock consumes one slot and unlock releases it.
type sem chan struct{}

func (r sem) lock() {
	r <- struct{}{}
}

func (r sem) unlock() {
	<-r
}

// newSem creates a semaphore allowing up to limit concurrent holders.
// It panics if limit <= 0, since a zero-capacity semaphore would
// dead-lock on the first lock call.
func newSem(limit int) sem {
	if limit <= 0 {
		panic("concutils: limit must be greater than zero")
	}

	return make(sem, limit)
}
