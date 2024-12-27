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
