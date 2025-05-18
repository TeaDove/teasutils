package errors_utils

import "github.com/pkg/errors"

func WithStackIfRequired(err error) error {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	_, ok := err.(stackTracer)
	if ok {
		return err
	}

	return errors.WithStack(err)
}
