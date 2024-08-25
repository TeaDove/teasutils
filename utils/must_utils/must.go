package must_utils

import (
	"context"

	"github.com/pkg/errors"
)

func Must(ctx context.Context, err error) {
	if err != nil {
		FancyPanic(ctx, errors.Wrap(err, "must failed"))
	}
}

func MustWithReturn[T any](ctx context.Context, obj T, err error) T {
	if err != nil {
		FancyPanic(ctx, errors.Wrap(err, "must with return failed"))
	}
	return obj
}

func MustWithReturnWithoutContext[T any](obj T, err error) T {
	if err != nil {
		panic(errors.Wrap(err, "must with return failed"))
	}
	return obj
}
