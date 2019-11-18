package entities

import "github.com/omerkaya1/abf-guard/internal/domain/bucket"

type IpBucket struct {
	settings *bucket.Settings
	flash    chan struct{}
}

func NewIpBucket() *IpBucket {
	return nil
}

func (ib *IpBucket) Add() error {
	return nil
}

func (ib *IpBucket) Flush() error {
	return nil
}

func (ib *IpBucket) Delete() error {
	return nil
}
