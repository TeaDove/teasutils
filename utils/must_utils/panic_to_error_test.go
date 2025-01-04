package must_utils

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestUnit_MustUtils_PanicToError(t *testing.T) {
	t.Parallel()

	defer func() {
		err := AnyToErr(recover())
		require.Error(t, err)
		assert.Equal(t, "aaaa: paniced", err.Error())
	}()

	panic("aaaa")
}
