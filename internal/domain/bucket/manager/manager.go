package manager

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/omerkaya1/abf-guard/internal/domain/bucket/entity"
	"github.com/omerkaya1/abf-guard/internal/domain/bucket/settings"
	"github.com/omerkaya1/abf-guard/internal/domain/bucket/store"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/bucket"
)

// Manager is an object that controls all the functionality to manage buckets
type Manager struct {
	settings *settings.Settings
	store    bucket.Storage
	emptied  chan string
	errChan  chan error
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewManager creates a new Manager object and returns it to the callee
func NewManager(settings *settings.Settings) (*Manager, error) {
	if settings == nil {
		return nil, errors.ErrNilSettings
	}
	ctx, cancel := context.WithCancel(context.Background())
	mgr := &Manager{
		settings: settings,
		store:    store.NewActiveBucketsStore(),
		emptied:  make(chan string, 3),
		errChan:  make(chan error, 10),
		ctx:      ctx,
		cancel:   cancel,
	}
	go mgr.monitor()
	return mgr, nil
}

// Dispatch accepts authorisation request parameters and creates a new or decrements a counter for each bucket
func (m *Manager) Dispatch(login, pwd, ip string) (bool, error) {
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
func (m *Manager) FlushBuckets(login, ip string) error {
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
func (m *Manager) PurgeBucket(name string) error {
	if name == "" {
		return errors.ErrEmptyBucketName
	}
	if !m.store.CheckBucket(name) {
		return errors.ErrNoBucketFound
	}
	return m.store.RemoveBucket(name)
}

// GetErrChan returns an error channel to monitor the Manager's activity
func (m *Manager) GetErrChan() chan error {
	return m.errChan
}

func (m *Manager) concurrentDispatch(wg *sync.WaitGroup, name string, bucketType int, result chan bool) {
	if b, _ := m.store.GetBucket(name); b != nil {
		result <- b.Decrement()
	} else {
		switch bucketType {
		case 0:
			m.store.AddBucket(name, entity.NewBucket(m.ctx, name, m.settings.LoginLimit, m.settings.Expire, m.emptied))
			break
		case 1:
			m.store.AddBucket(name, entity.NewBucket(m.ctx, name, m.settings.PasswordLimit, m.settings.Expire, m.emptied))
			break
		default:
			m.store.AddBucket(name, entity.NewBucket(m.ctx, name, m.settings.IPLimit, m.settings.Expire, m.emptied))
			break
		}
		result <- true
	}
	wg.Done()
	return
}

func (m *Manager) monitor() {
	// Handle interrupt
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt)
	for {
		select {
		// This case handles buckets that reported their removal
		case name := <-m.emptied:
			m.errChan <- m.store.RemoveBucket(name)
		case <-exitChan:
			m.cancel()
			close(m.errChan)
			return
		}
	}
}
