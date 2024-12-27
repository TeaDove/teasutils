package must_utils

import "fmt"

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(fmt.Errorf("must failed: %w", err))
	}

	return obj
}

func MustNoReturn(err error) {
	if err != nil {
		panic(fmt.Errorf("must failed: %w", err))
	}
}
