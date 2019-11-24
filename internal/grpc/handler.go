package grpc

import (
	"context"

	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	api "github.com/omerkaya1/abf-guard/internal/grpc/api"
)

// Authorisation is a handler for the GRPC Authorisation request
func (s *ABFGuardServer) Authorisation(ctx context.Context, r *api.AuthorisationRequest) (*api.Response, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	err := s.StorageService.GreenLightPass(ctx, r.GetIp())
	switch err {
	case errors.ErrDoesNotExist:
		return PrepareGRPCResponse(s.BucketService.Dispatch(r.GetLogin(), r.GetPassword(), r.GetIp())), nil
	case errors.ErrIsInTheBlacklist:
		return PrepareGRPCResponse(false, nil), nil
	case nil:
		return PrepareGRPCResponse(true, nil), nil
	default:
		return nil, err
	}
}

// FlushBuckets is a handler for the GRPC FlushBuckets request
func (s *ABFGuardServer) FlushBuckets(ctx context.Context, r *api.FlushBucketRequest) (*api.Response, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	if err := s.BucketService.FlushBuckets(r.GetLogin(), r.GetIp()); err != nil {
		return PrepareGRPCResponse(true, nil), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// PurgeBucket is a handler for the GRPC PurgeBucket request
func (s *ABFGuardServer) PurgeBucket(ctx context.Context, r *api.PurgeBucketRequest) (*api.Response, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	if err := s.BucketService.PurgeBucket(r.GetName()); err != nil {
		return PrepareGRPCResponse(true, err), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// AddIPToWhitelist is a handler for the GRPC AddIPToWhitelist request
func (s *ABFGuardServer) AddIPToWhitelist(ctx context.Context, r *api.SubnetRequest) (*api.Response, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	if err := s.StorageService.AddIP(ctx, r.GetIp(), r.GetList()); err != nil {
		return PrepareGRPCResponse(false, nil), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// DeleteIPFromWhitelist is a handler for the GRPC DeleteIPFromWhitelist request
func (s *ABFGuardServer) DeleteIPFromWhitelist(ctx context.Context, r *api.SubnetRequest) (*api.Response, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	if err := s.StorageService.DeleteIP(ctx, r.GetIp(), r.GetList()); err != nil {
		return PrepareGRPCResponse(false, nil), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// AddIPToBlacklist is a handler for the GRPC AddIPToBlacklist request
func (s *ABFGuardServer) AddIPToBlacklist(ctx context.Context, r *api.SubnetRequest) (*api.Response, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	if err := s.StorageService.AddIP(ctx, r.GetIp(), r.GetList()); err != nil {
		return PrepareGRPCResponse(false, nil), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// DeleteIPFromBlacklist is a handler for the GRPC DeleteIPFromBlacklist request
func (s *ABFGuardServer) DeleteIPFromBlacklist(ctx context.Context, r *api.SubnetRequest) (*api.Response, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	if err := s.StorageService.DeleteIP(ctx, r.GetIp(), r.GetList()); err != nil {
		return PrepareGRPCResponse(false, nil), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// GetIPList is a handler for the GRPC GetIPList request
func (s *ABFGuardServer) GetIPList(ctx context.Context, r *api.ListRequest) (*api.ListResponse, error) {
	if r == nil {
		return nil, errors.ErrBadRequest
	}
	return PrepareGRPCListIPResponse(s.StorageService.GetIPList(ctx, r.GetListType()))
}
