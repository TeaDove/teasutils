package telebotutils_test

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/teadove/teasutils/telebotutils"
	tele "gopkg.in/telebot.v4"
)

func TestWithContextInheritsParent(t *testing.T) {
	t.Parallel()

	type key struct{}

	parent, cancel := context.WithCancel(context.WithValue(t.Context(), key{}, "parent"))
	defer cancel()

	c := tele.NewContext(nil, tele.Update{ID: 42})
	wrapped := telebotutils.WithContext(parent)(func(c tele.Context) error {
		assert.Equal(t, "parent", telebotutils.Context(c).Value(key{}))

		return nil
	})

	err := wrapped(c)
	if assert.NoError(t, err) {
		cancel()
		assert.ErrorIs(t, telebotutils.Context(c).Err(), context.Canceled)
	}
}

func TestContextRequiresMiddleware(t *testing.T) {
	t.Parallel()

	c := tele.NewContext(nil, tele.Update{ID: 42})

	assert.PanicsWithValue(t, "telebotutils: WithContext middleware is required", func() {
		telebotutils.Context(c)
	})
}

func TestReportOnErrWithoutUpdateUsesParent(t *testing.T) {
	t.Parallel()

	var output bytes.Buffer

	logger := zerolog.New(&output)
	parent := logger.WithContext(t.Context())
	telebotutils.ReportOnErr(parent)(errors.New("test error"), nil)

	assert.Contains(t, output.String(), "test error")
}
