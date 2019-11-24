package grpc

import (
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	api "github.com/omerkaya1/abf-guard/internal/grpc/api"
)

// PrepareGRPCResponse forms a GRPC Response object
func PrepareGRPCResponse(ok bool, err error) *api.Response {
	if err != nil {
		return &api.Response{
			Result: &api.Response_Error{
				Error: err.Error(),
			},
		}
	}
	return &api.Response{
		Result: &api.Response_Ok{
			Ok: ok,
		},
	}
}

// PrepareGRPCListIPResponse forms a GRPC ListResponse object
func PrepareGRPCListIPResponse(IPs []string, err error) (*api.ListResponse, error) {
	if err != nil {
		return &api.ListResponse{
			Result: &api.ListResponse_Error{
				Error: err.Error(),
			},
		}, nil
	}
	if IPs == nil {
		return nil, errors.ErrEmptyIPList
	}
	return &api.ListResponse{
		Result: &api.ListResponse_Ips{
			Ips: &api.IPList{List: IPs},
		},
	}, nil
}
