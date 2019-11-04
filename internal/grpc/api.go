package grpc

import (
	"context"

	abfg "github.com/omerkaya1/abf-guard/internal/grpc/api"
)

func (s *ABFGServer) Authorisation(ctx context.Context, r *abfg.AuthorisationRequest) (*abfg.SimpleResponse, error) {
	panic("implement me")
}

func (s *ABFGServer) FlashBucket(ctx context.Context, r *abfg.FlashBucketRequest) (*abfg.SimpleResponse, error) {
	panic("implement me")
}

func (s *ABFGServer) AddIpToWhitelist(ctx context.Context, r *abfg.SubnetRequest) (*abfg.SimpleResponse, error) {
	panic("implement me")
}

func (s *ABFGServer) DeleteIpFromWhitelist(ctx context.Context, r *abfg.SubnetRequest) (*abfg.SimpleResponse, error) {
	panic("implement me")
}

func (s *ABFGServer) AddIpToBlacklist(ctx context.Context, r *abfg.SubnetRequest) (*abfg.SimpleResponse, error) {
	panic("implement me")
}

func (s *ABFGServer) DeleteIpFromBlacklist(ctx context.Context, r *abfg.SubnetRequest) (*abfg.SimpleResponse, error) {
	panic("implement me")
}
