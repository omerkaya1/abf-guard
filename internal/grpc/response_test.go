package grpc

import (
	"fmt"
	"testing"

	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/grpc/api"
	"github.com/stretchr/testify/assert"
)

func TestPrepareGRPCListIpResponse(t *testing.T) {
	testCases := []struct {
		header   string
		ips      []string
		err      error
		response *api.ListResponse
		err1     error
	}{
		{"Error", nil, fmt.Errorf("some error"), &api.ListResponse{
			Result: &api.ListResponse_Error{Error: "some error"},
		}, nil},
		{"Empty ips slice", nil, nil, nil, errors.ErrEmptyIPList},
		{
			"Error is present, some ips retrieved",
			[]string{"111.111.111.111", "111.111.111.112"},
			fmt.Errorf("some error"),
			&api.ListResponse{
				Result: &api.ListResponse_Error{
					Error: "some error",
				},
			}, nil},
		{
			"Response with a list of ips",
			[]string{"111.111.111.111", "111.111.111.112"},
			nil,
			&api.ListResponse{
				Result: &api.ListResponse_Ips{
					Ips: &api.IPList{List: []string{"111.111.111.111", "111.111.111.112"}},
				},
			}, nil},
	}
	for _, c := range testCases {
		t.Run(c.header, func(t *testing.T) {
			r, err := PrepareGRPCListIPResponse(c.ips, c.err)
			if err == errors.ErrEmptyIPList {
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
		response *api.Response
	}{
		{"No error", true, nil, &api.Response{Result: &api.Response_Ok{Ok: true}}},
		{"Error", false, fmt.Errorf("some error"), &api.Response{Result: &api.Response_Error{Error: "some error"}}},
	}
	for _, c := range testCases {
		t.Run(c.header, func(t *testing.T) {
			assert.Equal(t, c.response, PrepareGRPCResponse(c.ok, c.err))
		})
	}
}
