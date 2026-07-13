package fiberutils

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/serviceutils/loggerutils"
)

// ErrHandler returns a fiber error handler that logs the error (at error level
// with a stack for 5xx, warn for 4xx) and writes a JSON {"error": ...} body
// with the resolved status code.
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
				Stack().Err(err).
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

// MiddlewareLogger injects a request-scoped logger (with ip, method+path and
// user-agent fields) into the context and logs one debug line per request with
// latency, status code and request/response sizes.
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

// MiddlewareCtxTimeout bounds each request's context with a dur timeout,
// so downstream handlers and DB calls are cancelled when it elapses.
func MiddlewareCtxTimeout(dur time.Duration) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), dur)
		defer cancel()

		c.SetContext(ctx)

		return c.Next()
	}
}
