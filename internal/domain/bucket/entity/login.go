package entity

import (
	"context"
	"sync"
)

// Login .
type Login struct {
	limit int
	name  string
	stop  chan struct{}
	m     sync.RWMutex
}

// NewLoginBucket .
func NewLoginBucket(name string, limit int) *Login {
	return &Login{
		name:  name,
		limit: limit,
		stop:  make(chan struct{}, 0),
		m:     sync.RWMutex{},
	}
}

// Start .
func (l *Login) Start(ctx context.Context, blacklist chan string, close chan string) {
CYCLE:
	for {
		select {
		case <-l.stop:
			blacklist <- l.name
			break CYCLE
		case <-ctx.Done():
			// The lifetime of the bucket has expired, leave now
			close <- l.name
			break CYCLE
		}
	}
}

// Decrement .
func (l *Login) Decrement() {
	l.m.Lock()
	defer l.m.Unlock()
	if l.limit--; l.limit < 0 {
		l.stop <- struct{}{}
	}
	return
}
