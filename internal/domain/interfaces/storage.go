package interfaces

import "context"

// StorageProcessor is an interface to communicate with the DB
type StorageProcessor interface {
	// Authenticate tries to authenticate the user to the resource
	Authenticate(context.Context, string, string, string) (bool, error)
	// Flash resets the selected bucket
	Flash(context.Context, string, string, string) (bool, error)
	// Add either adds the IP to the whitelist or the blacklist of the app
	Add(context.Context, string, string, string) (bool, error)
	// Delete either deletes the IP from the whitelist or the blacklist of the app
	Delete(context.Context, string, string, string) (bool, error)
}
