package bucket

// Manager is a representation of a Bucket Manager interface
type Manager interface {
	// Dispatch accepts authorisation request parameters and creates a new or decrements a counter for each bucket
	Dispatch(login string, pwd string, ip string) (bool, error)
	// FlushBuckets removes all buckets with the specified login and ip
	FlushBuckets(login string, ip string) error
	// PurgeBucket removes a bucket which name was specified as an argument
	PurgeBucket(name string) error
	// GetErrChan returns an error channel to monitor the Manager's activity
	GetErrChan() chan error
}
