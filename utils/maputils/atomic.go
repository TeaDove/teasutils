package maputils

import (
	"sync/atomic"
)

// AtomicMap is a read-mostly, thread-safe map backed by an atomic.Pointer.
// It replaces the whole map on every write, so it avoids an RWMutex when
// writes are rare, done in the background and swap the entire map at once.
// The zero value is ready to use and safe for concurrent reads.
type AtomicMap[K comparable, V any] struct {
	v atomic.Pointer[map[K]V]
}

// Store atomically replaces the current map with v.
// The caller must not read or write v afterwards, since it is now shared.
func (r *AtomicMap[K, V]) Store(v map[K]V) {
	r.v.Store(&v)
}

// Get returns the value stored for key and whether it was present.
// On a never-stored (zero) map it returns the zero value and false.
func (r *AtomicMap[K, V]) Get(key K) (V, bool) {
	mapPtr := r.v.Load()
	if mapPtr == nil {
		return *new(V), false
	}

	v, ok := (*mapPtr)[key]

	return v, ok
}

// Len reports the number of entries, or 0 if nothing was stored yet.
func (r *AtomicMap[K, V]) Len() int {
	mapPtr := r.v.Load()
	if mapPtr == nil {
		return 0
	}

	return len(*mapPtr)
}
