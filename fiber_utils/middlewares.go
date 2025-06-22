package fiber_utils

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/teadove/teasutils/utils/errors_utils"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/logger_utils"
)

func ErrHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		var e *fiber.Error

		if errors.As(err, &e) {
			code = e.Code
		}

		if code >= http.StatusInternalServerError {
			zerolog.Ctx(c.UserContext()).
				Error().
				Stack().Err(errors_utils.WithStackIfRequired(err)).
				Int("code", code).
				Msg("http.internal.error")
		} else {
			zerolog.Ctx(c.UserContext()).
				Warn().
				Err(err).
				Int("code", code).
				Msg("http.client.error")
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

		return c.Status(code).SendString(err.Error())
	}
}

func MiddlewareLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		t0 := time.Now()
		ctx := logger_utils.AddLoggerToCtx(c.UserContext())
		ctx = logger_utils.WithValue(ctx, "ip", c.IP())
		ctx = logger_utils.WithValue(ctx, "app_method", fmt.Sprintf("%s %s", c.Method(), c.Path()))
		ctx = logger_utils.WithValue(ctx, "user_agent", strings.Clone(c.Get(fiber.HeaderUserAgent)))

		c.SetUserContext(ctx)

		err := c.Next()

		log := zerolog.Ctx(c.UserContext()).
			Debug().
			Str("latency", time.Since(t0).String()).
			Int("code", StatusFromContext(c, err))

		if c.Request().Header.ContentLength() > 0 {
			log.Int("req_len", c.Request().Header.ContentLength())
		}

		if c.Response().Header.ContentLength() > 0 {
			log.Int("resp_len", c.Response().Header.ContentLength())
		}

		log.Msg("request.processed")

		return err //nolint: wrapcheck // fp
	}
}

func MiddlewareCtxTimeout(dur time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.UserContext(), dur)
		defer cancel()
		c.SetUserContext(ctx)

		return c.Next()
	}
}
