package store

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/bucket"
	"github.com/stretchr/testify/assert"
)

func TestNewActiveBucketsStore(t *testing.T) {
	assert.NotNil(t, NewActiveBucketsStore())
}

func TestActiveBucketsStore_AddBucket(t *testing.T) {
	bs := NewActiveBucketsStore()
	assert.NotNil(t, bs)

	cases := []struct {
		header string
		name   string
	}{
		{"First add bucket request", "morty"},
		{"Second add bucket request", "123"},
		{"Third add bucket request", "10.0.0.1"},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	b := bucket.NewMockBucket(ctrl)

	for _, c := range cases {
		bs.AddBucket(c.name, b)
		assert.NotNil(t, bs.CheckBucket(c.name))
	}
}

func TestActiveBucketsStore_RemoveBucket(t *testing.T) {
	bs := NewActiveBucketsStore()
	assert.NotNil(t, bs)

	cases := []string{"morty", "123", "10.0.0.1"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	b := bucket.NewMockBucket(ctrl)

	for _, c := range cases {
		bs.AddBucket(c, b)
		assert.NotNil(t, bs.CheckBucket(c))
	}

	b.EXPECT().Stop().AnyTimes()

	for _, c := range cases {
		t.Run("Remove succeeds", func(t *testing.T) {
			if err := bs.RemoveBucket(c); assert.NoError(t, err) {
				assert.Equal(t, false, bs.CheckBucket(c))
			}
		})
	}

	for _, c := range cases {
		t.Run("Remove fails", func(t *testing.T) {
			assert.Error(t, bs.RemoveBucket(c))
		})
	}
}

func TestActiveBucketsStore_CheckBucket(t *testing.T) {
	bs := NewActiveBucketsStore()
	assert.NotNil(t, bs)

	cases := []string{"morty", "123", "10.0.0.1"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	b := bucket.NewMockBucket(ctrl)

	for _, c := range cases {
		bs.AddBucket(c, b)
		assert.NotNil(t, bs.CheckBucket(c))
	}

	b.EXPECT().Stop().AnyTimes()

	for _, c := range cases {
		t.Run("Check succeeds", func(t *testing.T) {
			assert.Equal(t, true, bs.CheckBucket(c))
		})
	}

	for _, c := range cases {
		t.Run("Remove succeeds", func(t *testing.T) {
			if err := bs.RemoveBucket(c); assert.NoError(t, err) {
				assert.Equal(t, false, bs.CheckBucket(c))
			}
		})
	}

	for _, c := range cases {
		t.Run("Check fails", func(t *testing.T) {
			assert.Equal(t, false, bs.CheckBucket(c))
		})
	}
}

func TestActiveBucketsStore_GetBucket(t *testing.T) {
	bs := NewActiveBucketsStore()
	assert.NotNil(t, bs)

	cases := []string{"morty", "123", "10.0.0.1"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	b := bucket.NewMockBucket(ctrl)

	for _, c := range cases {
		t.Run("Get bucket fails", func(t *testing.T) {
			if b, err := bs.GetBucket(c); assert.Error(t, err) {
				assert.Nil(t, b)
			}
		})
	}

	for _, c := range cases {
		bs.AddBucket(c, b)
		assert.NotNil(t, bs.CheckBucket(c))
	}

	for _, c := range cases {
		t.Run("Get bucket succeeds", func(t *testing.T) {
			if b, err := bs.GetBucket(c); assert.NoError(t, err) {
				assert.NotNil(t, b)
			}
		})
	}
}
