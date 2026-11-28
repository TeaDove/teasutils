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

// LogOnErr returns a Telebot error handler that logs with the update context when available.
// It uses parent for errors reported outside a handler.
func LogOnErr(parent context.Context) func(error, tele.Context) {
	return func(err error, c tele.Context) {
		logUpdateError(contextForError(parent, c), err)
	}
}

// ReportOnErr returns a Telebot error handler that logs the error and replies in the update's chat.
// Client errors expose their message and other errors receive a generic reply.
func ReportOnErr(parent context.Context) func(error, tele.Context) {
	return func(err error, c tele.Context) {
		ctx := contextForError(parent, c)
		logUpdateError(ctx, err)

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

// IsClientError reports whether err wraps an error created by NewClientError.
func IsClientError(err error) bool {
	var handlerError Error

	return errors.As(err, &handlerError) && handlerError.code == http.StatusBadRequest
}

func contextForError(parent context.Context, c tele.Context) context.Context {
	ctx := storedContext(c)
	if ctx == nil {
		return parent
	}

	return ctx
}

func logUpdateError(ctx context.Context, err error) {
	zerolog.Ctx(ctx).
		Error().
		Stack().
		Err(withoutRequestURL(err)).
		Msg("failed.to.process.tg.update")
}

func withoutRequestURL(err error) error {
	var requestErr *url.Error
	if errors.As(err, &requestErr) {
		return requestErr.Err
	}

	return err
}

// Error wraps a handler error with an HTTP status code for error reporting.
type Error struct {
	code int
	err  error
}

// Error renders the code and the wrapped error.
func (e Error) Error() string {
	return fmt.Sprintf("code=%d, err=%v", e.code, e.err)
}

// Unwrap returns the wrapped error.
func (e Error) Unwrap() error {
	return e.err
}

// NewClientError wraps an error whose message may be shown to the user.
func NewClientError(err error) Error {
	return Error{code: http.StatusBadRequest, err: err}
}

// NewServerError wraps an internal error reported with a generic reply.
func NewServerError(err error) Error {
	return Error{code: http.StatusInternalServerError, err: err}
}
