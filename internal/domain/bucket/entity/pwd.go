package entity

import (
	"context"
	"sync"
)

// Password .
type Password struct {
	limit int
	name  string
	stop  chan struct{}
	m     sync.RWMutex
}

// NewPasswordBucket .
func NewPasswordBucket(name string, limit int) *Password {
	return &Password{
		name:  name,
		limit: limit,
		stop:  make(chan struct{}, 0),
		m:     sync.RWMutex{},
	}
}

// Start .
func (p *Password) Start(ctx context.Context, blacklist chan string, close chan string) {
CYCLE:
	for {
		select {
		case <-p.stop:
			blacklist <- p.name
			break CYCLE
		case <-ctx.Done():
			// The lifetime of the bucket has expired, leave now
			close <- p.name
			break CYCLE
		}
	}
}

// Decrement .
func (p *Password) Decrement() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.limit--; p.limit < 0 {
		p.stop <- struct{}{}
	}
	return
}
