package config

import "path/filepath"

type Config struct {
	IsDebug             bool
	EnvPrefix           string
	JsonDebugFilePath   string
	JsonReleaseFilePath string
}

func createDefaultConfig(directory string) *Config {
	return &Config{
		EnvPrefix:    "",
		JsonDebugFilePath: filepath.Join(directory, "config.debug.json"),
		JsonReleaseFilePath: filepath.Join(directory, "config.json"),
	}
}
