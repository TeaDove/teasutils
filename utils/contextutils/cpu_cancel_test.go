package contextutils

import (
	"context"
	"testing"
	"time"

	"github.com/teadove/teasutils/utils/testutils"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestUnit_ContextUtils_CPUCancel_Ok(t *testing.T) {
	t.Parallel()

	ctx := testutils.GetLoggedContext()
	ctx, cancel := context.WithCancel(ctx)

	cancel()

	err := CPUCancel(ctx, func(_ context.Context) error {
		for {
			time.Sleep(time.Second)
		}
	})

	require.Error(t, err)
	assert.Equal(t, "panicked: context canceled", err.Error())
}
