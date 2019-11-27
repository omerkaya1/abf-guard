package grpc

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/bucket"
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/db"
	"github.com/omerkaya1/abf-guard/internal/domain/services"
	api "github.com/omerkaya1/abf-guard/internal/grpc/api"
	"github.com/stretchr/testify/assert"
)

func TestABFGuardServer_Authorisation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sp := db.NewMockStorageProcessor(ctrl)
	bm := bucket.NewMockManager(ctrl)

	c1 := sp.EXPECT().GreenLightPass(context.Background(), "1.1.1.0").Times(1).Return(errors.ErrDoesNotExist)
	c2 := bm.EXPECT().Dispatch("n", "1", "1.1.1.0").After(c1).Times(1).Return(true, nil)
	c3 := sp.EXPECT().GreenLightPass(context.Background(), "1.1.1.1").After(c2).Times(1).Return(errors.ErrIsInTheBlacklist)
	sp.EXPECT().GreenLightPass(context.Background(), "1.1.1.1").After(c3).Times(1).Return(nil)
	sp.EXPECT().GreenLightPass(context.Background(), "1.1.1.1").After(c3).Times(1).Return(assert.AnError)

	s := ABFGuardServer{
		Cfg:            nil,
		Logger:         nil,
		StorageService: &services.Storage{Processor: sp},
		BucketService:  &services.Bucket{Manager: bm},
	}
	t.Run("Empty request", func(t *testing.T) {
		if resp, err := s.Authorisation(context.Background(), nil); assert.Error(t, err) {
			assert.Equal(t, false, resp.GetOk())
		}
	})
	t.Run("Successful request", func(t *testing.T) {
		resp, err := s.Authorisation(context.Background(), PrepareGRPCAuthorisationBody("n", "1", "1.1.1.0"))
		if assert.NoError(t, err) {
			assert.Equal(t, true, resp.GetOk())
		}
	})
	t.Run("Blacklist", func(t *testing.T) {
		resp, err := s.Authorisation(context.Background(), PrepareGRPCAuthorisationBody("n", "1", "1.1.1.1"))
		if assert.NoError(t, err) {
			assert.Equal(t, false, resp.GetOk())
		}
	})
	t.Run("Whitelist", func(t *testing.T) {
		resp, err := s.Authorisation(context.Background(), PrepareGRPCAuthorisationBody("n", "1", "1.1.1.1"))
		if assert.NoError(t, err) {
			assert.Equal(t, true, resp.GetOk())
		}
	})
	t.Run("Unknown error", func(t *testing.T) {
		r, err := s.Authorisation(context.Background(), PrepareGRPCAuthorisationBody("n", "1", "1.1.1.1"))
		if assert.Error(t, err) {
			assert.Equal(t, false, r.GetOk())
		}
	})
}

func TestABFGuardServer_AddIPToBlacklist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sp := db.NewMockStorageProcessor(ctrl)

	sp.EXPECT().Add(context.Background(), "1.1.1.0", true).Times(1).Return(assert.AnError)
	sp.EXPECT().Add(context.Background(), "1.1.1.0", true).Times(1).Return(nil)

	s := ABFGuardServer{
		Cfg:            nil,
		Logger:         nil,
		StorageService: &services.Storage{Processor: sp},
		BucketService:  nil,
	}
	t.Run("Empty request", func(t *testing.T) {
		if resp, err := s.AddIPToBlacklist(context.Background(), nil); assert.Error(t, err) {
			assert.Equal(t, false, resp.GetOk())
		}
	})
	t.Run("Request failed", func(t *testing.T) {
		resp, err := s.AddIPToBlacklist(context.Background(), PrepareSubnetGrpcRequest("1.1.1.0", true))
		if assert.NoError(t, err) {
			assert.Equal(t, false, resp.GetOk())
		}
	})
	t.Run("Request succeeded", func(t *testing.T) {
		resp, err := s.AddIPToBlacklist(context.Background(), PrepareSubnetGrpcRequest("1.1.1.0", true))
		if assert.NoError(t, err) {
			assert.Equal(t, true, resp.GetOk())
		}
	})
}

func TestABFGuardServer_AddIPToWhitelist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sp := db.NewMockStorageProcessor(ctrl)

	sp.EXPECT().Add(context.Background(), "1.1.1.0", false).Times(1).Return(assert.AnError)
	sp.EXPECT().Add(context.Background(), "1.1.1.0", false).Times(1).Return(nil)

	s := ABFGuardServer{
		Cfg:            nil,
		Logger:         nil,
		StorageService: &services.Storage{Processor: sp},
		BucketService:  nil,
	}
	t.Run("Empty request", func(t *testing.T) {
		if resp, err := s.AddIPToWhitelist(context.Background(), nil); assert.Error(t, err) {
			assert.Equal(t, false, resp.GetOk())
		}
	})
	t.Run("Request failed", func(t *testing.T) {
		resp, err := s.AddIPToWhitelist(context.Background(), PrepareSubnetGrpcRequest("1.1.1.0", false))
		if assert.NoError(t, err) {
			assert.Equal(t, false, resp.GetOk())
		}
	})
	t.Run("Request succeeded", func(t *testing.T) {
		resp, err := s.AddIPToWhitelist(context.Background(), PrepareSubnetGrpcRequest("1.1.1.0", false))
		if assert.NoError(t, err) {
			assert.Equal(t, true, resp.GetOk())
		}
	})
}

func TestABFGuardServer_DeleteIPFromBlacklist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sp := db.NewMockStorageProcessor(ctrl)

	sp.EXPECT().Delete(context.Background(), "1.1.1.0", true).Times(1).Return(assert.AnError)
	sp.EXPECT().Delete(context.Background(), "1.1.1.0", true).Times(1).Return(nil)

	s := ABFGuardServer{
		Cfg:            nil,
		Logger:         nil,
		StorageService: &services.Storage{Processor: sp},
		BucketService:  nil,
	}
	t.Run("Empty request", func(t *testing.T) {
		if resp, err := s.DeleteIPFromBlacklist(context.Background(), nil); assert.Error(t, err) {
			assert.Equal(t, false, resp.GetOk())
		}
	})
	t.Run("Request failed", func(t *testing.T) {
		resp, err := s.DeleteIPFromBlacklist(context.Background(), PrepareSubnetGrpcRequest("1.1.1.0", true))
		if assert.NoError(t, err) {
			assert.Equal(t, false, resp.GetOk())
		}
	})
	t.Run("Request succeeded", func(t *testing.T) {
		resp, err := s.DeleteIPFromBlacklist(context.Background(), PrepareSubnetGrpcRequest("1.1.1.0", true))
		if assert.NoError(t, err) {
			assert.Equal(t, true, resp.GetOk())
		}
	})
}

func TestABFGuardServer_DeleteIPFromWhitelist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sp := db.NewMockStorageProcessor(ctrl)

	sp.EXPECT().Delete(context.Background(), "1.1.1.0", false).Times(1).Return(assert.AnError)
	sp.EXPECT().Delete(context.Background(), "1.1.1.0", false).Times(1).Return(nil)

	s := ABFGuardServer{
		Cfg:            nil,
		Logger:         nil,
		StorageService: &services.Storage{Processor: sp},
		BucketService:  nil,
	}
	t.Run("Empty request", func(t *testing.T) {
		if resp, err := s.DeleteIPFromWhitelist(context.Background(), nil); assert.Error(t, err) {
			assert.Equal(t, false, resp.GetOk())
		}
	})
	t.Run("Request failed", func(t *testing.T) {
		resp, err := s.DeleteIPFromWhitelist(context.Background(), PrepareSubnetGrpcRequest("1.1.1.0", false))
		if assert.NoError(t, err) {
			assert.Equal(t, false, resp.GetOk())
		}
	})
	t.Run("Request succeeded", func(t *testing.T) {
		resp, err := s.DeleteIPFromWhitelist(context.Background(), PrepareSubnetGrpcRequest("1.1.1.0", false))
		if assert.NoError(t, err) {
			assert.Equal(t, true, resp.GetOk())
		}
	})
}

func TestABFGuardServer_FlushBuckets(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bm := bucket.NewMockManager(ctrl)

	bm.EXPECT().FlushBuckets("n", "1.1.1.0").Times(1).Return(nil)
	bm.EXPECT().FlushBuckets("", "").Times(1).Return(assert.AnError)

	s := ABFGuardServer{
		Cfg:            nil,
		Logger:         nil,
		StorageService: nil,
		BucketService:  &services.Bucket{Manager: bm},
	}
	t.Run("Empty request", func(t *testing.T) {
		if resp, err := s.FlushBuckets(context.Background(), nil); assert.Error(t, err) {
			assert.Equal(t, false, resp.GetOk())
		}
	})
	t.Run("Successful request", func(t *testing.T) {
		resp, err := s.FlushBuckets(context.Background(), PrepareFlushBucketsGrpcRequest("n", "1.1.1.0"))
		if assert.NoError(t, err) {
			assert.Equal(t, true, resp.GetOk())
		}
	})
	t.Run("Unsuccessful request", func(t *testing.T) {
		resp, err := s.FlushBuckets(context.Background(), &api.FlushBucketRequest{Login: "", Ip: ""})
		if assert.NoError(t, err) {
			assert.Equal(t, false, resp.GetOk())
		}
	})
}

func TestABFGuardServer_GetIPList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sp := db.NewMockStorageProcessor(ctrl)

	wl := []string{"1.1.1.0", "1.1.1.1"}
	bl := []string{"1.1.1.2", "1.1.1.3"}

	sp.EXPECT().GetIPList(context.Background(), true).Times(1).Return(bl, nil)
	sp.EXPECT().GetIPList(context.Background(), false).Times(1).Return(wl, nil)

	s := ABFGuardServer{
		Cfg:            nil,
		Logger:         nil,
		StorageService: &services.Storage{Processor: sp},
		BucketService:  nil,
	}
	t.Run("Empty request", func(t *testing.T) {
		if resp, err := s.GetIPList(context.Background(), nil); assert.Error(t, err) {
			assert.Nil(t, nil, resp.GetIps().GetList())
		}
	})
	t.Run("Blacklist", func(t *testing.T) {
		resp, err := s.GetIPList(context.Background(), PrepareIPListGrpcRequest(true))
		if assert.NoError(t, err) {
			assert.Equal(t, len(bl), len(resp.GetIps().List))
			assert.Equal(t, bl, resp.GetIps().List)
		}
	})
	t.Run("Whitelist", func(t *testing.T) {
		resp, err := s.GetIPList(context.Background(), PrepareIPListGrpcRequest(false))
		if assert.NoError(t, err) {
			assert.Equal(t, len(wl), len(resp.GetIps().List))
			assert.Equal(t, wl, resp.GetIps().List)
		}
	})
}

func TestABFGuardServer_PurgeBucket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bm := bucket.NewMockManager(ctrl)

	bm.EXPECT().PurgeBucket("n").Times(1).Return(nil)
	bm.EXPECT().PurgeBucket("n").Times(1).Return(assert.AnError)

	s := ABFGuardServer{
		Cfg:            nil,
		Logger:         nil,
		StorageService: nil,
		BucketService:  &services.Bucket{Manager: bm},
	}
	t.Run("Empty request", func(t *testing.T) {
		if resp, err := s.PurgeBucket(context.Background(), nil); assert.Error(t, err) {
			assert.Equal(t, false, resp.GetOk())
		}
	})
	t.Run("Successful request", func(t *testing.T) {
		resp, err := s.PurgeBucket(context.Background(), PreparePurgeBucketGrpcRequest("n"))
		if assert.NoError(t, err) {
			assert.Equal(t, true, resp.GetOk())
		}
	})
	t.Run("Unsuccessful request", func(t *testing.T) {
		resp, err := s.PurgeBucket(context.Background(), PreparePurgeBucketGrpcRequest("n"))
		if assert.NoError(t, err) {
			assert.Equal(t, false, resp.GetOk())
		}
	})
}