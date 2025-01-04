package di_utils

import (
	"context"
	"testing"

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
			zerolog.Ctx(ctx).Info().Msg("stoped")
			return nil
		},
	}
}

// nolint: paralleltest // will fail otherwise
func TestUnit_DIUtils_MustBuildFromSettingsAndRun_Ok(t *testing.T) {
	ctx := logger_utils.NewLoggedCtx()

	container := MustBuildFromSettings[*TestContainer](
		ctx,
		func(_ context.Context) (*TestContainer, error) {
			return &TestContainer{}, nil
		},
	)
	assert.NoError(t, checkFromCheckers(ctx, container.HealthCheckers()))
}
