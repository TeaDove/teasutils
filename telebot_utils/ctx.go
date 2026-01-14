package telebot_utils

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/service_utils/logger_utils"
	"github.com/teadove/teasutils/utils/redact_utils"
	tele "gopkg.in/telebot.v4"
)

// GetOrSetCtx
// Gets context.Context from tele.Context, if nil - creates it and populates with update metadata.
func GetOrSetCtx(c tele.Context) context.Context {
	ctx, ok := c.Get("ctx").(context.Context)
	if ok {
		return ctx
	}

	ctx = logger_utils.NewLoggedCtx()
	if c.Chat() != nil && c.Chat().Title != "" {
		ctx = logger_utils.WithValue(ctx, "in", c.Chat().Title)
	}

	if c.Text() != "" {
		ctx = logger_utils.WithValue(ctx, "text", redact_utils.Trim(c.Text()))
	}

	if c.Sender() != nil {
		ctx = logger_utils.WithValue(ctx, "from", c.Sender().Username)
	}

	c.Set("ctx", ctx)

	return ctx
}

func Log(c tele.Context) *zerolog.Logger {
	return zerolog.Ctx(GetOrSetCtx(c))
}
