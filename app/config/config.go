package config

import (
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type DatabaseDetail struct {
	URL      string `yaml:"url"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	InsertCount    int            `yaml:"insertCount"`
	DatabaseConfig DatabaseDetail `yaml:"databaseConfig"`
}

func ParseConfig(configFilePath string, logger *zap.Logger) *Config {
	var config Config
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		logger.Error("Error in reading file ", zap.Error(err))
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logger.Error("Error in unmarshiling config object", zap.Error(err))
	}
	return &config
}
