package loggerutils

import (
	"context"

	"github.com/rs/zerolog"
)

// NewLoggedCtx returns a background context carrying the package's global logger.
func NewLoggedCtx() context.Context {
	return AddLoggerToCtx(context.Background())
}

// AddLoggerToCtx attaches the global logger to ctx and returns the new context.
func AddLoggerToCtx(ctx context.Context) context.Context {
	return globalLogger.With().Logger().WithContext(ctx)
}

// NewLvlLoggedCtx returns a background context carrying the global logger
// capped at level lvl.
func NewLvlLoggedCtx(lvl zerolog.Level) context.Context {
	return AddLvlLoggerToCtx(context.Background(), lvl)
}

// AddLvlLoggerToCtx attaches the global logger at level lvl to ctx.
func AddLvlLoggerToCtx(ctx context.Context, lvl zerolog.Level) context.Context {
	return globalLogger.With().Logger().Level(lvl).WithContext(ctx)
}

// WithValue returns a context whose logger carries the given key/value string
// fields. kv is read in pairs; a trailing unpaired element is ignored.
func WithValue(ctx context.Context, kv ...string) context.Context {
	logger := zerolog.Ctx(ctx).With()
	for idx := 0; idx < len(kv)-1; idx += 2 {
		logger = logger.Str(kv[idx], kv[idx+1])
	}

	return logger.Ctx(ctx).Logger().WithContext(ctx)
}

type loggerCtxKey string

// WithReadableValue adds key/value to the logger (like WithValue) and also
// stores it on the context so it can be retrieved later with ReadValue.
func WithReadableValue(ctx context.Context, key string, value string) context.Context {
	return WithValue(context.WithValue(ctx, loggerCtxKey(key), value), key, value)
}

// ReadValue returns the value stored by WithReadableValue for key, or "" if
// absent. It panics if a value is present but is not a string.
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
