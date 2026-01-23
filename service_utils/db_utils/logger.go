package db_utils

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
)

type ZerologAdapter struct{}

func (z ZerologAdapter) LogMode(_ logger.LogLevel) logger.Interface {
	panic("no implemented")
}

func (z ZerologAdapter) Info(ctx context.Context, s string, i ...interface{}) {
	zerolog.Ctx(ctx).Info().Msgf(s, i...)
}

func (z ZerologAdapter) Warn(ctx context.Context, s string, i ...interface{}) {
	zerolog.Ctx(ctx).Warn().Msgf(s, i...)
}

func (z ZerologAdapter) Error(ctx context.Context, s string, i ...interface{}) {
	zerolog.Ctx(ctx).Error().Msgf(s, i...)
}

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
