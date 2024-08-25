package settings_utils

import (
	"context"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

const (
	defaultEnvPrefix = "teas_"
	defaultEnvFile   = ".env"
)

func InitSetting[T any](ctx context.Context) (T, error) {
	_ = godotenv.Load(defaultEnvFile)

	settings, err := env.ParseAsWithOptions[T](env.Options{Prefix: defaultEnvPrefix})
	if err != nil {
		return *new(T), errors.Wrap(err, "failed to env parse")
	}

	zerolog.Ctx(ctx).
		Debug().
		Interface("v", settings).
		Msg("settings.loaded")

	return settings, nil
}
