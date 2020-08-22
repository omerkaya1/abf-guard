package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitConfig(t *testing.T) {
	t.Run("Empty path", func(t *testing.T) {
		cfg, err := InitConfig("")
		require.Error(t, err)
		require.Nil(t, cfg)
	})
	t.Run("Path not found", func(t *testing.T) {
		cfg, err := InitConfig("some/path/config.json")
		require.Error(t, err)
		require.Nil(t, cfg)
	})
	t.Run("Incorrect file name", func(t *testing.T) {
		cfg, err := InitConfig(" .json")
		require.Error(t, err)
		require.Nil(t, cfg)
	})
}

func TestDBConf_Valid(t *testing.T) {
	tests := []struct {
		name string
		cfg  DBConf
		want bool
	}{
		{name: "Empty config", cfg:  DBConf{}, want: false},
		{name: "Partially empty config", cfg:  DBConf{Host: "fsadf", Port: "sdfsd"}, want: false},
		{name: "Correct config", cfg:  DBConf{
			Host:     "localhost",
			Port:     "5432",
			Password: "some_pwd",
			Name:     "some_name",
			User:     "some_user",
			SSL:      "disabled",
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.cfg.Valid())
		})
	}
}

func TestLimits_Valid(t *testing.T) {
	tests := []struct {
		name string
		cfg  Limits
		want bool
	}{
		{ name: "Empty config", cfg:  Limits{}, want: false},
		{ name: "Partially empty config", cfg:  Limits{IP: 3, Login: 3}, want: false},
		{ name: "Correct config", cfg:  Limits{
			Login:    5,
			Password: 4,
			IP:       3,
			Expire:   "3m",
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.cfg.Valid())
		})
	}
}

func TestServer_Valid(t *testing.T) {
	tests := []struct {
		name string
		cfg  Server
		want bool
	}{
		{ name: "Empty config", cfg:  Server{}, want: false},
		{ name: "Partially empty config", cfg:  Server{Host: "localhost"}, want: false},
		{ name: "Correct config", cfg:  Server{
			Host:  "localhost",
			Port:  "8080",
			Level: 1,
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.cfg.Valid())
		})
	}
}
