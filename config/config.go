package config

import "github.com/BurntSushi/toml"

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
	MigPath     string `toml:"mig_path"`
	Root        string `toml:"root"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}

func DecodeConfigFromRelativePath(path string) (*Config, error) {
	config := NewConfig()

	_, err := toml.DecodeFile(path, config)
	if err != nil {
		return nil, err
	}
	return config, nil

}
