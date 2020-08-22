package bucket

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/stretchr/testify/require"
)

func TestNewManager(t *testing.T) {
	testCases := []struct {
		header   string
		err      error
		ctx      context.Context
		settings *Settings
	}{
		{"Correct settings seconds", nil, context.Background(), &Settings{
			LoginLimit:    2,
			PasswordLimit: 5,
			IPLimit:       10,
			Expire:        2 * time.Second,
		}},
		{"Correct settings minutes", nil, context.Background(), &Settings{
			LoginLimit:    2,
			PasswordLimit: 5,
			IPLimit:       10,
			Expire:        2 * time.Minute,
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.header, func(t *testing.T) {
			bm, err := NewManager(tc.ctx, tc.settings)
			require.Equal(t, tc.err, err)
			require.NotNil(t, bm)
		})
	}
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

	pr, err := NewManager(context.Background(), &Settings{
		LoginLimit:    2,
		PasswordLimit: 5,
		IPLimit:       10,
		Expire:        2 * time.Second,
	})
	require.NoError(t, err)

	for i := 0; i < 3; i++ {
		t.Run("Zero request", func(t *testing.T) {
			switch i {
			case 0:
				s, err := pr.Dispatch("", "qwre", "1.1.1.1")
				require.Error(t, err)
				require.Equal(t, false, s)
			case 1:
				s, err := pr.Dispatch("ad", "", "1.1.1.1")
				require.Error(t, err)
				require.Equal(t, false, s)
			default:
				s, err := pr.Dispatch("qewr", "qwr", "")
				require.Error(t, err)
				require.Equal(t, false, s)
			}
		})
	}

	for i := 0; i < len(testCases); i++ {
		t.Run(testCases[i].header, func(t *testing.T) {
			s, err := pr.Dispatch(testCases[i].login, testCases[i].pwd, testCases[i].ip)
			require.Equal(t, testCases[i].err, err)
			require.Equal(t, testCases[i].response, s)
		})
	}
}

func TestManager_FlushBuckets(t *testing.T) {
	pr, err := NewManager(context.Background(),
		&Settings{LoginLimit: 3, PasswordLimit: 5, IPLimit: 10, Expire: 2 * time.Second})
	require.NoError(t, err)

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
		s, err := pr.Dispatch(r.login, r.pwd, r.ip)
		require.NoError(t, err)
		require.Equal(t, true, s)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	time.AfterFunc(time.Second*2, func() {
		defer wg.Done()
		s, err := pr.Dispatch(requests[0].login, requests[0].pwd, requests[0].ip)
		require.NoError(t, err)
		require.Equal(t, true, s)
	})

	req := struct {
		header string
		errStr string
		login  string
		ip     string
	}{"Third request", "no bucket found in store for provided login: morty", "morty", "10.0.0.1"}

	t.Run("Zero request", func(t *testing.T) {
		require.Error(t, pr.FlushBuckets("", "1.1.1.1"))
	})
	t.Run("First request", func(t *testing.T) {
		require.Error(t, pr.FlushBuckets("test", ""))
	})

	t.Run("Second request", func(t *testing.T) {
		require.NoError(t, pr.FlushBuckets("morty", "10.0.0.1"))
	})

	t.Run(req.header, func(t *testing.T) {
		err := pr.FlushBuckets(req.login, req.ip)
		require.Error(t, err)
		require.Equal(t, req.errStr, err.Error())
	})
	wg.Wait()
	t.Run("Purge ip bucket", func(t *testing.T) {
		require.NoError(t, pr.PurgeBucket("10.0.0.1"))
	})
	t.Run("Err check", func(t *testing.T) {
		err := pr.FlushBuckets(req.login, req.ip)
		require.Error(t, err)
		require.Equal(t, "no bucket found in store for provided ip: 10.0.0.1", err.Error())
	})
}

func TestManager_PurgeBucket(t *testing.T) {
	pr, err := NewManager(context.Background(),
		&Settings{LoginLimit: 3, PasswordLimit: 5, IPLimit: 10, Expire: 2 * time.Second})
	require.NoError(t, err)

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
		s, err := pr.Dispatch(r.login, r.pwd, r.ip)
		require.NoError(t, err)
		require.Equal(t, true, s)
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
		require.Error(t, pr.PurgeBucket(""))
	})

	for i, r := range successRequests {
		t.Run(r.header, func(t *testing.T) {
			switch i {
			case 0:
				require.NoError(t, pr.PurgeBucket(r.login))
			case 1:
				require.NoError(t, pr.PurgeBucket(r.pwd))
			case 2:
				require.NoError(t, pr.PurgeBucket(r.ip))
			}
		})
	}
	for _, r := range failureRequests {
		t.Run(r.header, func(t *testing.T) {
			require.Equal(t, r.errStr, pr.PurgeBucket(r.bucket).Error())
		})
	}
}

func TestManager_PurgeBucket_Ctx(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()
	pr, err := NewManager(ctx,
		&Settings{LoginLimit: 3, PasswordLimit: 5, IPLimit: 10, Expire: 2 * time.Second})
	require.NoError(t, err)

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
		s, err := pr.Dispatch(r.login, r.pwd, r.ip)
		require.NoError(t, err)
		require.Equal(t, true, s)
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
		require.Error(t, pr.PurgeBucket(""))
	})

	for i, r := range successRequests {
		t.Run(r.header, func(t *testing.T) {
			switch i {
			case 0:
				require.NoError(t, pr.PurgeBucket(r.login))
			case 1:
				require.NoError(t, pr.PurgeBucket(r.pwd))
			case 2:
				require.NoError(t, pr.PurgeBucket(r.ip))
			}
		})
	}
	for _, r := range failureRequests {
		t.Run(r.header, func(t *testing.T) {
			require.Equal(t, r.errStr, pr.PurgeBucket(r.bucket).Error())
		})
	}
}

func TestManager_GetErrChan(t *testing.T) {
	pr, err := NewManager(
		context.Background(),
		&Settings{LoginLimit: 3, PasswordLimit: 5, IPLimit: 10, Expire: 2 * time.Second})
	require.NoError(t, err)

	go func() {
		var err error
		pr.GetErrChan() <- err
		close(pr.GetErrChan())
	}()

	for v := range pr.GetErrChan() {
		require.NoError(t, v)
	}
}
