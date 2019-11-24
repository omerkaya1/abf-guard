package grpc

import (
	"context"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	api "github.com/omerkaya1/abf-guard/internal/grpc/api"
)

// Authorisation .
func (s *ABFGuardServer) Authorisation(ctx context.Context, r *api.AuthorisationRequest) (*api.Response, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	ok := false
	var err error
	ok, err = s.StorageService.ExistInList(ctx, r.GetIp(), true)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrServiceCmdPrefix, err)
	}
	if ok {
		return PrepareGRPCResponse(false, nil), nil
	}
	ok, err = s.StorageService.ExistInList(ctx, r.GetIp(), false)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrServiceCmdPrefix, err)
	}
	if ok {
		return PrepareGRPCResponse(ok, nil), nil
	}
	return PrepareGRPCResponse(s.BucketService.Dispatch(r.GetLogin(), r.GetPassword(), r.GetIp())), nil
}

// FlashBuckets .
func (s *ABFGuardServer) FlushBuckets(ctx context.Context, r *api.FlushBucketRequest) (*api.Response, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	if err := s.BucketService.FlushBuckets(r.GetLogin(), r.GetIp()); err != nil {
		return PrepareGRPCResponse(true, nil), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// PurgeBucket .
func (s *ABFGuardServer) PurgeBucket(ctx context.Context, r *api.PurgeBucketRequest) (*api.Response, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	if err := s.BucketService.PurgeBucket(r.GetName()); err != nil {
		return PrepareGRPCResponse(true, err), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// AddIPToWhitelist .
func (s *ABFGuardServer) AddIPToWhitelist(ctx context.Context, r *api.SubnetRequest) (*api.Response, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	if err := s.StorageService.AddIP(ctx, r.GetIp(), r.GetList()); err != nil {
		return PrepareGRPCResponse(false, nil), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// DeleteIPFromWhitelist .
func (s *ABFGuardServer) DeleteIPFromWhitelist(ctx context.Context, r *api.SubnetRequest) (*api.Response, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	if err := s.StorageService.DeleteIP(ctx, r.GetIp(), r.GetList()); err != nil {
		return PrepareGRPCResponse(false, nil), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// AddIPToBlacklist .
func (s *ABFGuardServer) AddIPToBlacklist(ctx context.Context, r *api.SubnetRequest) (*api.Response, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	if err := s.StorageService.AddIP(ctx, r.GetIp(), r.GetList()); err != nil {
		return PrepareGRPCResponse(false, nil), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// DeleteIPFromBlacklist .
func (s *ABFGuardServer) DeleteIPFromBlacklist(ctx context.Context, r *api.SubnetRequest) (*api.Response, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	if err := s.StorageService.DeleteIP(ctx, r.GetIp(), r.GetList()); err != nil {
		return PrepareGRPCResponse(false, nil), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// GetIPList .
func (s *ABFGuardServer) GetIPList(ctx context.Context, r *api.ListRequest) (*api.ListResponse, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	return PrepareGRPCListIPResponse(s.StorageService.GetIPList(ctx, r.GetListType()))
}
