package settings_utils

import (
	"context"
	"time"
)

type LogSettings struct {
	// Level
	// Can be trace, debug, info, warning, error etc.
	Level string `env:"LEVEL" envDefault:"INFO"`

	// Factory
	// Zerolog factory, can be console (with fancy colors ✨🪄🔮💫) or json (⚙️)
	Factory string `env:"FACTORY" envDefault:"CONSOLE"`
}

type ProfSettings struct {
	Enabled            bool          `env:"ENABLED"               envDefault:"false"`
	ResultFilename     string        `env:"RESULT_FILENAME"       envDefault:"cpu.prof"`
	SpamMemUsage       bool          `env:"SPAM_MEM_USAGE"        envDefault:"true"`
	SpamMemUsagePeriod time.Duration `env:"SPAM_MEM_USAGE_PERIOD" envDefault:"1s"`
}

type MetricsSettings struct {
	URL            string        `env:"URL"             envDefault:"0.0.0.0:8083"`
	RequestTimeout time.Duration `env:"REQUEST_TIMEOUT" envDefault:"10s"`
}

type baseSettings struct {
	Release bool `env:"RELEASE" envDefault:"true"`

	Log     LogSettings     `envPrefix:"LOG__"`
	Prof    ProfSettings    `envPrefix:"PROF__"`
	Metrics MetricsSettings `envPrefix:"METRICS__"`
}

//nolint:gochecknoglobals // need this
var BaseSettings = MustInitSetting[baseSettings](context.Background(), "BASE_")
