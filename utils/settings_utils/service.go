package settings_utils

import (
	"context"
	"os"
	"regexp"
)

type logSettings struct {
	// Level
	// Can be trace, debug, info, warning, error etc.
	Level string `env:"LEVEL" envDefault:"INFO"`

	// Factory
	// Zerolog factory, can be console (with fancy colors ✨🪄🔮💫) or json (⚙️)
	Factory string `env:"FACTORY" envDefault:"CONSOLE"`
}

type serviceSettings struct {
	Release     bool   `env:"RELEASE"      envDefault:"true"`
	ServiceName string `env:"SERVICE_NAME" envDefault:""`

	Log logSettings `envPrefix:"LOG__"`
}

//nolint:gochecknoglobals // need this
var ServiceSettings serviceSettings

// nolint: gochecknoinits // required here
func init() {
	ServiceSettings = MustGetSetting[serviceSettings](context.Background(), "SERVICE_")
	setServiceName(&ServiceSettings)
}

// TODO add hooks
func setServiceName(settings *serviceSettings) {
	if settings.ServiceName != "" {
		return
	}

	hostName := os.Getenv("HOSTNAME")
	if hostName == "" {
		settings.ServiceName = "undefined"
		return
	}

	foundString := regexp.MustCompile(`^(.+)-\w+-\w+$`).FindStringSubmatch(hostName)
	// TODO make for sfs

	const maxAllowedGroups = 2

	if len(foundString) == maxAllowedGroups {
		settings.ServiceName = foundString[1]
		return
	}

	settings.ServiceName = hostName
}
