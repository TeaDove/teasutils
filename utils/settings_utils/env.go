package settings_utils

import (
	"context"
	"encoding/json"
	"os"
	"syscall"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/must_utils"
	"github.com/teadove/teasutils/utils/redact_utils"
)

func loadSettings[T any](envPrefix string) (T, error) {
	// Dangerous place! Dotenv files will override any set ENV settings!
	err := godotenv.Overload(getEnvFilePath())
	if err != nil {
		var pathErr *os.PathError
		if !errors.As(err, &pathErr) || !errors.Is(pathErr.Err, syscall.ENOENT) {
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
//		 type serviceSettings struct {
//			TG  tg  `env:"TG" envPrefix:"TG__"`
//		 }
//		 func init() {
//			  ctx := logger_utils.NewLoggedCtx()
//
//			  Settings = must_utils.Must(settings_utils.InitSetting[serviceSettings](
//			  	  ctx,
//	           "TEAS_",
//			  	  "TG.Token",
//			  ))
//		 }
//
//		 var Settings serviceSettings
//
// Panics if dotEnv file found, but corrupted.
func GetSettings[T any](ctx context.Context, envPrefix string) (*T, error) {
	lastLoad := time.Now().UTC()

	settings, err := loadSettings[T](envPrefix)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load settings")
	}

	marshaledSettings, err := json.Marshal(settings)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal settings")
	}

	prelog := zerolog.Ctx(ctx).
		Debug().
		RawJSON("v", redact_utils.RedactLongStrings(marshaledSettings))

	if getEnvFileRefreshEnabled() {
		go refresh(ctx, &settings, lastLoad, envPrefix)
		prelog.Str("refresh_period", getEnvFileRefreshInterval().String())
	}

	prelog.Msg("settings.loaded")

	return &settings, nil
}

func MustGetSetting[T any](ctx context.Context, envPrefix string) *T {
	return must_utils.Must(GetSettings[T](ctx, envPrefix))
}
