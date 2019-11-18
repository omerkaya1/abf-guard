package grpc

import (
	"context"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	abfg "github.com/omerkaya1/abf-guard/internal/grpc/api"
)

func (s *ABFGServer) Authorisation(ctx context.Context, r *abfg.AuthorisationRequest) (*abfg.Response, error) {
	if r == nil {
		return PrepareGRPCResponse(false, errors.ErrBadRequest), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

func (s *ABFGServer) FlashBucket(ctx context.Context, r *abfg.FlushBucketRequest) (*abfg.Response, error) {
	if r == nil {
		return PrepareGRPCResponse(false, errors.ErrBadRequest), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

func (s *ABFGServer) AddIpToWhitelist(ctx context.Context, r *abfg.SubnetRequest) (*abfg.Response, error) {
	if r == nil {
		return PrepareGRPCResponse(false, errors.ErrBadRequest), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

func (s *ABFGServer) DeleteIpFromWhitelist(ctx context.Context, r *abfg.SubnetRequest) (*abfg.Response, error) {
	if r == nil {
		return PrepareGRPCResponse(false, errors.ErrBadRequest), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

func (s *ABFGServer) AddIpToBlacklist(ctx context.Context, r *abfg.SubnetRequest) (*abfg.Response, error) {
	if r == nil {
		return PrepareGRPCResponse(false, errors.ErrBadRequest), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

func (s *ABFGServer) DeleteIpFromBlacklist(ctx context.Context, r *abfg.SubnetRequest) (*abfg.Response, error) {
	if r == nil {
		return PrepareGRPCResponse(false, errors.ErrBadRequest), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

func (s *ABFGServer) GetIpList(ctx context.Context, r *abfg.ListRequest) (*abfg.ListResponse, error) {
	if r == nil {
		return PrepareGRPCListIpResponse(nil, errors.ErrBadRequest)
	}
	return PrepareGRPCListIpResponse(nil, errors.ErrBadRequest)
}
