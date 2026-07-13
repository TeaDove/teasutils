package maputils

import (
	"maps"
	"slices"
	"sync"
)

// RWMap is a generic map guarded by an RWMutex, safe for concurrent use.
// Prefer it over AtomicMap when writes are frequent and touch single keys.
// The zero value is not usable; create one with NewRWMap.
type RWMap[K comparable, T any] struct {
	mapping map[K]T
	mu      sync.RWMutex
}

// NewRWMap returns an RWMap whose backing map is pre-sized to capacity.
func NewRWMap[K comparable, T any](capacity int) *RWMap[K, T] {
	return &RWMap[K, T]{mapping: make(map[K]T, capacity)}
}

// Get returns the value for k and whether it was present.
func (r *RWMap[K, T]) Get(k K) (T, bool) {
	r.mu.RLock()
	t, ok := r.mapping[k]
	r.mu.RUnlock()

	return t, ok
}

// Set stores t under key k, overwriting any existing value.
func (r *RWMap[K, T]) Set(k K, t T) {
	r.mu.Lock()
	r.mapping[k] = t
	r.mu.Unlock()
}

// Keys returns a snapshot of the current keys in unspecified order.
func (r *RWMap[K, T]) Keys() []K {
	r.mu.RLock()
	keys := slices.Collect(maps.Keys(r.mapping))
	r.mu.RUnlock()

	return keys
}

// Store replaces the whole backing map with mapping.
// The caller must not mutate mapping afterwards, as it is adopted by reference.
func (r *RWMap[K, T]) Store(mapping map[K]T) {
	r.mu.Lock()
	r.mapping = mapping
	r.mu.Unlock()
}

// Copy returns a shallow copy of the map that is safe to read without locking.
func (r *RWMap[K, T]) Copy() map[K]T {
	r.mu.RLock()
	mapping := make(map[K]T, len(r.mapping))
	maps.Copy(mapping, r.mapping)
	r.mu.RUnlock()

	return mapping
}

// Len reports the number of entries.
func (r *RWMap[K, T]) Len() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.mapping)
}
