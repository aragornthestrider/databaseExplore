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
}
