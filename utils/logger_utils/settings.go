package logger_utils

type settings struct {
	// LogLevel
	// Can be trace, debug, info, warning, error etc.
	LogLevel string `env:"log_level" envDefault:"debug"`

	// LoggerFactory
	// Zerolog factory, can be console (with fancy colors ✨🪄🔮💫) or json (⚙️)
	LoggerFactory string `env:"logger_factory" envDefault:"console"`
}
