package grpc

import (
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	abfg "github.com/omerkaya1/abf-guard/internal/grpc/api"
)

// PrepareGRPCResponse .
func PrepareGRPCResponse(ok bool, err error) *abfg.Response {
	if err != nil {
		return &abfg.Response{
			Result: &abfg.Response_Error{
				Error: err.Error(),
			},
		}
	}
	return &abfg.Response{
		Result: &abfg.Response_Ok{
			Ok: ok,
		},
	}
}

// PrepareGRPCListIPResponse .
func PrepareGRPCListIPResponse(IPs []string, err error) (*abfg.ListResponse, error) {
	if err != nil {
		return &abfg.ListResponse{
			Result: &abfg.ListResponse_Error{
				Error: err.Error(),
			},
		}, nil
	}
	if IPs == nil {
		return nil, errors.ErrEmptyIPList
	}
	return &abfg.ListResponse{
		Result: &abfg.ListResponse_Ips{
			Ips: &abfg.IPList{List: IPs},
		},
	}, nil
}
