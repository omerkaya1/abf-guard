package services

import (
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/bucket"
)

// Bucket is a service that provides all the needed functionality to manage buckets
type Bucket struct {
	// Manager provides means to communicate with the a bucket manager object
	Manager bucket.Manager
}

// Dispatch accepts authorisation request parameters and creates a new or decrements a counter for each bucket
func (b Bucket) Dispatch(login, pwd, ip string) (bool, error) {
	return b.Manager.Dispatch(login, pwd, ip)
}

// FlushBuckets removes all buckets with the specified login and ip
func (b Bucket) FlushBuckets(login, ip string) error {
	return b.Manager.FlushBuckets(login, ip)
}

// PurgeBucket removes a bucket which name was specified as an argument
func (b Bucket) PurgeBucket(name string) error {
	return b.Manager.PurgeBucket(name)
}

// MonitorErrors returns an error channel to monitor the bucket manager's activity
func (b Bucket) MonitorErrors() chan error {
	return b.Manager.GetErrChan()
}
