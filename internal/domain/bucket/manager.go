package bucket

import (
	"context"
	"github.com/omerkaya1/abf-guard/internal/domain/bucket/entity"
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces"
	"sync"
)

// Manager .
// One ring to rule em all!
type Manager struct {
	settings      *Settings
	mutex         sync.RWMutex
	activeBuckets map[string]interfaces.Bucket
	errChan       chan error
	blacklist     chan string
	emptied       chan string
	toBlackList   chan string
}

// NewManager .
func NewManager(settings *Settings, errChan chan error, blackChan chan string) *Manager {
	return &Manager{
		// NOTE: we need settings to configure buckets
		settings: settings,
		// NOTE: we need a mutex to guard the bucket map
		mutex: sync.RWMutex{},
		// NOTE: we need a map of started buckets so that we can remove them
		activeBuckets: make(map[string]interfaces.Bucket),
		// NOTE: we need an err chan so that we can log errors
		errChan: errChan,
		// NOTE: we need this channel so that we can place an intruder to the blacklist
		blacklist:   blackChan,
		emptied:     make(chan string, 0),
		toBlackList: make(chan string, 0),
	}
}

// Monitor handles all interactions with buckets
func (m *Manager) Monitor() {
CYCLE:
	for {
		select {
		case violator := <-m.toBlackList:
			m.blacklist <- violator
		case t := <-m.emptied:
			m.removeBucket(t)
		default:
			break CYCLE
		}
	}
}

// Dispatch .
func (m *Manager) Dispatch(login, password, ip string) {
	if bucket := m.getBucket(login); bucket != nil {
		bucket.Decrement()
	} else {
		m.prepareNewLoginBucket(login)
	}
	if bucket := m.getBucket(password); bucket != nil {
		bucket.Decrement()
	} else {
		m.prepareNewPasswordBucket(password)
	}
	if bucket := m.getBucket(ip); bucket != nil {
		bucket.Decrement()
	} else {
		m.prepareNewIPBucket(ip)
	}
}

// FlushBucket .
func (m *Manager) FlushBucket(login, ip string) {
	if bucket := m.getBucket(login); bucket != nil {
		m.removeBucket(login)
	}
	if bucket := m.getBucket(ip); bucket != nil {
		m.removeBucket(ip)
	}
}

func (m *Manager) getBucket(name string) interfaces.Bucket {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if result, ok := m.activeBuckets[name]; ok {
		return result
	}
	return nil
}

func (m *Manager) addBucket(name string, b interfaces.Bucket) {
	m.mutex.Lock()
	m.activeBuckets[name] = b
	m.mutex.Unlock()
}

func (m *Manager) removeBucket(name string) {
	m.mutex.Lock()
	delete(m.activeBuckets, name)
	m.mutex.Unlock()
}

func (m *Manager) prepareNewIPBucket(ip string) {
	b := entity.NewIPBucket(ip, m.settings.IPLimit)
	m.addBucket(ip, b)
	ctx, _ := context.WithTimeout(context.Background(), m.settings.Expire)
	go b.Start(ctx, m.blacklist, m.emptied)
}

func (m *Manager) prepareNewLoginBucket(login string) {
	b := entity.NewLoginBucket(login, m.settings.LoginLimit)
	m.addBucket(login, b)
	ctx, _ := context.WithTimeout(context.Background(), m.settings.Expire)
	go b.Start(ctx, m.blacklist, m.emptied)
}

func (m *Manager) prepareNewPasswordBucket(pwd string) {
	b := entity.NewPasswordBucket(pwd, m.settings.PasswordLimit)
	m.addBucket(pwd, b)
	ctx, _ := context.WithTimeout(context.Background(), m.settings.Expire)
	go b.Start(ctx, m.blacklist, m.emptied)
}
