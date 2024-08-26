package settings_utils

import (
	"context"
	"encoding/json"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/redact_utils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const (
	defaultEnvPrefix = "teas_"
	defaultEnvFile   = ".env"
)

func InitSetting[T any](ctx context.Context, omitFromLogValues ...string) (T, error) {
	_ = godotenv.Load(defaultEnvFile)

	settings, err := env.ParseAsWithOptions[T](env.Options{Prefix: defaultEnvPrefix})
	if err != nil {
		return *new(T), errors.Wrap(err, "failed to env parse")
	}

	settingsJson, err := json.Marshal(settings)
	if err != nil {
		return *new(T), errors.Wrap(err, "failed to marshal settings")
	}

	for _, valueKey := range omitFromLogValues {
		res := gjson.GetBytes(settingsJson, valueKey)
		if res.String() == "" {
			zerolog.Ctx(ctx).
				Warn().
				Str("key", valueKey).
				Msg("empty.value.for.omitted.field")
			continue
		}
		settingsJson, err = sjson.SetBytes(
			settingsJson,
			valueKey,
			redact_utils.RedactWithPrefix(res.String()),
		)
		if err != nil {
			return *new(T), errors.Wrap(err, "failed to redact settings")
		}
	}

	println(string(settingsJson))

	zerolog.Ctx(ctx).
		Debug().
		Interface("v", settings).
		Msg("settings.loaded")

	return settings, nil
}
