package settings

import (
	"github.com/omerkaya1/abf-guard/internal/domain/config"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/stretchr/testify/assert"
	"time"

	"testing"
)

func TestInitBucketSettings(t *testing.T) {
	testCases := []struct {
		header   string
		response error
		value    *Settings
		config.Limits
	}{
		{"Incorrect login limit", errors.ErrIncorrectCfgLogin, nil, config.Limits{Login: -1, Password: 10, IP: 100, Expire: "1m"}},
		{"Incorrect password limit", errors.ErrIncorrectCfgPWD, nil, config.Limits{Login: 10, Password: -1, IP: 100, Expire: "1m"}},
		{"Incorrect ip limit", errors.ErrIncorrectCfgIP, nil, config.Limits{Login: 10, Password: 10, IP: -1, Expire: "1m"}},
		{"Incorrect expire string", errors.ErrEmptyCfgDuration, nil, config.Limits{Login: 10, Password: 10, IP: 100, Expire: ""}},
		{"All is well", nil, &Settings{
			LoginLimit:    10,
			PasswordLimit: 10,
			IPLimit:       100,
			Expire:        1 * time.Minute,
		}, config.Limits{Login: 10, Password: 10, IP: 100, Expire: "1m"}},
	}
	for _, c := range testCases {
		t.Run(c.header, func(t *testing.T) {
			s, err := InitBucketSettings(c.Limits)
			assert.Equal(t, c.response, err)
			assert.Equal(t, c.value, s)
		})
	}
	t.Run("Invalid duration string", func(t *testing.T) {
		limits := config.Limits{Login: 10, Password: 10, IP: 100, Expire: "jhgl"}
		if s, err := InitBucketSettings(limits); assert.Error(t, err) {
			assert.Nil(t, nil, s)
		}
	})
}
