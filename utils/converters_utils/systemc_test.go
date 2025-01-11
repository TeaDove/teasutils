package converters_utils

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

	assert.Equal(t, "15 B", ToClosestByteAsString(15, 2))
	assert.Equal(t, "1.46 KB", ToClosestByteAsString(1500, 2))
	assert.Equal(t, "14.65 KB", ToClosestByteAsString(15000, 2))
	assert.Equal(t, "1.43 MB", ToClosestByteAsString(1500000, 2))
	assert.Equal(t, "14.31 MB", ToClosestByteAsString(15000000, 2))
	assert.Equal(t, "143.05 MB", ToClosestByteAsString(150000000, 2))
	assert.Equal(t, "1.4 GB", ToClosestByteAsString(1500000000, 2))
	assert.Equal(t, "139.7 GB", ToClosestByteAsString(150000000000, 2))
}
