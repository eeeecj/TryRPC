package config

import "github.com/TryRpc/internal/local/options"

type Config struct {
	*options.Options
}

func NewConfig(opt *options.Options) *Config {
	return &Config{opt}
}
