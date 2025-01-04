package di_utils

import (
	"context"
	"testing"
	"time"

	"github.com/teadove/teasutils/utils/settings_utils"

	"github.com/stretchr/testify/assert"

	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/logger_utils"
)

type TestContainer struct{}

func (r *TestContainer) HealthCheckers() []func(ctx context.Context) error {
	return []func(ctx context.Context) error{
		func(ctx context.Context) error {
			zerolog.Ctx(ctx).Info().Msg("health.checked")
			return nil
		},
	}
}

func (r *TestContainer) Stoppers() []func(ctx context.Context) error {
	return []func(ctx context.Context) error{
		func(ctx context.Context) error {
			time.Sleep(time.Minute)
			zerolog.Ctx(ctx).Info().Msg("stopped")
			return nil
		},
	}
}

func TestUnit_DIUtils_MustBuildFromSettingsAndRun_Ok(t *testing.T) {
	t.Parallel()

	ctx := logger_utils.NewLoggedCtx()
	settings_utils.BaseSettings.Metrics.URL = "0.0.0.0:8083"

	container := MustBuildFromSettings[*TestContainer](
		ctx,
		func(_ context.Context) (*TestContainer, error) {
			return &TestContainer{}, nil
		},
	)
	assert.NoError(t, checkFromCheckers(ctx, container.HealthCheckers()))
}

func TestUnit_DIUtils_MustBuildFromSettingsStopInf_LogsErr(t *testing.T) {
	t.Parallel()

	ctx := logger_utils.NewLoggedCtx()
	settings_utils.BaseSettings.Metrics.URL = "0.0.0.0:8084"

	container := MustBuildFromSettings[*TestContainer](
		ctx,
		func(_ context.Context) (*TestContainer, error) {
			return &TestContainer{}, nil
		},
	)

	settings_utils.BaseSettings.Metrics.CloseTimeout = time.Second
	err := stop(ctx, container)
	assert.Error(t, err)
}
