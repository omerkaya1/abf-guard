package manager

import (
	"fmt"
	"github.com/omerkaya1/abf-guard/internal/domain/bucket/entity"
	"github.com/omerkaya1/abf-guard/internal/domain/bucket/settings"
	"github.com/omerkaya1/abf-guard/internal/domain/bucket/store"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/bucket"
	"log"
	"os"
	"os/signal"
	"sync"
)

// Manager .
// One ring to rule em all!
type Manager struct {
	settings *settings.Settings
	store    bucket.Storage
	emptied  chan string
	errChan  chan error
}

// NewManager .
func NewManager(settings *settings.Settings) *Manager {
	mgr := &Manager{
		settings: settings,
		store:    store.NewActiveBucketsStore(),
		emptied:  make(chan string, 3),
		errChan:  make(chan error, 10),
	}
	go mgr.monitor()
	return mgr
}

// Dispatch accepts authorisation request parameters and creates a new or decrements  bucket
func (m *Manager) Dispatch(login, pwd, ip string) (bool, error) {
	if err := ValidateAuthorisationParams(login, pwd, ip); err != nil {
		return false, err
	}
	resultChan := make(chan bool, 3)
	wg := &sync.WaitGroup{}
	// Concurrently dispatch buckets
	for bucketName, bucketType := range PrepareAuthorisationMap(login, pwd, ip) {
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
			log.Println("not ok")
			return v, errors.ErrBucketFull
		}
	}
	// Everything is ok
	log.Println("ok")
	return true, nil
}

// FlushBucket .
func (m *Manager) FlushBuckets(login, ip string) error {
	if err := ValidateFlashParams(login, ip); err != nil {
		return err
	}
	if m.store.CheckBucket(login) {
		m.errChan <- m.store.RemoveBucket(login)
	} else {
		return fmt.Errorf("no bucket for login: %s", login)
	}
	if m.store.CheckBucket(ip) {
		m.errChan <- m.store.RemoveBucket(ip)
	} else {
		return fmt.Errorf("no bucket for ip: %s", ip)
	}
	return nil
}

// PurgeBucket .
func (m *Manager) PurgeBucket(name string) error {
	if name == "" {
		return errors.ErrEmptyBucketName
	}
	if !m.store.CheckBucket(name) {
		return errors.ErrNoBucketFound
	}
	return m.store.RemoveBucket(name)
}

// GetErrChan .
func (m *Manager) GetErrChan() chan error {
	return m.errChan
}

func (m *Manager) concurrentDispatch(wg *sync.WaitGroup, name string, bucketType int, result chan bool) {
	if b, _ := m.store.GetBucket(name); b != nil {
		result <- b.Decrement()
	} else {
		switch bucketType {
		case 0:
			m.store.AddBucket(name, entity.NewBucket(name, m.settings.LoginLimit, m.settings.Expire, m.emptied))
			break
		case 1:
			m.store.AddBucket(name, entity.NewBucket(name, m.settings.PasswordLimit, m.settings.Expire, m.emptied))
			break
		default:
			m.store.AddBucket(name, entity.NewBucket(name, m.settings.IPLimit, m.settings.Expire, m.emptied))
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
			return
		}
	}
}
