package main

import (
	"time"

	"github.com/rs/zerolog"

	"github.com/teadove/teasutils/utils/logger_utils"
	"github.com/teadove/teasutils/utils/settings_utils"
)

type Settings struct {
	Host string `env:"HOST" envDefault:"127.0.0.1"`
}

func main() {
	ctx := logger_utils.NewLoggedCtx()
	settings := settings_utils.MustGetSetting[Settings](ctx, "SIMPLE_")

	idx := 0

	for {
		time.Sleep(time.Second)

		idx++
		zerolog.Ctx(ctx).Info().
			Int("idx", idx).
			Interface("settings", settings).
			Interface("base_settings", settings_utils.BaseSettings).
			Msg("hello world")
	}
}
