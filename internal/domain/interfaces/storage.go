package interfaces

import "context"

// TODO: comments!
// StorageProcessor .
type StorageProcessor interface {
	// Authenticate .
	Authenticate(context.Context, string, string, string) (bool, error)
	// Flash .
	Flash(context.Context, string, string, string) (bool, error)
	// Add .
	Add(context.Context, string, string, string) (bool, error)
	// Delete .
	Delete(context.Context, string, string, string) (bool, error)
}
