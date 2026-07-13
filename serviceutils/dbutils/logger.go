package dbutils

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
)

// ZerologAdapter implements gorm's logger.Interface by writing through the
// zerolog logger found on the request context.
type ZerologAdapter struct{}

// LogMode returns the adapter unchanged: the level is governed by the context
// logger, not by gorm, so the requested level is ignored.
func (z ZerologAdapter) LogMode(_ logger.LogLevel) logger.Interface {
	return z
}

// Info logs at info level using the ctx logger.
func (z ZerologAdapter) Info(ctx context.Context, s string, i ...interface{}) {
	zerolog.Ctx(ctx).Info().Msgf(s, i...)
}

// Warn logs at warn level using the ctx logger.
func (z ZerologAdapter) Warn(ctx context.Context, s string, i ...interface{}) {
	zerolog.Ctx(ctx).Warn().Msgf(s, i...)
}

// Error logs at error level using the ctx logger.
func (z ZerologAdapter) Error(ctx context.Context, s string, i ...interface{}) {
	zerolog.Ctx(ctx).Error().Msgf(s, i...)
}

// Trace logs one executed SQL statement at trace level with its row count,
// start time and error.
func (z ZerologAdapter) Trace(
	ctx context.Context,
	begin time.Time,
	fc func() (sql string, rowsAffected int64),
	err error,
) {
	sql, rowAffected := fc()

	zerolog.Ctx(ctx).Trace().
		Int64("row_affected", rowAffected).
		Str("sql", sql).
		Time("begin", begin).
		Err(err).
		Msgf("gorm.trace")
}
