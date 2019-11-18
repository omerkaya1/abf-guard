package grpc

import (
	"context"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	abfg "github.com/omerkaya1/abf-guard/internal/grpc/api"
)

// Authorisation .
func (s *ABFGServer) Authorisation(ctx context.Context, r *abfg.AuthorisationRequest) (*abfg.Response, error) {
	if r == nil {
		return PrepareGRPCResponse(false, errors.ErrBadRequest), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// FlashBucket .
func (s *ABFGServer) FlashBucket(ctx context.Context, r *abfg.FlushBucketRequest) (*abfg.Response, error) {
	if r == nil {
		return PrepareGRPCResponse(false, errors.ErrBadRequest), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// AddIPToWhitelist .
func (s *ABFGServer) AddIPToWhitelist(ctx context.Context, r *abfg.SubnetRequest) (*abfg.Response, error) {
	if r == nil {
		return PrepareGRPCResponse(false, errors.ErrBadRequest), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// DeleteIPFromWhitelist .
func (s *ABFGServer) DeleteIPFromWhitelist(ctx context.Context, r *abfg.SubnetRequest) (*abfg.Response, error) {
	if r == nil {
		return PrepareGRPCResponse(false, errors.ErrBadRequest), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// AddIPToBlacklist .
func (s *ABFGServer) AddIPToBlacklist(ctx context.Context, r *abfg.SubnetRequest) (*abfg.Response, error) {
	if r == nil {
		return PrepareGRPCResponse(false, errors.ErrBadRequest), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// DeleteIPFromBlacklist .
func (s *ABFGServer) DeleteIPFromBlacklist(ctx context.Context, r *abfg.SubnetRequest) (*abfg.Response, error) {
	if r == nil {
		return PrepareGRPCResponse(false, errors.ErrBadRequest), nil
	}
	return PrepareGRPCResponse(true, nil), nil
}

// GetIPList .
func (s *ABFGServer) GetIPList(ctx context.Context, r *abfg.ListRequest) (*abfg.ListResponse, error) {
	if r == nil {
		return PrepareGRPCListIPResponse(nil, errors.ErrBadRequest)
	}
	return PrepareGRPCListIPResponse(nil, errors.ErrBadRequest)
}
