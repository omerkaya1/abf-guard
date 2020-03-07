package manager

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/omerkaya1/abf-guard/internal/domain/bucket/entity"
	"github.com/omerkaya1/abf-guard/internal/domain/bucket/settings"
	"github.com/omerkaya1/abf-guard/internal/domain/bucket/store"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/bucket"
)

// BucketManager is an object that controls all the functionality to manage buckets
type BucketManager struct {
	// TODO: same thing for the configuration struct (what we need here is the Validate method!)
	settings *settings.Settings
	store    bucket.Storage
	emptied  chan string
	errChan  chan error
}

// NewBucketManager creates a new Manager object and returns it to the callee
func NewBucketManager(ctx context.Context, settings *settings.Settings) (bucket.Manager, error) {
	if settings == nil {
		return nil, errors.ErrNilSettings
	}
	mgr := &BucketManager{
		settings: settings,
		store:    store.NewActiveBucketsStore(),
		emptied:  make(chan string, 3),
		errChan:  make(chan error, 10),
	}
	go mgr.monitor(ctx)
	return mgr, nil
}

// Dispatch accepts authorisation request parameters and creates a new or decrements a counter for each bucket
func (m *BucketManager) Dispatch(login, pwd, ip string) (bool, error) {
	if err := validateAuthorisationParams(login, pwd, ip); err != nil {
		return false, err
	}
	resultChan := make(chan bool, 3)
	wg := &sync.WaitGroup{}
	// Concurrently dispatch buckets
	for bucketName, bucketType := range prepareAuthorisationMap(login, pwd, ip) {
		wg.Add(1)
		go func(group *sync.WaitGroup, bucketN string, bucketT int) {
			m.concurrentDispatch(group, bucketN, bucketT, resultChan)
		}(wg, bucketName, bucketType)
	}
	// Wait for all the workers to complete
	wg.Wait()
	// Close the result channel
	close(resultChan)
	// Iterate over the results
	for v := range resultChan {
		// The first false result reports that the request cannot proceed
		if !v {
			return v, errors.ErrBucketFull
		}
	}
	// Everything is ok
	return true, nil
}

// FlushBuckets removes all buckets with the specified login and ip
func (m *BucketManager) FlushBuckets(login, ip string) error {
	if err := validateFlashParams(login, ip); err != nil {
		return err
	}
	if m.store.CheckBucket(login) {
		m.errChan <- m.store.RemoveBucket(login)
	} else {
		return fmt.Errorf("no bucket found in store for provided login: %s", login)
	}
	if m.store.CheckBucket(ip) {
		m.errChan <- m.store.RemoveBucket(ip)
	} else {
		return fmt.Errorf("no bucket found in store for provided ip: %s", ip)
	}
	return nil
}

// PurgeBucket removes a bucket which name was specified as an argument
func (m *BucketManager) PurgeBucket(name string) error {
	if name == "" {
		return errors.ErrEmptyBucketName
	}
	if !m.store.CheckBucket(name) {
		return errors.ErrNoBucketFound
	}
	return m.store.RemoveBucket(name)
}

// GetErrChan returns an error channel to monitor the Manager's activity
func (m *BucketManager) GetErrChan() chan error {
	return m.errChan
}

func (m *BucketManager) concurrentDispatch(wg *sync.WaitGroup, name string, bucketType int, result chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), m.settings.Expire)
	// Call cancel() to release the resources of the bucket
	time.AfterFunc(m.settings.Expire, cancel)
	if b, _ := m.store.GetBucket(name); b != nil {
		result <- b.Decrement()
	} else {
		switch bucketType {
		case 0:
			m.store.AddBucket(name, entity.NewBucket(ctx, name, m.settings.LoginLimit, m.emptied))
			break
		case 1:
			m.store.AddBucket(name, entity.NewBucket(ctx, name, m.settings.PasswordLimit, m.emptied))
			break
		default:
			m.store.AddBucket(name, entity.NewBucket(ctx, name, m.settings.IPLimit, m.emptied))
			break
		}
		result <- true
	}
	wg.Done()
	return
}

func (m *BucketManager) monitor(ctx context.Context) {
	for {
		select {
		// This case handles buckets that reported their removal
		case name := <-m.emptied:
			m.errChan <- m.store.RemoveBucket(name)
		// Handle context interrupt
		case <-ctx.Done():
			close(m.errChan)
			return
		}
	}
}
