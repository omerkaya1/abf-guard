package db

import "context"

// Storage is an interface to communicate with the DB
type Storage interface {
	// Add adds the IP to either the whitelist or the blacklist of the app
	Add(ctx context.Context, ip string, blacklist bool) error
	// Delete removes the IP from either the whitelist or the blacklist
	Delete(ctx context.Context, ip string, blacklist bool) error
	// GetIPList either deletes the IP from the whitelist or the blacklist of the app
	GetIPList(ctx context.Context, blacklist bool) ([]string, error)
	// ExistInList checks whether an ip exists in a specified list
	GreenLightPass(ctx context.Context, ip string) error
}
