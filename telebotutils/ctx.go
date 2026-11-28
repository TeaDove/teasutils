package telebotutils

import (
	"context"
	"strconv"

	"github.com/teadove/teasutils/serviceutils/loggerutils"
	tele "gopkg.in/telebot.v4"
)

const contextKey = "telebotutils.context"

// Context returns the update context installed by WithContext.
// It panics if the handler is running without that middleware.
func Context(c tele.Context) context.Context {
	ctx := storedContext(c)
	if ctx == nil {
		panic("telebotutils: WithContext middleware is required")
	}

	return ctx
}

// WithContext derives an update context from parent and adds the update ID to its logger.
func WithContext(parent context.Context) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			ctx := loggerutils.WithValue(parent, "update_id", strconv.Itoa(c.Update().ID))
			c.Set(contextKey, ctx)

			return next(c)
		}
	}
}

func storedContext(c tele.Context) context.Context {
	if c == nil {
		return nil
	}

	ctx, _ := c.Get(contextKey).(context.Context)

	return ctx
}
