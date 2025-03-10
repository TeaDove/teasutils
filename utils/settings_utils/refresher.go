package settings_utils

import (
	"context"
	"io/fs"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func refresh[T any](ctx context.Context, settings *T, loadedAt time.Time, envPrefix string) (time.Duration, error) {
	refreshIntervalSRaw, ok := os.LookupEnv(envFilePathRefreshIntervalS)
	if !ok {
		return 0, nil
	}

	filePath := getFilePath()

	file, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return 0, nil
		}

		return 0, errors.Wrap(err, "failed to check if file exists")
	}

	refreshIntervalS, err := strconv.Atoi(refreshIntervalSRaw)
	if err != nil {
		return 0, errors.Wrap(err, "invalid refresh interval")
	}

	var (
		newSettings T
		period      = time.Duration(refreshIntervalS) * time.Second
		ticker      = time.NewTimer(period)
	)

	go func() {
		for range ticker.C {
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

			newSettings, err = loadSettings[T](envPrefix)
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
	}()

	return period, nil
}
