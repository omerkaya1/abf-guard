package grpc

import (
	"github.com/omerkaya1/abf-guard/internal/domain/config"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/domain/services"
	"github.com/omerkaya1/abf-guard/log"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewServer(t *testing.T) {
	testCases := []struct {
		header   string
		response error
		cfg      *config.Server
		logger   *zap.Logger
		ss       *services.Storage
		bm       *services.Bucket
	}{
		{
			"No configuration",
			errors.ErrMissingServerParameters,
			nil,
			&zap.Logger{},
			&services.Storage{Processor: nil},
			&services.Bucket{Manager: nil},
		},
		{
			"No logger",
			errors.ErrMissingServerParameters,
			&config.Server{},
			nil,
			&services.Storage{Processor: nil},
			&services.Bucket{Manager: nil},
		},
		{
			"No storage",
			errors.ErrMissingServerParameters,
			&config.Server{},
			&zap.Logger{},
			nil,
			&services.Bucket{Manager: nil},
		},
		{
			"No bucket manager",
			errors.ErrMissingServerParameters,
			&config.Server{},
			&zap.Logger{},
			&services.Storage{Processor: nil},
			nil,
		},
	}
	for _, c := range testCases {
		t.Run(c.header, func(t *testing.T) {
			if s, err := NewServer(c.cfg, c.logger, c.ss, c.bm); assert.Error(t, err) {
				assert.Nil(t, s)
			}
		})
	}
	t.Run("Correct parameters", func(t *testing.T) {
		cfg := config.Server{
			Host:  "127.0.0.1",
			Port:  "8080",
			Level: 1,
		}
		l, err := log.InitLogger(cfg.Level)
		assert.NoError(t, err)
		if s, err := NewServer(&cfg, l, &services.Storage{Processor: nil}, &services.Bucket{Manager: nil}); assert.NoError(t, err) {
			assert.NotNil(t, s)
		}
	})
}
