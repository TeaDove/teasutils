package telebotutils_test

import (
	"context"
	"errors"
	"testing"

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
		ctx := telebotutils.Context(c)
		if ctx == nil || ctx.Value(key{}) != "parent" {
			return errors.New("update did not inherit parent context")
		}

		if ctx != telebotutils.GetOrSetCtx(parent)(c) {
			return errors.New("update context was created more than once")
		}

		return nil
	})

	err := wrapped(c)
	if err != nil {
		t.Fatal(err)
	}

	cancel()

	if !errors.Is(telebotutils.Context(c).Err(), context.Canceled) {
		t.Fatal("canceling the parent did not cancel the update")
	}
}
