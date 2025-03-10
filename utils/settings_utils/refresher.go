package settings_utils

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"io/fs"
	"os"
	"time"
)

func refresh[T any](ctx context.Context, settings *T, loadedAt time.Time, envPrefix string) (time.Duration, error) {
	refreshEnabled := os.Getenv(envFilePathRefreshEnabled)
	if refreshEnabled != "true" {
		return 0, nil
	}

	filePath := getFilePath()
	file, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return 0, nil
		} else {
			return 0, errors.Wrap(err, "failed to check if file exists")
		}
	}

	period, err := getRefreshInterval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get refresh interval")
	}

	go func() {
		ticker := time.NewTimer(period)
		for {
			select {
			case <-ticker.C:
				file, err = os.Stat(filePath)
				if err != nil {
					zerolog.Ctx(ctx).
						Error().Stack().
						Err(err).
						Str("file", filePath).
						Msg("failed.to.open.file")
					continue
				}

				if file.ModTime().Before(loadedAt) {
					continue
				}

				newSettings, err := loadSettings[T](envPrefix)
				if err != nil {
					zerolog.Ctx(ctx).Error().Stack().Err(err).Msg("env.setting.error")
				}
				*settings = newSettings

				// TODO log only diff
				zerolog.Ctx(ctx).
					Info().
					Str("file", filePath).
					Interface("new_settings", newSettings).
					Msg("settings.refreshed")
			}
		}
	}()

	return period, nil
}
