package fiberutils

import (
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v3"
)

// StatusFromContext resolves the HTTP status for a handler result: the actual
// response code when err is nil, the code of a *fiber.Error, or 500 otherwise.
func StatusFromContext(c fiber.Ctx, err error) int {
	if err == nil {
		return c.Response().StatusCode()
	}

	var e *fiber.Error

	code := fiber.StatusInternalServerError
	if errors.As(err, &e) {
		code = e.Code
	}

	return code
}
