package settings_utils

import (
	"context"
	"io/fs"
	"os"
	"strconv"
	"time"

	"github.com/wI2L/jsondiff"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func refresh[T any]( //nolint: gocognit // TODO refactor
	ctx context.Context,
	settings *T,
	loadedAt time.Time,
	envPrefix string,
) (time.Duration, error) {
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
		ticker      = time.NewTicker(period)
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

			if file.ModTime() == loadedAt {
				continue
			}

			newSettings, err = loadSettings[T](envPrefix)
			if err != nil {
				zerolog.Ctx(ctx).Error().Stack().Err(err).Msg("env.setting.error")
			}

			var patch jsondiff.Patch

			patch, err = jsondiff.Compare(settings, newSettings)
			if err != nil {
				zerolog.Ctx(ctx).
					Error().
					Stack().Err(err).
					Msg("err.checking.diff")
			}

			if len(patch) != 0 || err != nil {
				*settings = newSettings

				zerolog.Ctx(ctx).
					Info().
					Str("file", filePath).
					RawJSON("patch", []byte(patch.String())).
					Msg("settings.refreshed")
			}

			loadedAt = file.ModTime()
		}
	}()

	return period, nil
}
