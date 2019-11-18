package entities

import "github.com/omerkaya1/abf-guard/internal/domain/bucket"

type LoginBucket struct {
	settings *bucket.Settings
	flash    chan struct{}
}

func NewLoginBucket() *LoginBucket {
	return nil
}

func (ib *LoginBucket) Add() error {
	return nil
}

func (ib *LoginBucket) Flush() error {
	return nil
}

func (ib *LoginBucket) Delete() error {
	return nil
}
