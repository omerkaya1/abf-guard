package grpc

import (
	abfg "github.com/omerkaya1/abf-guard/internal/grpc/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrepareGRPCAuthorisationBody(t *testing.T) {
	testCases := []struct {
		header   string
		login    string
		pwd      string
		ip       string
		response *abfg.AuthorisationRequest
	}{
		{"Empty args", "", "", "", nil},
		{"Empty login", "", "1234", "111.111.111.111", nil},
		{"Empty password", "petya", "", "111.111.111.111", nil},
		{"Empty ip", "petya", "1234", "", nil},
		{"Args are present", "petya", "1234", "111.111.111.111", &abfg.AuthorisationRequest{
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
		response *abfg.FlushBucketRequest
	}{
		{"Empty args", "", "", nil},
		{"Empty login", "", "111.111.111.111", nil},
		{"Empty ip", "petya", "", nil},
		{"Args are present", "petya", "111.111.111.111", &abfg.FlushBucketRequest{
			Login: "petya",
			Ip:    "111.111.111.111",
		}},
	}
	for _, c := range testCases {
		t.Run(c.header, func(t *testing.T) {
			assert.Equal(t, c.response, PrepareFlushBucketGrpcRequest(c.login, c.ip))
		})
	}
}

func TestPrepareSubnetGrpcRequest(t *testing.T) {
	testCases := []struct {
		header   string
		ip       string
		response *abfg.SubnetRequest
	}{
		{"Empty ip", "", nil},
		{"Args are present", "111.111.111.111", &abfg.SubnetRequest{
			Ip: "111.111.111.111",
		}},
	}
	for _, c := range testCases {
		t.Run(c.header, func(t *testing.T) {
			assert.Equal(t, c.response, PrepareSubnetGrpcRequest(c.ip))
		})
	}
}
