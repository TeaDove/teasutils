package telebotutils

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/serviceutils/loggerutils"
	"github.com/teadove/teasutils/utils/redactutils"
	tele "gopkg.in/telebot.v4"
)

// GetOrSetCtx
// Gets context.Context from tele.Context, if nil - creates it and populates with update metadata.
func GetOrSetCtx(c tele.Context) context.Context {
	ctx, ok := c.Get("ctx").(context.Context)
	if ok {
		return ctx
	}

	ctx = loggerutils.NewLoggedCtx()
	if c.Chat() != nil && c.Chat().Title != "" {
		ctx = loggerutils.WithValue(ctx, "in", c.Chat().Title)
	}

	if c.Text() != "" {
		ctx = loggerutils.WithValue(ctx, "text", redactutils.Trim(c.Text()))
	}

	if c.Sender() != nil {
		ctx = loggerutils.WithValue(ctx, "from", c.Sender().Username)
	}

	c.Set("ctx", ctx)

	return ctx
}

// Log returns the request-scoped logger for the update, initialising the
// context via GetOrSetCtx if needed.
func Log(c tele.Context) *zerolog.Logger {
	return zerolog.Ctx(GetOrSetCtx(c))
}
