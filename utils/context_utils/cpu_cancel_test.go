package context_utils

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
	"github.com/teadove/teasutils/utils/logger_utils"
)

func TestUnit_ContextUtils_CPUCancel_Ok(t *testing.T) {
	t.Parallel()

	ctx := logger_utils.NewLoggedCtx()
	ctx, cancel := context.WithCancel(ctx)

	cancel()

	err := CPUCancel(ctx, func(_ context.Context) error {
		for {
			time.Sleep(time.Second)
		}
	})

	require.Error(t, err)
	assert.Equal(t, "paniced: context canceled", err.Error())
}
