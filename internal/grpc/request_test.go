package grpc

import (
	"testing"

	"github.com/omerkaya1/abf-guard/internal/grpc/api"
	"github.com/stretchr/testify/assert"
)

func TestPrepareGRPCAuthorisationBody(t *testing.T) {
	testCases := []struct {
		header   string
		login    string
		pwd      string
		ip       string
		response *api.AuthorisationRequest
	}{
		{"Empty args", "", "", "", nil},
		{"Empty login", "", "1234", "111.111.111.111", nil},
		{"Empty password", "petya", "", "111.111.111.111", nil},
		{"Empty ip", "petya", "1234", "", nil},
		{"Args are present", "petya", "1234", "111.111.111.111", &api.AuthorisationRequest{
			Login:    "petya",
			Password: "1234",
			Ip:       "111.111.111.111",
		}},
	}
	for _, c := range testCases {
		t.Run(c.header, func(t *testing.T) {
			assert.Equal(t, c.response, PrepareGRPCAuthorisationBody(c.login, c.pwd, c.ip))
		})
	}
}

func TestPrepareFlushBucketGrpcRequest(t *testing.T) {
	testCases := []struct {
		header   string
		login    string
		ip       string
		response *api.FlushBucketRequest
	}{
		{"Empty args", "", "", nil},
		{"Empty login", "", "111.111.111.111", nil},
		{"Empty ip", "petya", "", nil},
		{"Args are present", "petya", "111.111.111.111", &api.FlushBucketRequest{
			Login: "petya",
			Ip:    "111.111.111.111",
		}},
	}
	for _, c := range testCases {
		t.Run(c.header, func(t *testing.T) {
			assert.Equal(t, c.response, PrepareFlushBucketsGrpcRequest(c.login, c.ip))
		})
	}
}

func TestPrepareSubnetGrpcRequest(t *testing.T) {
	testCases := []struct {
		header   string
		ip       string
		result   bool
		response *api.SubnetRequest
	}{
		{"Empty ip", "", false, nil},
		{"Args are present", "111.111.111.111", true, &api.SubnetRequest{
			Ip:   "111.111.111.111",
			List: true,
		}},
	}
	for _, c := range testCases {
		t.Run(c.header, func(t *testing.T) {
			assert.Equal(t, c.response, PrepareSubnetGrpcRequest(c.ip, c.result))
		})
	}
}

func TestPreparePurgeBucketGrpcRequest(t *testing.T) {
	testCases := []struct {
		header   string
		name     string
		response *api.PurgeBucketRequest
	}{
		{"Empty name", "", nil},
		{"Args are present", "morty", &api.PurgeBucketRequest{
			Name: "morty",
		}},
	}
	for _, c := range testCases {
		t.Run(c.header, func(t *testing.T) {
			assert.Equal(t, c.response, PreparePurgeBucketGrpcRequest(c.name))
		})
	}
}
