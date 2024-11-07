package settings_utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/redact_utils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"os"
	"syscall"
)

const (
	envFile = ".env"
)

// InitSetting
// Initialize settings, example:
//
//		 type tg struct {
//		   Token string `env:"token,required" json:"token"`
//		 }
//
//		 type baseSettings struct {
//			Tg  tg  `env:"tg"  json:"tg"  envPrefix:"tg__"`
//		 }
//		 func init() {
//			  ctx := logger_utils.NewLoggedCtx()
//
//			  Settings = must_utils.Must(settings_utils.InitSetting[baseSettings](
//			  	  ctx,
//	           "teas_",
//			  	  "tg.token",
//			  ))
//		 }
//
//		 var Settings baseSettings
//
// Panics if dotEnv file found, but corrupted.
func InitSetting[T any](
	ctx context.Context,
	envPrefix string,
	omitFromLogValues ...string,
) (T, error) {
	err := godotenv.Load(envFile)
	if err != nil {
		var pathErr *os.PathError
		if !(errors.As(err, &pathErr) && errors.Is(pathErr.Err, syscall.ENOENT)) {
			panic(fmt.Sprintf("failed to load dotenv file %s: %v", envFile, err))
		}
	}

	settings, err := env.ParseAsWithOptions[T](env.Options{Prefix: envPrefix})
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
