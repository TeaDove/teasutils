package settings_utils

import (
	"context"
	"fmt"
	"os"
	"syscall"

	"github.com/teadove/teasutils/utils/json_utils"
	"github.com/teadove/teasutils/utils/redact_utils"

	"github.com/teadove/teasutils/utils/must_utils"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

const (
	envFile = ".env"
)

// InitSetting
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

	zerolog.Ctx(ctx).
		Debug().
		RawJSON("v", redact_utils.RedactJSONWithPrefix(
			ctx,
			json_utils.MarshalOrWarn(ctx, settings), omitFromLogValues...),
		).
		Msg("settings.loaded")

	return settings, nil
}

func MustInitSetting[T any](
	ctx context.Context,
	envPrefix string,
	omitFromLogValues ...string,
) T {
	return must_utils.Must(InitSetting[T](ctx, envPrefix, omitFromLogValues...))
}
