package config

import (
	"github.com/rs/zerolog"
	"go.soon.build/kit/config"
)

// application name
const AppName = "tarkov-tk-server-count"

// Config stores configuration options set by configuration file or env vars
type Config struct {
	Log      Log
	Discord  Discord
	BotToken string
}

// Log contains logging configuration
type Log struct {
	Console bool
	Verbose bool
	Level   string
}

type Discord struct {
	BotToken string
}

// Default is a default configuration setup with sane defaults
var Default = Config{
	Log: Log{
		Level: zerolog.InfoLevel.String(),
	},
	Discord:  Discord{},
	BotToken: "",
}

// New constructs a new Config instance
func New(opts ...config.Option) (Config, error) {
	c := Default
	v := config.ViperWithDefaults("ttsc")
	v.AutomaticEnv()
	err := config.ReadInConfig(v, &c, opts...)
	if err != nil {
		return c, err
	}
	return c, nil
}
