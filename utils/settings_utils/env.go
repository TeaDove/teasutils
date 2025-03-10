package settings_utils

import (
	"context"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/json_utils"
	"github.com/teadove/teasutils/utils/must_utils"
	"github.com/teadove/teasutils/utils/redact_utils"
	"os"
	"syscall"
	"time"
)

func loadSettings[T any](envPrefix string) (T, error) {
	// Dangerous place! Dotenv files will override any set ENV settings!
	err := godotenv.Overload(getFilePath())
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
func GetSettings[T any](
	ctx context.Context,
	envPrefix string,
	omitFromLogValues ...string,
) (*T, error) {
	lastLoad := time.Now().UTC()
	settings, err := loadSettings[T](envPrefix)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load settings")
	}

	var refreshPeriod time.Duration
	refreshPeriod, err = refresh(ctx, &settings, lastLoad, envPrefix)
	if err != nil {
		return nil, errors.Wrap(err, "failed to schedule refresher")
	}

	prelog := zerolog.Ctx(ctx).
		Debug().
		RawJSON("v", redact_utils.RedactJSONWithPrefix(
			ctx,
			json_utils.MarshalOrWarn(ctx, settings), omitFromLogValues...),
		)

	if refreshPeriod > 0 {
		prelog = prelog.Str("refresh_period", refreshPeriod.String())
	}

	prelog.Msg("settings.loaded")

	return &settings, nil
}

func MustGetSetting[T any](
	ctx context.Context,
	envPrefix string,
	omitFromLogValues ...string,
) *T {
	return must_utils.Must(GetSettings[T](ctx, envPrefix, omitFromLogValues...))
}
