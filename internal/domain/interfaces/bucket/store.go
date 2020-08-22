package bucket

// Storage provides a functionality to communicate with any bucket store that satisfies the interface
type Storage interface {
	// GetBucket returns the requested bucket to the callee
	GetBucket(name string) (Bucket, error)
	// CheckBucket checks whether a requested bucket is present in the active bucket store
	CheckBucket(name string) bool
	// AddBucket adds a new bucket to the active bucket store
	AddBucket(name string, b Bucket)
	// RemoveBucket removes a specified bucket from the active bucket store
	RemoveBucket(name string) error
}
