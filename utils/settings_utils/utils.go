package settings_utils

import (
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const (
	envFilePathEnv              = "ENV_FILE_PATH"
	envFilePathRefreshIntervalS = "ENV_REFRESH_INTERVAL_S"
	envFilePathRefreshEnabled   = "ENV_REFRESH_ENABLED"
)

func getFilePath() string {
	envFile, ok := os.LookupEnv("ENV_FILE_PATH")
	if !ok {
		envFile = ".env"
	}

	return envFile
}

func getRefreshInterval() (time.Duration, error) {
	var (
		refreshIntervalS = 10
		err              error
	)

	refreshIntervalSRaw, ok := os.LookupEnv(envFilePathRefreshIntervalS)
	if ok {
		refreshIntervalS, err = strconv.Atoi(refreshIntervalSRaw)
		if err != nil {
			return 0, errors.Wrap(err, "invalid refresh interval")
		}
	}

	return time.Duration(refreshIntervalS) * time.Second, nil
}
