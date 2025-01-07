package slices_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_SlicesUtils_PythoniseIdx_Ok(t *testing.T) {
	t.Parallel()

	a := "123"

	assert.Equal(t, 1, PSIdx(a, 1))
	assert.Equal(t, 6, PSIdx(a, 6))
	assert.Equal(t, 0, PSIdx(a, 0))
	assert.Equal(t, 2, PSIdx(a, -1))
	assert.Equal(t, 1, PSIdx(a, -2))
	assert.Equal(t, 0, PSIdx(a, -3))
}

func TestUnit_SlicesUtils_PythoniseIdxGet_Ok(t *testing.T) {
	t.Parallel()

	a := "123"

	assert.Equal(t, byte('2'), PSGet(a, 1))
	assert.Equal(t, byte('1'), PSGet(a, 0))
	assert.Equal(t, byte('3'), PSGet(a, -1))
	assert.Equal(t, byte('2'), PSGet(a, -2))
	assert.Equal(t, byte('1'), PSGet(a, -3))
}

func TestUnit_SlicesUtils_PythoniseIdxSlice_Ok(t *testing.T) {
	t.Parallel()

	a := []int{1, 2, 3}

	assert.Equal(t, 1, PIdx(a, 1))
	assert.Equal(t, 6, PIdx(a, 6))
	assert.Equal(t, 0, PIdx(a, 0))
	assert.Equal(t, 2, PIdx(a, -1))
	assert.Equal(t, 1, PIdx(a, -2))
	assert.Equal(t, 0, PIdx(a, -3))
}

func TestUnit_SlicesUtils_PythoniseIdxGetSlice_Ok(t *testing.T) {
	t.Parallel()

	a := []int{1, 2, 3}

	assert.Equal(t, 2, PGet(a, 1))
	assert.Equal(t, 1, PGet(a, 0))
	assert.Equal(t, 3, PGet(a, -1))
	assert.Equal(t, 2, PGet(a, -2))
	assert.Equal(t, 1, PGet(a, -3))
}
