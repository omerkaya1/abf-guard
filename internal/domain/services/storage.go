package services

import (
	"context"
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/db"
)

// Storage .
type Storage struct {
	Processor db.StorageProcessor
}

// AddIP .
func (ss *Storage) AddIP(ctx context.Context, ip string, blacklist bool) error {
	// TODO: Add verifiers!
	return ss.Processor.Add(ctx, ip, blacklist)
}

// DeleteIP .
func (ss *Storage) DeleteIP(ctx context.Context, ip string, blacklist bool) error {
	// TODO: Add verifiers!
	return ss.Processor.Delete(ctx, ip, blacklist)
}

// GetIPList .
func (ss *Storage) GetIPList(ctx context.Context, blacklist bool) ([]string, error) {
	return ss.Processor.GetIPList(ctx, blacklist)
}

// GreenLightPass .
func (ss *Storage) GreenLightPass(ctx context.Context, ip string) error {
	return ss.Processor.GreenLightPass(ctx, ip)
}
