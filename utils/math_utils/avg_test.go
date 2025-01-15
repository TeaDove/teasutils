package math_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_MathUtils_Avg_CalculatesOk(t *testing.T) {
	t.Parallel()

	assert.InDelta(t, 2.0, Avg([]int{1, 2, 3}), 0.000001)
	assert.InDelta(t, 2.5, Avg([]int{1, 2, 3, 4}), 0.000001)
	assert.InDelta(t, 3230.75, Avg([]int{123, 6354, 123, 6323}), 0.0000001)
}

func TestUnit_MathUtils_AddToAvg_CalculatesOk(t *testing.T) {
	t.Parallel()

	assert.InDelta(t, 2.5, AddToAvg(2.0, 3, 4), 0.00000001)
	assert.InDelta(t, 3.0, AddToAvg(2.5, 4, 5), 0.00000001)
	assert.InDelta(t, 2784.4, AddToAvg(3230.75, 4, 999), 0.000001)
}

func TestUnit_MathUtils_AvgWithAvg_CalculatesOk(t *testing.T) {
	t.Parallel()

	assert.InDelta(t, 2.2857142857142856, AvgWithAvg(2.0, 3, 2.5, 4), 0.00000000001)
}
