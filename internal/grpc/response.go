package grpc

import (
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	abfg "github.com/omerkaya1/abf-guard/internal/grpc/api"
)

// FlushBucketModelToGrpc .
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

// PrepareGRPCListIpResponse .
func PrepareGRPCListIpResponse(ips []string, err error) (*abfg.ListResponse, error) {
	if err != nil {
		return &abfg.ListResponse{
			Result: &abfg.ListResponse_Error{
				Error: err.Error(),
			},
		}, nil
	}
	if ips == nil {
		return nil, errors.ErrEmptyIpList
	}
	return &abfg.ListResponse{
		Result: &abfg.ListResponse_Ips{
			Ips: &abfg.IpList{List: ips},
		},
	}, nil
}
