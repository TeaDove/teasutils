package mustutils

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

// Must returns obj, or panics if err is non-nil.
// Use only where an error is a programming error, e.g. package init.
func Must[T any](obj T, err error) T {
	if err != nil {
		panic(fmt.Errorf("must failed: %w", err))
	}

	return obj
}

// MustOk returns obj, or panics if ok is false.
// Handy for the comma-ok forms (map lookup, type assertion).
func MustOk[T any](obj T, ok bool) T {
	if !ok {
		panic(errors.New("ok failed"))
	}

	return obj
}

// MustNoReturn panics if err is non-nil, for calls that only return an error.
func MustNoReturn(err error) {
	if err != nil {
		panic(fmt.Errorf("must failed: %w", err))
	}
}
