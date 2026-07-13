package fiberutils

import (
	"github.com/cockroachdb/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

// StructValidator adapts a go-playground validator to fiber's StructValidator
// interface so bound structs are validated automatically.
type StructValidator struct {
	Validator *validator.Validate
}

// Validate validates out using the underlying validator, returning its error.
func (v *StructValidator) Validate(out any) error {
	return v.Validator.Struct(out)
}

// NewDefaultStructValidator returns a StructValidator with required-struct
// validation enabled, ready to assign to fiber.Config.StructValidator.
func NewDefaultStructValidator() *StructValidator {
	return &StructValidator{validator.New(validator.WithRequiredStructEnabled())}
}

// BindJSON decodes the request body into T. On failure it returns the zero T
// and a 422 Unprocessable Entity fiber error (already suitable for returning
// from a handler).
func BindJSON[T any](c fiber.Ctx) (T, error) {
	var req T

	err := c.Bind().JSON(&req)
	if err != nil {
		return *new(T), SendUnprocessable(errors.Wrap(err, "failed to bind JSON body"))
	}

	return req, nil
}

// SendUnprocessable wraps err as a 422 Unprocessable Entity fiber error.
func SendUnprocessable(err error) error {
	return &fiber.Error{Code: fiber.StatusUnprocessableEntity, Message: err.Error()}
}

// SendBadRequest wraps err as a 400 Bad Request fiber error.
func SendBadRequest(err error) error {
	return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
}
