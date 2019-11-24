package store

import (
	"fmt"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/bucket"
	"sync"
)

// ActiveBucketsStore .
type ActiveBucketsStore struct {
	mutex         sync.RWMutex
	activeBuckets map[string]bucket.Bucket
}

// NewActiveBucketsStore .
func NewActiveBucketsStore() *ActiveBucketsStore {
	return &ActiveBucketsStore{
		mutex:         sync.RWMutex{},
		activeBuckets: make(map[string]bucket.Bucket),
	}
}

// GetBucket .
func (abs *ActiveBucketsStore) GetBucket(name string) (bucket.Bucket, error) {
	if b := abs.checkPresence(name); b != nil {
		return b, nil
	}
	return nil, fmt.Errorf("%s: %s", errors.ErrBucketStoragePrefix, errors.ErrNoBucketFound)
}

// AddBucket .
func (abs *ActiveBucketsStore) AddBucket(name string, b bucket.Bucket) {
	abs.mutex.Lock()
	abs.activeBuckets[name] = b
	abs.mutex.Unlock()
}

// CheckBucket .
func (abs *ActiveBucketsStore) CheckBucket(name string) bool {
	if b := abs.checkPresence(name); b != nil {
		return true
	}
	return false
}

// RemoveBucket .
func (abs *ActiveBucketsStore) RemoveBucket(name string) error {
	b := abs.checkPresence(name)
	if b == nil {
		return fmt.Errorf("%s: %s", errors.ErrBucketStoragePrefix, errors.ErrDeleteMissingBucket)
	}
	// Just ot be extra sure, we release all the resources
	b.Stop()
	abs.mutex.Lock()
	delete(abs.activeBuckets, name)
	abs.mutex.Unlock()
	return nil
}

func (abs *ActiveBucketsStore) checkPresence(name string) bucket.Bucket {
	abs.mutex.RLock()
	defer abs.mutex.RUnlock()
	return abs.activeBuckets[name]
}
