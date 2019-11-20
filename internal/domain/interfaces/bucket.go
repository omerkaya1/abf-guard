package interfaces

import "context"

// Bucket .
type Bucket interface {
	// Start .
	Start(context.Context, chan string, chan string)
	// Decrement .
	Decrement()
}
