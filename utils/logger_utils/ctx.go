package logger_utils

import (
	"context"

	"github.com/rs/zerolog"
)

func NewLoggedCtx() context.Context {
	return AddLoggerToCtx(context.Background())
}

func AddLoggerToCtx(ctx context.Context) context.Context {
	return globalLogger.With().Logger().WithContext(ctx)
}

func NewLvlLoggedCtx(lvl zerolog.Level) context.Context {
	return AddLvlLoggerToCtx(context.Background(), lvl)
}

func AddLvlLoggerToCtx(ctx context.Context, lvl zerolog.Level) context.Context {
	return globalLogger.With().Logger().Level(lvl).WithContext(ctx)
}

func WithValue(ctx context.Context, key string, value string) context.Context {
	return zerolog.Ctx(ctx).With().Str(key, value).Ctx(ctx).Logger().WithContext(ctx)
}

type loggerCtxKey string

func WithReadableValue(ctx context.Context, key string, value string) context.Context {
	return WithValue(context.WithValue(ctx, loggerCtxKey(key), value), key, value)
}

func ReadValue(ctx context.Context, key string) string {
	v := ctx.Value(loggerCtxKey(key))
	if v == nil {
		return ""
	}

	res, ok := v.(string)
	if !ok {
		panic("value is not a string")
	}

	return res
}
