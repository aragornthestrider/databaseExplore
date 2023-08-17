package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestConfig(t *testing.T) {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		t.Error("Error in reading file: ", err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		t.Error("Error in unmarshiling config object: ", err)
	}
	assert.Equal(t, config.InsertCount, 100)
	assert.Equal(t, config.DatabaseConfig.URL, "postgres")
	assert.Equal(t, config.DatabaseConfig.Port, 5432)
	assert.Equal(t, config.DatabaseConfig.Database, "postgres")
	assert.Equal(t, config.DatabaseConfig.Username, "postgres")
	assert.Equal(t, config.DatabaseConfig.Password, "postgres")
}
