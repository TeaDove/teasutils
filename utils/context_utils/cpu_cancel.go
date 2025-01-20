package context_utils

import (
	"context"

	"github.com/teadove/teasutils/utils/must_utils"
)

// CPUCancel
// Useful when you need to cancel no-io-bound cpu routine.
func CPUCancel(ctx context.Context, f func(ctx context.Context) error) (err error) {
	defer func() {
		err = must_utils.AnyToErr(recover())
	}()

	errChan := make(chan error, 1)
	go func() {
		errChan <- f(ctx)
	}()

	select {
	case err = <-errChan:
		return err
	case <-ctx.Done():
		panic(ctx.Err())
	}
}
