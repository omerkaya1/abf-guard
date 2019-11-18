package interfaces

// Bucket .
type Bucket interface {
	// Flush .
	Flush() error
	// Add .
	Add() error
	// Delete .
	Delete() error
}
