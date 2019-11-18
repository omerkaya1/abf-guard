package grpc

import (
	"fmt"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	abfg "github.com/omerkaya1/abf-guard/internal/grpc/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrepareGRPCListIpResponse(t *testing.T) {
	testCases := []struct {
		header   string
		ips      []string
		err      error
		response *abfg.ListResponse
		err1     error
	}{
		{"Error", nil, fmt.Errorf("some error"), &abfg.ListResponse{
			Result: &abfg.ListResponse_Error{Error: "some error"},
		}, nil},
		{"Empty ips slice", nil, nil, nil, errors.ErrEmptyIpList},
		{
			"Error is present, some ips retrieved",
			[]string{"111.111.111.111", "111.111.111.112"},
			fmt.Errorf("some error"),
			&abfg.ListResponse{
				Result: &abfg.ListResponse_Error{
					Error: "some error",
				},
			}, nil},
		{
			"Response with a list of ips",
			[]string{"111.111.111.111", "111.111.111.112"},
			nil,
			&abfg.ListResponse{
				Result: &abfg.ListResponse_Ips{
					Ips: &abfg.IpList{List: []string{"111.111.111.111", "111.111.111.112"}},
				},
			}, nil},
	}
	for _, c := range testCases {
		t.Run(c.header, func(t *testing.T) {
			r, err := PrepareGRPCListIpResponse(c.ips, c.err)
			if err == errors.ErrEmptyIpList {
				assert.Equal(t, c.response, r)
				return
			}
			if assert.NoError(t, err) {
				assert.Equal(t, c.response, r)
			}
		})
	}
}

func TestPrepareGRPCResponse(t *testing.T) {
	testCases := []struct {
		header   string
		ok       bool
		err      error
		response *abfg.Response
	}{
		{"No error", true, nil, &abfg.Response{Result: &abfg.Response_Ok{Ok: true}}},
		{"Error", false, fmt.Errorf("some error"), &abfg.Response{Result: &abfg.Response_Error{Error: "some error"}}},
	}
	for _, c := range testCases {
		t.Run(c.header, func(t *testing.T) {
			assert.Equal(t, c.response, PrepareGRPCResponse(c.ok, c.err))
		})
	}
}
