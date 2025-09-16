package closer_utils

import (
	"context"
	"sync"
	"testing"

	"github.com/teadove/teasutils/utils/test_utils"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type Closable struct {
	isClosed   bool
	errOnClose bool
}

func (c *Closable) Close() error {
	if c.isClosed {
		return errors.New("already closed")
	}

	if c.errOnClose {
		return errors.New("failed to close")
	}

	c.isClosed = true

	return nil
}

func TestUnit_CloserUtils_CloseOrLog_Ok(t *testing.T) {
	t.Parallel()

	ctx := test_utils.GetLoggedContext()

	closable := Closable{}
	CloseOrLog(ctx, &closable)

	assert.True(t, closable.isClosed)
}

func TestUnit_CloserUtils_CloseOrLog_DontPanic(t *testing.T) {
	t.Parallel()

	ctx := test_utils.GetLoggedContext()

	closable := Closable{}
	closable.errOnClose = true
	CloseOrLog(ctx, &closable)

	assert.False(t, closable.isClosed)
}

func TestUnit_CloserUtils_CloseOrLogOnCtxDone_CancelOk(t *testing.T) {
	t.Parallel()

	ctx := test_utils.GetLoggedContext()
	ctx, cancel := context.WithCancel(ctx)

	var wg sync.WaitGroup

	closable := Closable{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		CloseOrLogOnCtxDone(ctx, &closable)
	}()

	assert.False(t, closable.isClosed)
	cancel()
	wg.Wait()
	assert.True(t, closable.isClosed)
}
