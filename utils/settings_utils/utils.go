package settings_utils

import (
	"os"
)

const (
	envFilePathEnv              = "ENV_FILE_PATH"
	envFilePathRefreshIntervalS = "ENV_REFRESH_INTERVAL_S"
)

func getFilePath() string {
	envFile, ok := os.LookupEnv(envFilePathEnv)
	if !ok {
		envFile = ".env"
	}

	return envFile
}
