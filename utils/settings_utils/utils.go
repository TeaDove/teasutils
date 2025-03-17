package settings_utils

import (
	"os"
	"strconv"
	"time"
)

func getEnvFilePath() string {
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = ".env"
	}

	return envFilePath
}

func getEnvFileRefreshEnabled() bool {
	refreshIntervalS, _ := strconv.Atoi(os.Getenv("ENV_REFRESH_INTERVAL_S"))
	return refreshIntervalS != 0
}

func getEnvFileRefreshInterval() time.Duration {
	refreshIntervalS, _ := strconv.Atoi(os.Getenv("ENV_REFRESH_INTERVAL_S"))

	return time.Duration(refreshIntervalS) * time.Second
}
