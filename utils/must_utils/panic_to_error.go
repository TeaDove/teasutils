package must_utils

import (
	"fmt"

	"github.com/pkg/errors"
)

func AnyToErr(v any) error {
	if v == nil {
		return nil
	}

	err, ok := v.(error)
	if ok {
		return errors.Wrap(err, "panicked")
	}

	strErr, ok := v.(string)
	if ok {
		return errors.Wrap(errors.New("panicked"), strErr)
	}

	return fmt.Errorf("panicked: %v", v)
}
