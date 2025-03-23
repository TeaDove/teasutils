package fiber_utils

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/logger_utils"
)

func ErrHandler() func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		var e *fiber.Error

		if errors.As(err, &e) {
			code = e.Code
		}

		if code >= http.StatusInternalServerError {
			zerolog.Ctx(c.UserContext()).
				Error().
				Stack().Err(err).
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

type contextAppender func(c *fiber.Ctx, ctx context.Context) context.Context

type LogCtxConfig struct {
	DisableLogRequest bool
	DisableIP         bool
	DisableAPPMethod  bool
	DisableUserAgent  bool
}

var logCtxKey = "logCtx"

func MiddlewareLogger(config *LogCtxConfig) func(c *fiber.Ctx) error {
	contexts := make([]contextAppender, 0)
	if !config.DisableIP {
		contexts = append(contexts, func(c *fiber.Ctx, ctx context.Context) context.Context {
			return logger_utils.WithValue(ctx, "ip", c.IP())
		})
	}

	if !config.DisableAPPMethod {
		contexts = append(contexts, func(c *fiber.Ctx, ctx context.Context) context.Context {
			return logger_utils.WithValue(
				ctx,
				"app_method",
				fmt.Sprintf("%s %s", c.Method(), c.Path()),
			)
		})
	}

	if !config.DisableUserAgent {
		contexts = append(contexts, func(c *fiber.Ctx, ctx context.Context) context.Context {
			return logger_utils.WithValue(ctx, "user_agent", strings.Clone(c.Get("User-Agent")))
		})
	}

	return func(c *fiber.Ctx) error {
		ctx := logger_utils.AddLoggerToCtx(c.UserContext())
		for _, appender := range contexts {
			ctx = appender(c, ctx) //nolint: fatcontext // fp
		}

		c.SetUserContext(ctx)
		c.Locals(logCtxKey, ctx)

		err := c.Next()
		if err == nil && !config.DisableLogRequest {
			statusCode := c.Response().StatusCode()

			switch {
			case statusCode < http.StatusBadRequest:
				zerolog.Ctx(c.UserContext()).
					Info().
					Int("code", statusCode). // TODO add resp-size and duration
					Msg("request.processed")
			case statusCode < http.StatusInternalServerError:
				zerolog.Ctx(c.UserContext()).
					Warn().
					Int("code", statusCode). // TODO add resp-size and duration
					Msg("request.processed")
			default:
				zerolog.Ctx(c.UserContext()).
					Error().
					Int("code", statusCode). // TODO add resp-size and duration
					Msg("request.processed")
			}
		}

		return err //nolint: wrapcheck // fp
	}
}

func GetLogCtx(c interface {
	Locals(key string, value ...any) any
}) context.Context {
	return c.Locals(logCtxKey).(context.Context)
}
