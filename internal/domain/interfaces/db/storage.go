package db

import "context"

// Storage is an interface to communicate with the DB
type Storage interface {
	// Add adds the IP to either the whitelist or the blacklist of the app
	Add(context.Context, string, bool) error
	// Delete removes the IP from either the whitelist or the blacklist
	Delete(context.Context, string, bool) error
	// GetIPList either deletes the IP from the whitelist or the blacklist of the app
	GetIPList(context.Context, bool) ([]string, error)
	// ExistInList checks whether an ip exists in a specified list
	GreenLightPass(context.Context, string) error
}
