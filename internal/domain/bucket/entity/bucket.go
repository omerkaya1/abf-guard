package entity

import (
	"context"
	"log"
	"sync"
)

// Bucket is a structure that represents a bucket object
type Bucket struct {
	count int
	name  string
	mutex sync.RWMutex
	stop  chan struct{}
}

// NewBucket returns an new bucket object to the callee
func NewBucket(ctx context.Context, name string, count int, done chan<- string) *Bucket {
	stopChan := make(chan struct{})
	// We make sure that the bucket gets deleted after a certain time
	go func(ctx context.Context, stopCh <-chan struct{}) {
		select {
		case <-ctx.Done():
			// Time's up, a bucket has to die
			if ctx.Err() == context.DeadlineExceeded {
				done <- name
				return
			}
			// Log other errors
			log.Println(ctx.Err())
			return
		case <-stopCh:
			done <- name
			return
		}
	}(ctx, stopChan)

	return &Bucket{
		name:  name,
		count: count,
		mutex: sync.RWMutex{},
		stop:  stopChan,
	}
}

// Decrement reduces the bucket request counter; it return true if the request can pass and false otherwise
func (b *Bucket) Decrement() bool {
	if ok := b.checkAvailable(); !ok {
		return ok
	}
	b.mutex.Lock()
	b.count--
	b.mutex.Unlock()
	return true
}

// Stop releases all the resources associated with the bucket
func (b *Bucket) Stop() {
	b.stop <- struct{}{}
}

func (b *Bucket) checkAvailable() bool {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	if b.count > 1 {
		return true
	}
	return false
}
