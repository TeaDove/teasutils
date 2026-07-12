package fiberutils

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/teadove/teasutils/utils/errorsutils"

	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/serviceutils/loggerutils"
)

func ErrHandler() fiber.ErrorHandler {
	return func(c fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		var e *fiber.Error

		if errors.As(err, &e) {
			code = e.Code
		}

		if code >= http.StatusInternalServerError {
			zerolog.Ctx(c.Context()).
				Error().
				Stack().Err(errorsutils.WithStackIfRequired(err)).
				Int("code", code).
				Msg("http.internal.error")
		} else {
			zerolog.Ctx(c.Context()).
				Warn().
				Err(err).
				Int("code", code).
				Msg("http.client.error")
		}

		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		return c.Status(code).JSON(fiber.Map{"error": err.Error()})
	}
}

func MiddlewareLogger() fiber.Handler {
	return func(c fiber.Ctx) error {
		t0 := time.Now()
		ctx := loggerutils.AddLoggerToCtx(c.Context())
		ctx = loggerutils.WithValue(ctx, "ip", c.IP())
		ctx = loggerutils.WithValue(ctx, "app_method", fmt.Sprintf("%s %s", c.Method(), c.Path()))
		ctx = loggerutils.WithValue(ctx, "user_agent", strings.Clone(c.Get(fiber.HeaderUserAgent)))

		c.SetContext(ctx)

		err := c.Next()

		log := zerolog.Ctx(c.Context()).
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
	return func(c fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), dur)
		defer cancel()

		c.SetContext(ctx)

		return c.Next()
	}
}
