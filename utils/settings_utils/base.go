package settings_utils

import (
	"context"
	"os"
	"regexp"
	"time"
)

type logSettings struct {
	// Level
	// Can be trace, debug, info, warning, error etc.
	Level string `env:"LEVEL" envDefault:"INFO"`

	// Factory
	// Zerolog factory, can be console (with fancy colors ✨🪄🔮💫) or json (⚙️)
	Factory string `env:"FACTORY" envDefault:"CONSOLE"`
}

type profSettings struct {
	Enabled            bool          `env:"ENABLED"               envDefault:"false"`
	ResultFilename     string        `env:"RESULT_FILENAME"       envDefault:"cpu.prof"`
	SpamMemUsage       bool          `env:"SPAM_MEM_USAGE"        envDefault:"false"`
	SpamMemUsagePeriod time.Duration `env:"SPAM_MEM_USAGE_PERIOD" envDefault:"15s"`
}

type metricsSettings struct {
	URL            string        `env:"URL"             envDefault:"0.0.0.0:8093"`
	RequestTimeout time.Duration `env:"REQUEST_TIMEOUT" envDefault:"10s"`
	CloseTimeout   time.Duration `env:"CLOSE_TIMEOUT"   envDefault:"10s"`
}

type serviceSettings struct {
	Release     bool      `env:"RELEASE"      envDefault:"true"`
	StartedAt   time.Time `env:"START_TIME"   envDefault:""`
	ServiceName string    `env:"SERVICE_NAME" envDefault:""`

	Log     logSettings     `envPrefix:"LOG__"`
	Prof    profSettings    `envPrefix:"PROF__"`
	Metrics metricsSettings `envPrefix:"METRICS__"`
}

func (r *serviceSettings) Uptime() time.Duration {
	return time.Since(r.StartedAt)
}

//nolint:gochecknoglobals // need this
var ServiceSettings = MustGetSetting[serviceSettings](context.Background(), "BASE_")

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

func setStartedAt(settings *serviceSettings) {
	if settings.StartedAt.IsZero() {
		settings.StartedAt = time.Now().UTC()
	}
}

// nolint: gochecknoinits // required here
func init() {
	setServiceName(ServiceSettings)
	setStartedAt(ServiceSettings)
}
