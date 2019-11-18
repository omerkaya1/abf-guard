package entity

import "github.com/omerkaya1/abf-guard/internal/domain/bucket"

// Password .
type Password struct {
	settings *bucket.Settings
	flash    chan struct{}
}

// NewPasswordBucket .
func NewPasswordBucket() *Password {
	return nil
}

// Add .
func (ib *Password) Add() error {
	return nil
}

// Flush .
func (ib *Password) Flush() error {
	return nil
}

// Delete .
func (ib *Password) Delete() error {
	return nil
}
