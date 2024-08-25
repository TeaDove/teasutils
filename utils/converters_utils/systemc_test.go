package converters_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_Converters_SystemCToFixed_Ok(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 3.14, ToFixed(3.1415926, 2))
}

func TestUnit_Converters_SystemCToKilo_Ok(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 2.0, ToKilo(2000))
}

func TestUnit_Converters_SystemCToKiloByte_Ok(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 2.0, ToKiloByte(2048))
}
