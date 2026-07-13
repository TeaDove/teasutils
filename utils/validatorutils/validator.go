package validatorutils

import "github.com/go-playground/validator/v10"

// Validator is a shared, ready-to-use go-playground validator instance with
// required-struct validation enabled. Safe for concurrent use.
var Validator = validator.New(validator.WithRequiredStructEnabled()) //nolint: gochecknoglobals // allowed as singleton
