package bucket

// Storage .
type Storage interface {
	// GetBucket .
	GetBucket(string) (Bucket, error)
	// CheckBucket .
	CheckBucket(string) bool
	// AddBucket .
	AddBucket(string, Bucket)
	// RemoveBucket .
	RemoveBucket(string) error
}
