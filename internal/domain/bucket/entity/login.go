package entity

import "github.com/omerkaya1/abf-guard/internal/domain/bucket"

// Login .
type Login struct {
	settings *bucket.Settings
	flash    chan struct{}
}

// NewLoginBucket .
func NewLoginBucket() *Login {
	return nil
}

// Add .
func (ib *Login) Add() error {
	return nil
}

// Flush .
func (ib *Login) Flush() error {
	return nil
}

// Delete .
func (ib *Login) Delete() error {
	return nil
}
