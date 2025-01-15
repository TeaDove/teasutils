package math_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_MathUtils_SliceLinear_Ok(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, SliceLinear(10, 10))
	assert.Equal(
		t,
		[]float64{0.909, 1.818, 2.727, 3.636, 4.545, 5.455, 6.364, 7.273, 8.182, 9.091, 10},
		SliceLinear(10, 11),
	)
	assert.Equal(t, []float64{1, 2, 3}, SliceLinear(3, 3))
}

func TestUnit_MathUtils_SliceGeometricProgression_Ok(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []float64{2, 6, 18, 54, 162}, SliceGeometricProgression(2, 3, 4))
	assert.Equal(
		t,
		[]float64{
			0.1,
			0.3,
			0.9,
			2.7,
			8.1,
			24.3,
			72.9,
			218.7,
			656.1,
			1968.3,
			5904.9,
			17714.7,
			53144.1,
			159432.3,
			478296.9,
			1.4348907e+06,
		},
		SliceGeometricProgression(0.1, 3, 15),
	)
}
