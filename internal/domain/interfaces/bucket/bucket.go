package bucket

// Bucket .
type Bucket interface {
	// Decrement .
	Decrement() bool
	// Stop .
	Stop()
}
