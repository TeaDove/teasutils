package settings_utils

import (
	"os"
)

func getEnvFilePath() string {
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = ".env"
	}

	return envFilePath
}
