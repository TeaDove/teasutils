package telebot_utils

import (
	"fmt"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
	tele "gopkg.in/telebot.v4"
)

func LogOnErr(err error, c tele.Context) {
	zerolog.Ctx(GetOrSetCtx(c)).
		Error().
		Stack().Err(err).
		Msg("failed.to.process.tg.update")
}

func ReportOnErr(err error, c tele.Context) {
	err = errors.Wrap(err, "failed to process update")

	var teleErr Error

	if !errors.As(err, &teleErr) {
		teleErr = Error{
			code: http.StatusInternalServerError,
			err:  err,
		}
	}

	if teleErr.code >= http.StatusBadRequest {
		zerolog.Ctx(GetOrSetCtx(c)).
			Warn().
			Err(err).
			Msg("failed.to.process.tg.update")

		innerErr := c.Reply(err.Error(), tele.ModeDefault)
		if innerErr != nil {
			zerolog.Ctx(GetOrSetCtx(c)).
				Error().
				Stack().Err(err).
				Msg("failed.to.reply")
		}
	} else {
		zerolog.Ctx(GetOrSetCtx(c)).
			Error().
			Stack().Err(err).
			Msg("failed.to.process.tg.update")
	}
}

type Error struct {
	code int
	err  error
}

func (e Error) Error() string {
	return fmt.Sprintf("code=%d, err=%v", e.code, e.err)
}

func NewClientError(err error) Error {
	return Error{
		code: http.StatusBadRequest,
		err:  err,
	}
}

func NewServerError(err error) Error {
	return Error{
		code: http.StatusInternalServerError,
		err:  err,
	}
}
