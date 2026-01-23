package fiber_utils

import (
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

func StatusFromContext(c *fiber.Ctx, err error) int {
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
