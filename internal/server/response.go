package server

import (
	"errors"

	"github.com/omerkaya1/abf-guard/internal/server/api"
)

var errEmptyIPList = errors.New("empty ip list received")

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
		return nil, errEmptyIPList
	}
	return &api.ListResponse{
		Result: &api.ListResponse_Ips{
			Ips: &api.IPList{List: IPs},
		},
	}, nil
}
