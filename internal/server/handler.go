package server

import (
	"context"
	"errors"

	"github.com/omerkaya1/abf-guard/internal/db"
	"github.com/omerkaya1/abf-guard/internal/server/api"
)

var errBadRequest = errors.New("bad request")

// Authorisation is a handler for the GRPC Authorisation request
func (s *ABFGuardServer) Authorisation(ctx context.Context, r *api.AuthorisationRequest) (*api.Response, error) {
	if r == nil {
		return nil, errBadRequest
	}
	err := s.Storage.GreenLightPass(ctx, r.GetIp())
	switch err {
	case db.ErrDoesNotExist:
		return PrepareGRPCResponse(s.BucketManager.Dispatch(r.GetLogin(), r.GetPassword(), r.GetIp())), nil
	case db.ErrIsInTheBlacklist:
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
		return nil, errBadRequest
	}
	if err := s.BucketManager.FlushBuckets(r.GetLogin(), r.GetIp()); err != nil {
		return PrepareGRPCResponse(false, err), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// PurgeBucket is a handler for the GRPC PurgeBucket request
func (s *ABFGuardServer) PurgeBucket(ctx context.Context, r *api.PurgeBucketRequest) (*api.Response, error) {
	if r == nil || r.GetName() == "" {
		return nil, errBadRequest
	}
	if err := s.BucketManager.PurgeBucket(r.GetName()); err != nil {
		return PrepareGRPCResponse(true, err), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// AddIPToWhitelist is a handler for the GRPC AddIPToWhitelist request
func (s *ABFGuardServer) AddIPToWhitelist(ctx context.Context, r *api.SubnetRequest) (*api.Response, error) {
	if r == nil || r.GetIp() == "" {
		return nil, errBadRequest
	}
	if err := s.Storage.Add(ctx, r.GetIp(), r.GetList()); err != nil {
		return PrepareGRPCResponse(false, nil), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// DeleteIPFromWhitelist is a handler for the GRPC DeleteIPFromWhitelist request
func (s *ABFGuardServer) DeleteIPFromWhitelist(ctx context.Context, r *api.SubnetRequest) (*api.Response, error) {
	if r == nil || r.GetIp() == "" {
		return nil, errBadRequest
	}
	if err := s.Storage.Delete(ctx, r.GetIp(), r.GetList()); err != nil {
		return PrepareGRPCResponse(false, nil), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// AddIPToBlacklist is a handler for the GRPC AddIPToBlacklist request
func (s *ABFGuardServer) AddIPToBlacklist(ctx context.Context, r *api.SubnetRequest) (*api.Response, error) {
	if r == nil || r.GetIp() == "" {
		return nil, errBadRequest
	}
	if err := s.Storage.Add(ctx, r.GetIp(), r.GetList()); err != nil {
		return PrepareGRPCResponse(false, nil), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// DeleteIPFromBlacklist is a handler for the GRPC DeleteIPFromBlacklist request
func (s *ABFGuardServer) DeleteIPFromBlacklist(ctx context.Context, r *api.SubnetRequest) (*api.Response, error) {
	if r == nil || r.GetIp() == "" {
		return nil, errBadRequest
	}
	if err := s.Storage.Delete(ctx, r.GetIp(), r.GetList()); err != nil {
		return PrepareGRPCResponse(false, nil), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// GetIPList is a handler for the GRPC GetIPList request
func (s *ABFGuardServer) GetIPList(ctx context.Context, r *api.ListRequest) (*api.ListResponse, error) {
	if r == nil {
		return nil, errBadRequest
	}
	return PrepareGRPCListIPResponse(s.Storage.GetIPList(ctx, r.GetListType()))
}
