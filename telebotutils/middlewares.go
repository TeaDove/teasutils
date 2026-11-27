package telebotutils

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
	tele "gopkg.in/telebot.v4"
)

func LogOnErr(parent context.Context) func(error, tele.Context) {
	getContext := GetOrSetCtx(parent)

	return func(err error, c tele.Context) {
		zerolog.Ctx(getContext(c)).
			Error().
			Stack().
			Err(withoutRequestURL(err)).
			Msg("failed.to.process.tg.update")
	}
}

func ReportOnErr(parent context.Context) func(error, tele.Context) {
	logError := LogOnErr(parent)
	getContext := GetOrSetCtx(parent)

	return func(err error, c tele.Context) {
		logError(err, c)

		ctx := getContext(c)
		if c == nil || c.Chat() == nil || ctx.Err() != nil {
			return
		}

		message := "An internal error occurred."

		var handlerError Error
		if errors.As(err, &handlerError) && handlerError.code == http.StatusBadRequest {
			message = handlerError.err.Error()
		}

		replyErr := c.Reply(message, tele.ModeDefault)
		if replyErr != nil {
			zerolog.Ctx(ctx).Error().Err(withoutRequestURL(replyErr)).Msg("failed.to.reply")
		}
	}
}

func IsClientError(err error) bool {
	var handlerError Error

	return errors.As(err, &handlerError) && handlerError.code == http.StatusBadRequest
}

func withoutRequestURL(err error) error {
	var requestErr *url.Error
	if errors.As(err, &requestErr) {
		return requestErr.Err
	}

	return err
}

type Error struct {
	code int
	err  error
}

func (e Error) Error() string {
	return fmt.Sprintf("code=%d, err=%v", e.code, e.err)
}

func (e Error) Unwrap() error {
	return e.err
}

func NewClientError(err error) Error {
	return Error{code: http.StatusBadRequest, err: err}
}

func NewServerError(err error) Error {
	return Error{code: http.StatusInternalServerError, err: err}
}
