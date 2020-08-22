package domain

import (
	"errors"
	"fmt"
	"sync"
)

type (
	// Storer provides a functionality to communicate with any bucket store that satisfies the interface
	Storer interface {
		// GetBucket returns the requested bucket to the callee
		GetBucket(name string) (Bucketer, error)
		// CheckBucket checks whether a requested bucket is present in the active bucket store
		CheckBucket(name string) bool
		// AddBucket adds a new bucket to the active bucket store
		AddBucket(name string, b Bucketer)
		// RemoveBucket removes a specified bucket from the active bucket store
		RemoveBucket(name string) error
	}
	// ActiveBucketsStore is an object that provides a storage for all buckets.
	ActiveBucketsStore struct {
		mutex         sync.RWMutex
		activeBuckets map[string]Bucketer
	}
)

const prefix = "bucket storage error"

// ErrDeleteMissingBucket reports an attempt to delete not-existing buckets
var ErrDeleteMissingBucket = errors.New("no bucket found in store for deletion")

// NewActiveBucketsStore returns a new ActiveBucketsStore object to the callee
func NewActiveBucketsStore() *ActiveBucketsStore {
	return &ActiveBucketsStore{
		mutex:         sync.RWMutex{},
		activeBuckets: make(map[string]Bucketer),
	}
}

// GetBucket returns the requested bucket to the callee
func (abs *ActiveBucketsStore) GetBucket(name string) (Bucketer, error) {
	if b := abs.checkPresence(name); b != nil {
		return b, nil
	}
	return nil, fmt.Errorf("%s: %s", prefix, ErrNoBucketFound)
}

// AddBucket adds a new bucket to the active bucket store
func (abs *ActiveBucketsStore) AddBucket(name string, b Bucketer) {
	abs.mutex.Lock()
	abs.activeBuckets[name] = b
	abs.mutex.Unlock()
}

// CheckBucket checks whether a requested bucket is present in the active bucket store
func (abs *ActiveBucketsStore) CheckBucket(name string) bool {
	return abs.checkPresence(name) != nil
}

// RemoveBucket removes a specified bucket from the active bucket store
func (abs *ActiveBucketsStore) RemoveBucket(name string) error {
	b := abs.checkPresence(name)
	if b == nil {
		return fmt.Errorf("%s: %s", prefix, ErrDeleteMissingBucket)
	}
	// Just to be extra sure, we release all the resources
	b.Stop()
	abs.mutex.Lock()
	delete(abs.activeBuckets, name)
	abs.mutex.Unlock()
	return nil
}

func (abs *ActiveBucketsStore) checkPresence(name string) Bucketer {
	abs.mutex.RLock()
	defer abs.mutex.RUnlock()
	return abs.activeBuckets[name]
}
