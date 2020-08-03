package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {

	t.Run("Empty path", func(t *testing.T) {
		if cfg, err := InitConfig(""); assert.Error(t, err) {
			assert.Nil(t, cfg)
		}
	})
	t.Run("Path not found", func(t *testing.T) {
		if cfg, err := InitConfig("some/path/config.json"); assert.Error(t, err) {
			assert.Nil(t, cfg)
		}
	})
	t.Run("Incorrect file name", func(t *testing.T) {
		if cfg, err := InitConfig(" .json"); assert.Error(t, err) {
			assert.Nil(t, cfg)
		}
	})
}
