package entity

import (
	"context"
	"log"
	"sync"
	"time"
)

// Bucket is a structure that represents a bucket object
type Bucket struct {
	count  int
	name   string
	mutex  sync.RWMutex
	ctx    context.Context
	cancel context.CancelFunc
}

// NewBucket returns an new bucket object to the callee
func NewBucket(ctxParent context.Context, name string, count int, expires time.Duration, done chan<- string) *Bucket {
	// We create a context for each bucket to handle the removal of buckets by calling the cancel() function
	innerCtx, cancel := context.WithCancel(ctxParent)
	// We make sure that the bucket gets deleted after a certain time
	go func(ctxInn context.Context, expire time.Duration) {
		clickClick := time.NewTicker(expire)
		select {
		case <-ctxInn.Done():
			if ctxInn.Err() == context.Canceled {
				return
			}
			// Log unexpected error
			log.Println(ctxInn.Err())
			return
		case <-clickClick.C:
			// Boom!
			done <- name
			return
		}
	}(innerCtx, expires)

	return &Bucket{
		name:   name,
		count:  count,
		mutex:  sync.RWMutex{},
		ctx:    innerCtx,
		cancel: cancel,
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
	b.cancel()
}

func (b *Bucket) checkAvailable() bool {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	if b.count > 1 {
		return true
	}
	return false
}
