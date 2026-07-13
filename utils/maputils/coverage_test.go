package maputils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomicMap_EmptyIsSafe(t *testing.T) {
	t.Parallel()

	var m AtomicMap[string, int]

	v, ok := m.Get("missing")
	assert.False(t, ok)
	assert.Zero(t, v)
	assert.Equal(t, 0, m.Len())
}

func TestRWMap_StoreKeysLenAndMiss(t *testing.T) {
	t.Parallel()

	m := NewRWMap[string, int](4)

	_, ok := m.Get("missing")
	assert.False(t, ok)
	assert.Equal(t, 0, m.Len())

	m.Store(map[string]int{"a": 1, "b": 2})
	assert.Equal(t, 2, m.Len())
	assert.ElementsMatch(t, []string{"a", "b"}, m.Keys())

	m.Set("c", 3)
	got, ok := m.Get("c")
	assert.True(t, ok)
	assert.Equal(t, 3, got)
	assert.Equal(t, 3, m.Len())
}
