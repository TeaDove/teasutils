package slices_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_SlicesUtils_PIdx_Ok(t *testing.T) {
	t.Parallel()

	a := []int{1, 2, 3}

	assert.Equal(t, 1, PIdx(a, 1))
	assert.Equal(t, 6, PIdx(a, 6))
	assert.Equal(t, 0, PIdx(a, 0))
	assert.Equal(t, 2, PIdx(a, -1))
	assert.Equal(t, 1, PIdx(a, -2))
	assert.Equal(t, 0, PIdx(a, -3))
}

func TestUnit_SlicesUtils_PGet_Ok(t *testing.T) {
	t.Parallel()

	a := []int{1, 2, 3}

	assert.Equal(t, 2, PGet(a, 1))
	assert.Equal(t, 1, PGet(a, 0))
	assert.Equal(t, 3, PGet(a, -1))
	assert.Equal(t, 2, PGet(a, -2))
	assert.Equal(t, 1, PGet(a, -3))
}

func TestUnit_SlicesUtils_PAGet_Ok(t *testing.T) {
	t.Parallel()

	a := []int{1, 2, 3}

	assert.Equal(t, []int{2}, PAGet(a, 1, 2))
	assert.Equal(t, []int{1, 2, 3}, PAGet(a, 0, 3))
	assert.Equal(t, []int{1, 2}, PAGet(a, 0, -1))
	assert.Equal(t, []int{1, 2}, PAGet(a, -3, -1))
}

func TestUnit_SlicesUtils_PSIdx_Ok(t *testing.T) {
	t.Parallel()

	a := "123"

	assert.Equal(t, 1, PSIdx(a, 1))
	assert.Equal(t, 6, PSIdx(a, 6))
	assert.Equal(t, 0, PSIdx(a, 0))
	assert.Equal(t, 2, PSIdx(a, -1))
	assert.Equal(t, 1, PSIdx(a, -2))
	assert.Equal(t, 0, PSIdx(a, -3))
}

func TestUnit_SlicesUtils_PSGet_Ok(t *testing.T) {
	t.Parallel()

	a := "123"

	assert.Equal(t, byte('2'), PSGet(a, 1))
	assert.Equal(t, byte('1'), PSGet(a, 0))
	assert.Equal(t, byte('3'), PSGet(a, -1))
	assert.Equal(t, byte('2'), PSGet(a, -2))
	assert.Equal(t, byte('1'), PSGet(a, -3))
}

func TestUnit_SlicesUtils_PSAGet_Ok(t *testing.T) {
	t.Parallel()

	a := "123"

	assert.Equal(t, "2", PSAGet(a, 1, 2))
	assert.Equal(t, "123", PSAGet(a, 0, 3))
	assert.Equal(t, "12", PSAGet(a, 0, -1))
	assert.Equal(t, "12", PSAGet(a, -3, -1))
}
