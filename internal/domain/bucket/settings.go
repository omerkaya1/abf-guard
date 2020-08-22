package bucket

import (
	"time"

	"github.com/omerkaya1/abf-guard/internal/domain/config"
)

// Settings is an object that holds all the main settings for the Bucket Manager
type Settings struct {
	LoginLimit    int
	PasswordLimit int
	IPLimit       int
	Expire        time.Duration
}

// Valid validates the bucket manager settings
func (s Settings) Valid() bool {
	return s.IPLimit > 0 && s.LoginLimit > 0 && s.PasswordLimit > 0 && s.Expire > 0
}

// InitBucketManagerSettings initiates setting for the bucket manager
func InitBucketManagerSettings(cfg *config.Limits) (*Settings, error) {
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
