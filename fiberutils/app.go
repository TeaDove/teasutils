package fiberutils

import (
	"github.com/cockroachdb/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type StructValidator struct {
	Validator *validator.Validate
}

func (v *StructValidator) Validate(out any) error {
	return v.Validator.Struct(out)
}

func NewDefaultStructValidator() *StructValidator {
	return &StructValidator{validator.New(validator.WithRequiredStructEnabled())}
}

func BindJSON[T any](c fiber.Ctx) (T, error) {
	var req T

	err := c.Bind().JSON(&req)
	if err != nil {
		return *new(T), SendUnprocessable(errors.Wrap(err, "failed to bind JSON body"))
	}

	return req, nil
}

func SendUnprocessable(err error) error {
	return &fiber.Error{Code: fiber.StatusUnprocessableEntity, Message: err.Error()}
}

func SendBadRequest(err error) error {
	return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
}
