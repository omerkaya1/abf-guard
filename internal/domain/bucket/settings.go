package bucket

import (
	"strings"
	"time"

	"github.com/omerkaya1/abf-guard/internal/domain/config"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
)

// Settings is an object that holds all the main settings for the Bucket Manager
type Settings struct {
	LoginLimit    int
	PasswordLimit int
	IPLimit       int
	Expire        time.Duration
}

// InitBucketManagerSettings initiates setting for the bucket manager
func InitBucketManagerSettings(cfg config.Limits) (*Settings, error) {
	switch {
	case cfg.Login <= 0:
		return nil, errors.ErrIncorrectCfgLogin
	case cfg.Password <= 0:
		return nil, errors.ErrIncorrectCfgPWD
	case cfg.IP <= 0:
		return nil, errors.ErrIncorrectCfgIP
	case len(cfg.Expire) == 0 || !strings.ContainsAny(cfg.Expire, "msh"):
		return nil, errors.ErrEmptyCfgDuration
	}
	duration, err := time.ParseDuration(cfg.Expire)
	if err != nil {
		return nil, err
	}
	return &Settings{
		LoginLimit:    cfg.Login,
		PasswordLimit: cfg.Password,
		IPLimit:       cfg.IP,
		Expire:        duration,
	}, nil
}
