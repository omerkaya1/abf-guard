package interfaces

import "context"

// StorageProcessor is an interface to communicate with the DB
type StorageProcessor interface {
	// Add adds the IP to either the whitelist or the blacklist of the app
	Add(context.Context, bool, string) error
	// Delete removes the IP from either the whitelist or the blacklist
	Delete(context.Context, bool, string) error
	// GetIpList either deletes the IP from the whitelist or the blacklist of the app
	GetIpList(context.Context, bool) ([]string, error)
}
