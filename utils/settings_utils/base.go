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

// InitSetting
// Initialize settings, example:
//
//	 type tg struct {
//	   Token string `env:"token,required" json:"token"`
//	 }
//
//	 type baseSettings struct {
//		Tg  tg  `env:"tg"  json:"tg"  envPrefix:"tg__"`
//	 }
//	 func init() {
//		  ctx := logger_utils.NewLoggedCtx()
//
//		  Settings = must_utils.Must(settings_utils.InitSetting[baseSettings](
//		  	  ctx,
//		  	  "tg.token",
//		  ))
//	 }
//
//	 var Settings baseSettings
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
		settingsJson, err = sjson.SetBytes(
			settingsJson,
			valueKey,
			redact_utils.RedactWithPrefix(res.String()),
		)
		if err != nil {
			return *new(T), errors.Wrap(err, "failed to redact settings")
		}
	}

	zerolog.Ctx(ctx).
		Debug().
		RawJSON("v", settingsJson).
		Msg("settings.loaded")

	return settings, nil
}
