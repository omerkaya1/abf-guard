package entities

import "github.com/omerkaya1/abf-guard/internal/domain/bucket"

type PwdBucket struct {
	settings *bucket.Settings
	flash    chan struct{}
}

func NewPwdBucket() *PwdBucket {
	return nil
}

func (ib *PwdBucket) Add() error {
	return nil
}

func (ib *PwdBucket) Flush() error {
	return nil
}

func (ib *PwdBucket) Delete() error {
	return nil
}
