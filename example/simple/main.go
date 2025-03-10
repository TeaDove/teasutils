package main

import (
	"github.com/teadove/teasutils/utils/json_utils"
	"github.com/teadove/teasutils/utils/logger_utils"
	"github.com/teadove/teasutils/utils/settings_utils"
	"time"
)

type Settings struct {
	Host string `env:"HOST" envDefault:"127.0.0.1"`
}

func main() {
	ctx := logger_utils.NewLoggedCtx()
	settings := settings_utils.MustGetSetting[Settings](ctx, "SIMPLE_")

	idx := 0
	for {
		time.Sleep(2 * time.Second)
		idx++
		println(idx, string(json_utils.MarshalOrWarn(ctx, settings)))
	}

}
