package slices_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_SlicesUtils_PythoniseIdx_Ok(t *testing.T) {
	t.Parallel()

	a := []int{1, 2, 3}

	assert.Equal(t, 1, PythoniseIdx(a, 1))
	assert.Equal(t, 6, PythoniseIdx(a, 6))
	assert.Equal(t, 0, PythoniseIdx(a, 0))
	assert.Equal(t, 2, PythoniseIdx(a, -1))
	assert.Equal(t, 1, PythoniseIdx(a, -2))
	assert.Equal(t, 0, PythoniseIdx(a, -3))
}

func TestUnit_SlicesUtils_PythoniseIdxGet_Ok(t *testing.T) {
	t.Parallel()

	a := []string{"1", "2", "3"}

	assert.Equal(t, "2", PythoniseIdxGet(a, 1))
	assert.Equal(t, "1", PythoniseIdxGet(a, 0))
	assert.Equal(t, "3", PythoniseIdxGet(a, -1))
	assert.Equal(t, "2", PythoniseIdxGet(a, -2))
	assert.Equal(t, "1", PythoniseIdxGet(a, -3))
}
