package bucket

import (
	"testing"
	"time"

	"github.com/omerkaya1/abf-guard/internal/domain/config"
	"github.com/stretchr/testify/require"
)

func TestInitBucketManagerSettings(t *testing.T) {
	testCases := []struct {
		header   string
		response error
		value    *Settings
		config.Limits
	}{
		{"Hours", nil, &Settings{
			LoginLimit:    10,
			PasswordLimit: 10,
			IPLimit:       100,
			Expire:        1 * time.Hour,
		}, config.Limits{Login: 10, Password: 10, IP: 100, Expire: "1h"}},
		{"Minutes", nil, &Settings{
			LoginLimit:    10,
			PasswordLimit: 10,
			IPLimit:       100,
			Expire:        1 * time.Minute,
		}, config.Limits{Login: 10, Password: 10, IP: 100, Expire: "1m"}},
	}
	for _, c := range testCases {
		t.Run(c.header, func(t *testing.T) {
			s, err := InitBucketManagerSettings(&c.Limits)
			require.Equal(t, c.response, err)
			require.Equal(t, c.value, s)
		})
	}
	t.Run("Invalid duration string", func(t *testing.T) {
		s, err := InitBucketManagerSettings(&config.Limits{Login: 10, Password: 10, IP: 100, Expire: "jhgl"})
		require.Error(t, err)
		require.Nil(t, nil, s)
	})
}

func TestSettings_Valid(t *testing.T) {
	tests := []struct {
		name string
		s    Settings
		want bool
	}{
		{"Empty settings", Settings{}, false},
		{"Partially empty settings", Settings{
			LoginLimit:    10,
			PasswordLimit: 15,
			IPLimit:       0,
			Expire:        0,
		}, false},
		{"Correct settings", Settings{
			LoginLimit:    5,
			PasswordLimit: 10,
			IPLimit:       100,
			Expire:        5,
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.s.Valid())
		})
	}
}
