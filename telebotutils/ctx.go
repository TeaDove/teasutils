package telebotutils

import (
	"context"
	"strconv"

	"github.com/teadove/teasutils/serviceutils/loggerutils"
	tele "gopkg.in/telebot.v4"
)

const contextKey = "telebotutils.context"

func Context(c tele.Context) context.Context {
	if c == nil {
		return nil
	}

	ctx, _ := c.Get(contextKey).(context.Context)

	return ctx
}

func GetOrSetCtx(parent context.Context) func(tele.Context) context.Context {
	return func(c tele.Context) context.Context {
		if c == nil {
			return parent
		}

		ctx := Context(c)
		if ctx != nil {
			return ctx
		}

		ctx = loggerutils.WithValue(parent, "update_id", strconv.Itoa(c.Update().ID))
		c.Set(contextKey, ctx)

		return ctx
	}
}

func WithContext(parent context.Context) tele.MiddlewareFunc {
	getContext := GetOrSetCtx(parent)

	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			getContext(c)

			return next(c)
		}
	}
}
