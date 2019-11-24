package services

import (
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/bucket"
)

// Bucket .
type Bucket struct {
	// Manager .
	Manager bucket.Manager
}

// Dispatch .
func (b Bucket) Dispatch(login, pwd, ip string) (bool, error) {
	return b.Manager.Dispatch(login, pwd, ip)
}

// FlushBuckets .
func (b Bucket) FlushBuckets(login, ip string) error {
	return b.Manager.FlushBuckets(login, ip)
}

// PurgeBucket .
func (b Bucket) PurgeBucket(name string) error {
	return b.Manager.PurgeBucket(name)
}

// MonitorErrors .
func (b Bucket) MonitorErrors() chan error {
	return b.Manager.GetErrChan()
}
