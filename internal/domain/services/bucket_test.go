package services

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/bucket"
	"github.com/stretchr/testify/assert"
)

func TestBucket_Dispatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bm := bucket.NewMockManager(ctrl)

	tc := []struct {
		header   string
		response bool
		login    string
		pwd      string
		ip       string
	}{
		{"First dispatch request", true, "morty", "123", "10.0.0.1"},
		{"Second dispatch request", false, "morty", "123", "10.0.0.1"},
	}

	bm.EXPECT().Dispatch(tc[0].login, tc[0].pwd, tc[0].ip).Times(1).Return(true, nil)
	bm.EXPECT().Dispatch(tc[1].login, tc[1].pwd, tc[1].ip).Times(1).Return(false, errors.ErrBucketFull)

	bs := Bucket{Manager: bm}
	t.Run(tc[0].header, func(t *testing.T) {
		if ok, err := bs.Dispatch(tc[0].login, tc[0].pwd, tc[0].ip); assert.NoError(t, err) {
			assert.Equal(t, tc[0].response, ok)
		}
	})

	t.Run(tc[1].header, func(t *testing.T) {
		if ok, err := bs.Dispatch(tc[1].login, tc[1].pwd, tc[1].ip); assert.Equal(t, errors.ErrBucketFull, err) {
			assert.Equal(t, tc[1].response, ok)
		}
	})
}

func TestBucket_FlushBuckets(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bm := bucket.NewMockManager(ctrl)

	tc := []struct {
		header string
		login  string
		ip     string
	}{
		{"First flush request", "morty", "10.0.0.1"},
		{"Second flush request", "morty", "10.0.0.1"},
	}

	bm.EXPECT().FlushBuckets(tc[0].login, tc[0].ip).Times(1).Return(nil)
	bm.EXPECT().FlushBuckets(tc[1].login, tc[1].ip).Times(1).Return(
		fmt.Errorf("no bucket found in store for provided login: %s", tc[1].login))

	bs := Bucket{Manager: bm}
	t.Run(tc[0].header, func(t *testing.T) {
		assert.NoError(t, bs.FlushBuckets(tc[0].login, tc[0].ip))
	})

	t.Run(tc[1].header, func(t *testing.T) {
		assert.Error(t, bs.FlushBuckets(tc[1].login, tc[1].ip))
	})
}

func TestBucket_MonitorErrors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bm := bucket.NewMockManager(ctrl)

	r := make(chan error)
	bm.EXPECT().GetErrChan().Times(1).Return(r)
	bs := Bucket{Manager: bm}
	assert.Equal(t, r, bs.MonitorErrors())
}

func TestBucket_PurgeBucket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bm := bucket.NewMockManager(ctrl)

	tc := []string{"", "morty", "10.0.0.1"}

	bm.EXPECT().PurgeBucket(tc[0]).Times(1).Return(errors.ErrEmptyBucketName)
	bm.EXPECT().PurgeBucket(tc[1]).Times(1).Return(errors.ErrNoBucketFound)
	bm.EXPECT().PurgeBucket(tc[1]).Times(1).Return(nil)

	bs := Bucket{Manager: bm}
	t.Run("Empty name", func(t *testing.T) {
		assert.Error(t, bs.PurgeBucket(""))
	})

	t.Run("Missing bucket requested", func(t *testing.T) {
		assert.Error(t, bs.PurgeBucket(tc[1]))
	})

	t.Run("Successful request", func(t *testing.T) {
		assert.NoError(t, bs.PurgeBucket(tc[1]))
	})
}
