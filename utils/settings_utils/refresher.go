package settings_utils

import (
	"context"
	"os"
	"time"

	"github.com/wI2L/jsondiff"

	"github.com/rs/zerolog"
)

//nolint: gochecknoglobals, mnd // required
var Refreshed = make(chan struct{}, 10)

func refresh[T any](
	ctx context.Context,
	settings *T,
	loadedAt time.Time,
	envPrefix string,
) {
	var (
		newSettings T
		period      = envFileRefreshInterval
		ticker      = time.NewTicker(period)
		file        os.FileInfo
		err         error
	)

	for range ticker.C {
		file, err = os.Stat(envFilePath)
		if err != nil {
			zerolog.Ctx(ctx).
				Error().Stack().
				Err(err).
				Str("file", envFilePath).
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
				Str("file", envFilePath).
				RawJSON("patch", []byte(patch.String())).
				Msg("settings.refreshed")

			if len(Refreshed) < cap(Refreshed) {
				Refreshed <- struct{}{}
			}
		}

		loadedAt = file.ModTime()
	}
}
