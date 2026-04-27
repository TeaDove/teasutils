package fiber_utils

import "github.com/go-playground/validator/v10"

type StructValidator struct {
	Validator *validator.Validate
}

func (v *StructValidator) Validate(out any) error {
	return v.Validator.Struct(out)
}

func NewDefaultStructValidator() *StructValidator {
	return &StructValidator{validator.New(validator.WithRequiredStructEnabled())}
}
