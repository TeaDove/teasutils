package fiber_utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
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
