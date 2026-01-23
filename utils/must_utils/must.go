package must_utils

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(fmt.Errorf("must failed: %w", err))
	}

	return obj
}

func MustOk[T any](obj T, ok bool) T {
	if !ok {
		panic(errors.New("ok failed"))
	}

	return obj
}

func MustNoReturn(err error) {
	if err != nil {
		panic(fmt.Errorf("must failed: %w", err))
	}
}
