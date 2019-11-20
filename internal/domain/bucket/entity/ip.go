package entity

import (
	"context"
	"sync"
)

// IP .
type IP struct {
	limit int
	name  string
	stop  chan struct{}
	m     sync.RWMutex
}

// NewIPBucket .
func NewIPBucket(name string, limit int) *IP {
	return &IP{
		name:  name,
		limit: limit,
		stop:  make(chan struct{}, 0),
		m:     sync.RWMutex{},
	}
}

// Start .
func (i *IP) Start(ctx context.Context, blacklist chan string, close chan string) {
CYCLE:
	for {
		select {
		case <-i.stop:
			blacklist <- i.name
			break CYCLE
		case <-ctx.Done():
			// The lifetime of the bucket has expired, leave now
			close <- i.name
			break CYCLE
		}
	}
}

// Decrement .
func (i *IP) Decrement() {
	i.m.Lock()
	defer i.m.Unlock()
	if i.limit--; i.limit < 0 {
		i.stop <- struct{}{}
	}
	return
}
