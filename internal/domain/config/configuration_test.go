package config

import (
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitConfig(t *testing.T) {

	t.Run("Empty path", func(t *testing.T) {
		if cfg, err := InitConfig(""); assert.Error(t, err) {
			assert.Nil(t, cfg)
		}
	})
	t.Run("Incorrect file extension", func(t *testing.T) {
		if cfg, err := InitConfig(""); assert.Equal(t, errors.ErrCorruptConfigFileExtension, err) {
			assert.Nil(t, cfg)
		}
	})
	t.Run("Incorrect file name", func(t *testing.T) {
		if cfg, err := InitConfig(" .json"); assert.Error(t, err) {
			assert.Nil(t, cfg)
		}
	})
}
