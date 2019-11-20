package services

import (
	"context"
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces"
)

// Storage .
type Storage struct {
	Processor interfaces.StorageProcessor
}

// AddIP .
func (ss *Storage) AddIP(ctx context.Context, ip string, blacklist bool) error {
	return nil
}

// DeleteIP .
func (ss *Storage) DeleteIP(ctx context.Context, ip string, blacklist bool) error {
	return nil
}

// GetIPList .
func (ss *Storage) GetIPList(ctx context.Context, blacklist bool) ([]string, error) {
	return nil, nil
}
