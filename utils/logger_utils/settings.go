package logger_utils

type settings struct {
	// Level
	// Can be trace, debug, info, warning, error etc.
	Level string `env:"level" envDefault:"DEBUG"`

	// Factory
	// Zerolog factory, can be console (with fancy colors ✨🪄🔮💫) or json (⚙️)
	Factory string `env:"factory" envDefault:"CONSOLE"`
}
