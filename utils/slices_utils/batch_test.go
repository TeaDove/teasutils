package slices_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_SlicesUtils_CutInBatchesWithRemainder_Ok(t *testing.T) {
	t.Parallel()

	slice := []int{1, 2, 3, 4, 5}
	batched := CutInBatches(slice, 2)

	assert.Equal(t, [][]int{{1, 2}, {3, 4}, {5}}, batched)
}

func TestUnit_SlicesUtils_CutInBatchesWithoutRemainder_Ok(t *testing.T) {
	t.Parallel()

	slice := []int{1, 2, 3, 4, 5, 6}
	batched := CutInBatches(slice, 3)

	assert.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}}, batched)
}

func TestUnit_SlicesUtils_CutInBatchesBatchSizeBiggerThanSliceLen_Ok(t *testing.T) {
	t.Parallel()

	slice := []int{1, 2, 3, 4, 5, 6}
	batched := CutInBatches(slice, 10)

	assert.Equal(t, [][]int{{1, 2, 3, 4, 5, 6}}, batched)
}

func TestUnit_SlicesUtils_CutInBatchesEmptySlice_Ok(t *testing.T) {
	t.Parallel()

	var slice []int
	batched := CutInBatches(slice, 10)

	assert.Equal(t, [][]int{}, batched)
}
