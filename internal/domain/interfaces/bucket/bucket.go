package bucket

// Bucket allows for the communication with any created bucket
type Bucket interface {
	// Decrement reduces the bucket request counter; it return true if the request can pass and false otherwise
	Decrement() bool
	// Stop releases all the resources associated with the bucket
	Stop()
}
