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

type testService struct{}

func (r *testService) Health(ctx context.Context) error {
	zerolog.Ctx(ctx).Info().Msg("health.checked")
	return nil
}

func (r *testService) Close(ctx context.Context) error {
	time.Sleep(time.Minute)
	zerolog.Ctx(ctx).Info().Msg("stopped")

	return nil
}

func (r *TestContainer) Healths() []Health {
	return []Health{&testService{}}
}

func (r *TestContainer) Closers() []CloserWithContext {
	return []CloserWithContext{&testService{}}
}

//nolint: paralleltest // fails otherwise
func TestUnit_DIUtils_MustBuildFromSettingsAndRun_Ok(t *testing.T) {
	ctx := logger_utils.NewLoggedCtx()
	settings_utils.ServiceSettings.Metrics.URL = "0.0.0.0:8083"

	container := MustBuildFromSettings[*TestContainer](
		ctx,
		func(_ context.Context) (*TestContainer, error) {
			return &TestContainer{}, nil
		},
	)
	assert.NoError(t, checkFromCheckers(ctx, container.Healths()))
}

//nolint: paralleltest // fails otherwise
func TestUnit_DIUtils_MustBuildFromSettingsStopInf_LogsErr(t *testing.T) {
	ctx := logger_utils.NewLoggedCtx()
	settings_utils.ServiceSettings.Metrics.URL = "0.0.0.0:8084"

	container := MustBuildFromSettings[*TestContainer](
		ctx,
		func(_ context.Context) (*TestContainer, error) {
			return &TestContainer{}, nil
		},
	)

	settings_utils.ServiceSettings.Metrics.CloseTimeout = time.Second
	err := stop(ctx, container)
	assert.Error(t, err)
}
