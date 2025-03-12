package settings_utils

import (
	"os"
	"strconv"
	"time"
)

var (
	//nolint: gochecknoglobals // required
	envFilePath string
	//nolint: gochecknoglobals // required
	envFileRefreshEnabled bool
	//nolint: gochecknoglobals // required
	envFileRefreshInterval time.Duration
)

//nolint: gochecknoinits // required
func init() {
	envFilePath = os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = ".env"
	}

	refreshIntervalS, _ := strconv.Atoi(os.Getenv("ENV_REFRESH_INTERVAL_S"))
	if refreshIntervalS != 0 {
		envFileRefreshInterval = time.Duration(refreshIntervalS) * time.Second
		envFileRefreshEnabled = true
	}
}
