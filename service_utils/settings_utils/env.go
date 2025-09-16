package settings_utils

import (
	"os"
	"syscall"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/teadove/teasutils/utils/must_utils"
)

// GetSettings
// Initialize settings, example:
//
//	 type tg struct {
//	   Token string `env:"TOKEN,required"`
//	 }
//
//	 type serviceSettings struct {
//		TG  tg  `env:"TG" envPrefix:"TG__"`
//	 }
//
//	 func init() {
//		  Settings = settings_utils.MustGetSetting[serviceSettings]("TEAS_")
//	 }
//
//	 var Settings serviceSettings
//
// Returns error if dotEnv file found, but corrupted.
func GetSettings[T any](envPrefix string) (T, error) {
	settings, err := loadSettingsWithDotenv[T](envPrefix)
	if err != nil {
		return *new(T), errors.Wrap(err, "failed to load settings")
	}

	return settings, nil
}

func MustGetSetting[T any](envPrefix string) T {
	return must_utils.Must(GetSettings[T](envPrefix))
}

func loadSettingsWithDotenv[T any](envPrefix string) (T, error) {
	// ! Dangerous place! Dotenv files will override any set ENV settings!
	err := godotenv.Overload(getEnvFilePath())
	if err != nil {
		var pathErr *os.PathError
		if !errors.As(err, &pathErr) || !errors.Is(pathErr.Err, syscall.ENOENT) {
			return *new(T), errors.Wrap(err, "load dotenv file")
		}
	}

	return loadSettings[T](envPrefix)
}

func loadSettings[T any](envPrefix string) (T, error) {
	settings, err := env.ParseAsWithOptions[T](env.Options{Prefix: envPrefix})
	if err != nil {
		return *new(T), errors.Wrap(err, "env parse")
	}

	return settings, nil
}
