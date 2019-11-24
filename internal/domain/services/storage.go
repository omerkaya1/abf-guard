package services

import (
	"context"

	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/db"
)

// Storage is the object that provides the means to communicate with a project's DB storage
type Storage struct {
	Processor db.StorageProcessor
}

// AddIP adds the IP to either the whitelist or the blacklist of the app
func (ss *Storage) AddIP(ctx context.Context, ip string, blacklist bool) error {
	return ss.Processor.Add(ctx, ip, blacklist)
}

// DeleteIP removes the IP from either the whitelist or the blacklist
func (ss *Storage) DeleteIP(ctx context.Context, ip string, blacklist bool) error {
	return ss.Processor.Delete(ctx, ip, blacklist)
}

// GetIPList either deletes the IP from the whitelist or the blacklist of the app
func (ss *Storage) GetIPList(ctx context.Context, blacklist bool) ([]string, error) {
	return ss.Processor.GetIPList(ctx, blacklist)
}

// GreenLightPass checks whether an ip exists in a specified list
func (ss *Storage) GreenLightPass(ctx context.Context, ip string) error {
	return ss.Processor.GreenLightPass(ctx, ip)
}
