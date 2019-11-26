package manager

import (
	"github.com/omerkaya1/abf-guard/internal/domain/bucket/settings"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewManager(t *testing.T) {
	testCases := []struct {
		header   string
		err      error
		settings *settings.Settings
	}{
		{"Nil settings passed", errors.ErrNilSettings, nil},
		{"Correct settings passed", nil, &settings.Settings{
			LoginLimit:    2,
			PasswordLimit: 5,
			IPLimit:       10,
			Expire:        2 * time.Second,
		}},
	}

	t.Run(testCases[0].header, func(t *testing.T) {
		if bm, err := NewManager(testCases[0].settings); assert.Equal(t, testCases[0].err, err) {
			assert.Nil(t, bm)
		}
	})
	t.Run(testCases[1].header, func(t *testing.T) {
		if bm, err := NewManager(testCases[1].settings); assert.Equal(t, testCases[1].err, err) {
			assert.NotNil(t, bm)
		}
	})
}

func TestManager_Dispatch(t *testing.T) {
	testCases := []struct {
		header   string
		response bool
		err      error
		login    string
		pwd      string
		ip       string
	}{
		{"First request", true, nil, "morty", "123", "10.0.0.1"},
		{"Second request", true, nil, "morty", "123", "10.0.0.1"},
		{"Third request", false, errors.ErrBucketFull, "morty", "123", "10.0.0.1"},
	}

	pr, err := NewManager(&settings.Settings{
		LoginLimit:    2,
		PasswordLimit: 5,
		IPLimit:       10,
		Expire:        2 * time.Second,
	})
	assert.NoError(t, err)

	for i := 0; i < 3; i++ {
		t.Run("Zero request", func(t *testing.T) {
			switch i {
			case 0:
				if s, err := pr.Dispatch("", "qwre", "1.1.1.1"); assert.Error(t, err) {
					assert.Equal(t, false, s)
				}
			case 1:
				if s, err := pr.Dispatch("ad", "", "1.1.1.1"); assert.Error(t, err) {
					assert.Equal(t, false, s)
				}
			default:
				if s, err := pr.Dispatch("qewr", "qwr", ""); assert.Error(t, err) {
					assert.Equal(t, false, s)
				}
			}
		})
	}

	for i := 0; i < len(testCases); i++ {
		t.Run(testCases[i].header, func(t *testing.T) {
			if s, err := pr.Dispatch(testCases[i].login, testCases[i].pwd, testCases[i].ip); assert.Equal(t, testCases[i].err, err) {
				assert.Equal(t, testCases[i].response, s)
			}
		})
	}
}

func TestManager_FlushBuckets(t *testing.T) {
	pr, err := NewManager(&settings.Settings{LoginLimit: 3, PasswordLimit: 5, IPLimit: 10, Expire: 2 * time.Second})
	assert.NoError(t, err)

	requests := []struct {
		login string
		pwd   string
		ip    string
	}{
		{"morty", "123", "10.0.0.1"},
		{"morty", "123", "10.0.0.1"},
		{"morty", "123", "10.0.0.1"},
	}

	for _, r := range requests {
		if s, err := pr.Dispatch(r.login, r.pwd, r.ip); assert.NoError(t, err) {
			assert.Equal(t, true, s)
		}
	}

	req := struct {
		header string
		errStr string
		login  string
		ip     string
	}{"Third request", "no bucket found in store for provided login: morty", "morty", "10.0.0.1"}

	t.Run("Zero request", func(t *testing.T) {
		assert.Error(t, pr.FlushBuckets("", "1.1.1.1"))
	})
	t.Run("First request", func(t *testing.T) {
		assert.Error(t, pr.FlushBuckets("test", ""))
	})

	t.Run("Second request", func(t *testing.T) {
		assert.NoError(t, pr.FlushBuckets("morty", "10.0.0.1"))
	})

	t.Run(req.header, func(t *testing.T) {
		if err := pr.FlushBuckets(req.login, req.ip); assert.Error(t, err) {
			assert.Equal(t, req.errStr, err.Error())
		}
	})
}

func TestManager_PurgeBucket(t *testing.T) {
	pr, err := NewManager(&settings.Settings{LoginLimit: 3, PasswordLimit: 5, IPLimit: 10, Expire: 2 * time.Second})
	assert.NoError(t, err)

	successRequests := []struct {
		header string
		login  string
		pwd    string
		ip     string
	}{
		{"First purge request", "morty", "123", "10.0.0.1"},
		{"Second purge request", "morty", "123", "10.0.0.1"},
		{"Third purge request", "morty", "123", "10.0.0.1"},
	}

	for _, r := range successRequests {
		if s, err := pr.Dispatch(r.login, r.pwd, r.ip); assert.NoError(t, err) {
			assert.Equal(t, true, s)
		}
	}

	failureRequests := []struct {
		header string
		errStr string
		bucket string
	}{
		{"Fourth purge request", errors.ErrNoBucketFound.Error(), "morty"},
		{"Sixth purge request", errors.ErrNoBucketFound.Error(), "123"},
		{"Seventh purge request", errors.ErrNoBucketFound.Error(), "10.0.0.1"},
	}

	t.Run("Zero purge request", func(t *testing.T) {
		assert.Error(t, pr.PurgeBucket(""))
	})

	for i, r := range successRequests {
		t.Run(r.header, func(t *testing.T) {
			switch i {
			case 0:
				assert.NoError(t, pr.PurgeBucket(r.login))
			case 1:
				assert.NoError(t, pr.PurgeBucket(r.pwd))
			case 2:
				assert.NoError(t, pr.PurgeBucket(r.ip))
			}
		})
	}
	for _, r := range failureRequests {
		t.Run(r.header, func(t *testing.T) {
			assert.Equal(t, r.errStr, pr.PurgeBucket(r.bucket).Error())
		})
	}
}

func TestManager_GetErrChan(t *testing.T) {
	pr, err := NewManager(&settings.Settings{LoginLimit: 3, PasswordLimit: 5, IPLimit: 10, Expire: 2 * time.Second})
	assert.NoError(t, err)

	go func() {
		var err error
		pr.errChan <- err
		close(pr.errChan)
	}()

	for v := range pr.GetErrChan() {
		assert.NoError(t, v)
	}
}
