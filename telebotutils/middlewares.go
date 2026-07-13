package telebotutils

import (
	"fmt"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
	tele "gopkg.in/telebot.v4"
)

// LogOnErr logs err at error level (with a stack) against the update's
// context. Suitable as a telebot OnError handler that only logs.
func LogOnErr(err error, c tele.Context) {
	zerolog.Ctx(GetOrSetCtx(c)).
		Error().
		Stack().Err(err).
		Msg("failed.to.process.tg.update")
}

// ReportOnErr logs err and, for errors whose code is >= 400, replies the error
// text to the user. Non-Error values are treated as 500. Use NewClientError /
// NewServerError to control the code carried by err.
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

// Error is a telebot handler error carrying an HTTP-like status code that
// ReportOnErr uses to decide how to log and whether to reply to the user.
type Error struct {
	code int
	err  error
}

// Error implements the error interface, rendering the code and wrapped error.
func (e Error) Error() string {
	return fmt.Sprintf("code=%d, err=%v", e.code, e.err)
}

// NewClientError wraps err as a client-side Error (code 400).
func NewClientError(err error) Error {
	return Error{
		code: http.StatusBadRequest,
		err:  err,
	}
}

// NewServerError wraps err as a server-side Error (code 500).
func NewServerError(err error) Error {
	return Error{
		code: http.StatusInternalServerError,
		err:  err,
	}
}
