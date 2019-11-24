package bucket

// Manager .
type Manager interface {
	// Dispatch .
	Dispatch(string, string, string) (bool, error)
	// FlushBuckets .
	FlushBuckets(string, string) error
	// PurgeBucket .
	PurgeBucket(string) error
	// GetErrChan .
	GetErrChan() chan error
}
