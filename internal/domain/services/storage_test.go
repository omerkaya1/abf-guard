package services

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/db"
	"github.com/stretchr/testify/assert"
)

func TestStorage_AddIP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sp := db.NewMockStorageProcessor(ctrl)

	sp.EXPECT().Add(context.Background(), "", false).Times(1).Return(errors.ErrEmptyIP)
	sp.EXPECT().Add(context.Background(), "1.1.1.0", false).Times(1).Return(errors.ErrAlreadyStored)
	sp.EXPECT().Add(context.Background(), "1.1.1.1", false).Times(1).Return(assert.AnError)
	sp.EXPECT().Add(context.Background(), "1.1.1.2", false).Times(1).Return(nil)
	sp.EXPECT().Add(context.Background(), "1.1.1.3", true).Times(1).Return(nil)

	tsp := Storage{Processor: sp}
	// Empty parameter test
	t.Run("Empty parameter", func(t *testing.T) {
		assert.Error(t, tsp.AddIP(context.Background(), "", false))
	})
	// Already stored IP test
	t.Run("Already stored IP", func(t *testing.T) {
		assert.Error(t, tsp.AddIP(context.Background(), "1.1.1.0", false))
	})
	// DB error test
	t.Run("DB error", func(t *testing.T) {
		assert.Error(t, tsp.AddIP(context.Background(), "1.1.1.1", false))
	})
	// Add to whitelist succeeded
	t.Run("Add to whitelist succeeded", func(t *testing.T) {
		assert.NoError(t, tsp.AddIP(context.Background(), "1.1.1.2", false))
	})
	// Add to blacklist succeeded
	t.Run("Add to blacklist succeeded", func(t *testing.T) {
		assert.NoError(t, tsp.AddIP(context.Background(), "1.1.1.3", true))
	})
}

func TestStorage_DeleteIP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sp := db.NewMockStorageProcessor(ctrl)

	sp.EXPECT().Delete(context.Background(), "", false).Times(1).Return(errors.ErrEmptyIP)
	sp.EXPECT().Delete(context.Background(), "1.1.1.0", false).Times(1).Return(errors.ErrDoesNotExist)
	sp.EXPECT().Delete(context.Background(), "1.1.1.1", false).Times(1).Return(assert.AnError)
	sp.EXPECT().Delete(context.Background(), "1.1.1.2", false).Times(1).Return(nil)
	sp.EXPECT().Delete(context.Background(), "1.1.1.3", true).Times(1).Return(nil)

	tsp := Storage{Processor: sp}
	// Empty parameter test
	t.Run("Empty parameter", func(t *testing.T) {
		assert.Error(t, tsp.DeleteIP(context.Background(), "", false))
	})
	// Already stored IP test
	t.Run("Ip does not exist", func(t *testing.T) {
		assert.Error(t, tsp.DeleteIP(context.Background(), "1.1.1.0", false))
	})
	// DB error test
	t.Run("DB error", func(t *testing.T) {
		assert.Error(t, tsp.DeleteIP(context.Background(), "1.1.1.1", false))
	})
	// Delete from whitelist succeeded
	t.Run("Delete from whitelist succeeded", func(t *testing.T) {
		assert.NoError(t, tsp.DeleteIP(context.Background(), "1.1.1.2", false))
	})
	// Delete from blacklist succeeded
	t.Run("Delete from blacklist succeeded", func(t *testing.T) {
		assert.NoError(t, tsp.DeleteIP(context.Background(), "1.1.1.3", true))
	})
}

func TestStorage_GetIPList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sp := db.NewMockStorageProcessor(ctrl)

	wl := []string{"1.1.1.0", "1.1.1.1"}
	bl := []string{"1.1.1.2", "1.1.1.3"}

	sp.EXPECT().GetIPList(context.Background(), false).Times(1).Return(wl, nil)
	sp.EXPECT().GetIPList(context.Background(), true).Times(1).Return(bl, nil)
	sp.EXPECT().GetIPList(context.Background(), false).Times(1).Return(wl[:1], assert.AnError)
	sp.EXPECT().GetIPList(context.Background(), true).Times(1).Return(nil, assert.AnError)

	tsp := Storage{Processor: sp}
	// A list of ips from the whitelist test
	t.Run("A list of ips from the whitelist", func(t *testing.T) {
		if l, err := tsp.GetIPList(context.Background(), false); assert.NoError(t, err) {
			assert.EqualValues(t, wl, l)
		}
	})
	// A list of ips from the blacklist test
	t.Run("A list of ips from the blacklist", func(t *testing.T) {
		if l, err := tsp.GetIPList(context.Background(), true); assert.NoError(t, err) {
			assert.EqualValues(t, bl, l)
		}
	})
	// A partial return from the whitelist test
	t.Run("A partial return from the whitelist", func(t *testing.T) {
		if l, err := tsp.GetIPList(context.Background(), false); assert.Error(t, err) {
			assert.Equal(t, 1, len(l))
			assert.Contains(t, wl, l[0])
		}
	})
	// Unsuccessful result
	t.Run("Empty parameter", func(t *testing.T) {
		if l, err := tsp.GetIPList(context.Background(), true); assert.Error(t, err) {
			assert.Nil(t, l)
		}
	})
}

func TestStorage_GreenLightPass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sp := db.NewMockStorageProcessor(ctrl)

	sp.EXPECT().GreenLightPass(context.Background(), "1.1.1.0").Times(1).Return(nil)
	sp.EXPECT().GreenLightPass(context.Background(), "1.1.1.1").Times(1).Return(assert.AnError)

	tsp := Storage{Processor: sp}
	// A list of ips from the whitelist test
	t.Run("Successful call", func(t *testing.T) {
		assert.NoError(t, tsp.GreenLightPass(context.Background(), "1.1.1.0"))
	})
	t.Run("unsuccessful call", func(t *testing.T) {
		assert.Error(t, tsp.GreenLightPass(context.Background(), "1.1.1.1"))
	})
}
