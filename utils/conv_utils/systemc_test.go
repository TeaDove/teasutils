package conv_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_Converters_SystemCToFixed_Ok(t *testing.T) {
	t.Parallel()

	assert.InDelta(t, 3.14, ToFixed(3.1415926, 2), 0.00001)
}

func TestUnit_Converters_SystemCToKilo_Ok(t *testing.T) {
	t.Parallel()

	assert.InDelta(t, 2.0, ToKilo(2000), 0.01)
}

func TestUnit_Converters_SystemCToKiloByte_Ok(t *testing.T) {
	t.Parallel()

	assert.InDelta(t, 2.0, ToKiloByte(2048), 0.01)
}

func TestUnit_Converters_SystemCToClosestByteAsString_Ok(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "15 B", ClosestByte(15))
	assert.Equal(t, "1.46 kB", ClosestByte(1500))
	assert.Equal(t, "14.65 kB", ClosestByte(15000))
	assert.Equal(t, "1.43 MB", ClosestByte(1500000))
	assert.Equal(t, "14.31 MB", ClosestByte(15000000))
	assert.Equal(t, "143.05 MB", ClosestByte(150000000))
	assert.Equal(t, "1.4 GB", ClosestByte(1500000000))
	assert.Equal(t, "139.7 GB", ClosestByte(150000000000))
}

func TestUnit_Converters_SystemCToClosestKAsString_Ok(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "1.5 n", Closest(0.0000000015))
	assert.Equal(t, "1.5 µ", Closest(0.0000015))
	assert.Equal(t, "15 m", Closest(0.015))
	assert.Equal(t, "15", Closest(15))
	assert.Equal(t, "1.5 k", Closest(1500))
	assert.Equal(t, "15 k", Closest(15000))
	assert.Equal(t, "1.5 M", Closest(1500000))
	assert.Equal(t, "15 M", Closest(15000000))
	assert.Equal(t, "150 M", Closest(150000000))
	assert.Equal(t, "1.5 G", Closest(1500000000))
	assert.Equal(t, "150 G", Closest(150000000000))
}
