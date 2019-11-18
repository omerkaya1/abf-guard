package entity

import "github.com/omerkaya1/abf-guard/internal/domain/bucket"

// IP .
type IP struct {
	settings *bucket.Settings
	flash    chan struct{}
}

// NewIPBucket .
func NewIPBucket() *IP {
	return nil
}

// Add .
func (i *IP) Add() error {
	return nil
}

// Flush .
func (i *IP) Flush() error {
	return nil
}

// Delete .
func (i *IP) Delete() error {
	return nil
}
