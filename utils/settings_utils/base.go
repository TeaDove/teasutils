package settings_utils

import (
	"context"
	"github.com/pkg/errors"
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
	SpamMemUsage       bool          `env:"SPAM_MEM_USAGE"        envDefault:"true"`
	SpamMemUsagePeriod time.Duration `env:"SPAM_MEM_USAGE_PERIOD" envDefault:"1s"`
}

type metricsSettings struct {
	URL            string        `env:"URL"             envDefault:"0.0.0.0:8083"`
	RequestTimeout time.Duration `env:"REQUEST_TIMEOUT" envDefault:"10s"`
	CloseTimeout   time.Duration `env:"CLOSE_TIMEOUT"   envDefault:"10s"`
}

type baseSettings struct {
	Release     bool      `env:"RELEASE"    envDefault:"true"`
	StartedAt   time.Time `env:"START_TIME" envDefault:""`
	ServiceName string    `env:"SERVICE_NAME" envDefault:""`

	Log     logSettings     `envPrefix:"LOG__"`
	Prof    profSettings    `envPrefix:"PROF__"`
	Metrics metricsSettings `envPrefix:"METRICS__"`
}

func (r *baseSettings) Uptime() time.Duration {
	return time.Since(r.StartedAt)
}

//nolint:gochecknoglobals // need this
var BaseSettings = MustInitSetting[baseSettings](context.Background(), "BASE_")

func setServiceName(settings *baseSettings) {
	if settings.ServiceName != "" {
		return
	}

	hostName := os.Getenv("HOSTNAME")
	if hostName == "" {
		settings.ServiceName = "undefined"
		return
	}

	kubepodNameRegexp, err := regexp.Compile(`^(.+)-\w+-\w+$`)
	if err != nil {
		panic(errors.Wrap(err, "failed to compile kubepod name regexp"))
	}

	foundString := kubepodNameRegexp.FindStringSubmatch(hostName)
	if len(foundString) == 2 {
		settings.ServiceName = foundString[1]
		return
	}

	settings.ServiceName = hostName
}

func setStartedAt(settings *baseSettings) {
	if settings.StartedAt.IsZero() {
		settings.StartedAt = time.Now().UTC()
	}
}

// nolint: gochecknoinits // required here
func init() {
	setServiceName(&BaseSettings)
	setStartedAt(&BaseSettings)
}
