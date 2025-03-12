package settings_utils

import (
	"context"
	"os"
	"syscall"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/json_utils"
	"github.com/teadove/teasutils/utils/must_utils"
	"github.com/teadove/teasutils/utils/redact_utils"
)

func loadSettings[T any](envPrefix string) (T, error) {
	// Dangerous place! Dotenv files will override any set ENV settings!
	err := godotenv.Overload(envFilePath)
	if err != nil {
		var pathErr *os.PathError
		if !(errors.As(err, &pathErr) && errors.Is(pathErr.Err, syscall.ENOENT)) {
			return *new(T), errors.Wrap(err, "failed to load dotenv file")
		}
	}

	settings, err := env.ParseAsWithOptions[T](env.Options{Prefix: envPrefix})
	if err != nil {
		return *new(T), errors.Wrap(err, "failed to env parse")
	}

	return settings, nil
}

// GetSettings
// Initialize settings, example:
//
//		 type tg struct {
//		   Token string `env:"TOKEN,required"`
//		 }
//
//		 type baseSettings struct {
//			TG  tg  `env:"TG" envPrefix:"TG__"`
//		 }
//		 func init() {
//			  ctx := logger_utils.NewLoggedCtx()
//
//			  Settings = must_utils.Must(settings_utils.InitSetting[baseSettings](
//			  	  ctx,
//	           "TEAS_",
//			  	  "TG.Token",
//			  ))
//		 }
//
//		 var Settings baseSettings
//
// Panics if dotEnv file found, but corrupted.
func GetSettings[T any](ctx context.Context, envPrefix string) (*T, error) {
	lastLoad := time.Now().UTC()

	settings, err := loadSettings[T](envPrefix)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load settings")
	}

	prelog := zerolog.Ctx(ctx).
		Debug().
		RawJSON("v", redact_utils.RedactLongStrings(json_utils.MarshalOrWarn(ctx, settings)))

	if envFileRefreshEnabled {
		go refresh(ctx, &settings, lastLoad, envPrefix)
		prelog.Str("refresh_period", envFileRefreshInterval.String())
	}

	prelog.Msg("settings.loaded")

	return &settings, nil
}

func MustGetSetting[T any](ctx context.Context, envPrefix string) *T {
	return must_utils.Must(GetSettings[T](ctx, envPrefix))
}
