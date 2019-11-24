package grpc

import api "github.com/omerkaya1/abf-guard/internal/grpc/api"

// Authorisation .
type Authorisation struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	IP       string `json:"ip"`
}

// PrepareGRPCAuthorisationBody .
func PrepareGRPCAuthorisationBody(login, password, ip string) *api.AuthorisationRequest {
	if login == "" || password == "" || ip == "" {
		return nil
	}
	return &api.AuthorisationRequest{
		Login:    login,
		Password: password,
		Ip:       ip,
	}
}

// PrepareFlushBucketGrpcRequest .
func PrepareFlushBucketsGrpcRequest(login, ip string) *api.FlushBucketRequest {
	if login == "" || ip == "" {
		return nil
	}
	return &api.FlushBucketRequest{
		Login: login,
		Ip:    ip,
	}
}

// PreparePurgeBucketGrpcRequest .
func PreparePurgeBucketGrpcRequest(name string) *api.PurgeBucketRequest {
	if name == "" {
		return nil
	}
	return &api.PurgeBucketRequest{
		Name: name,
	}
}

// PrepareSubnetGrpcRequest .
func PrepareSubnetGrpcRequest(ip string, black bool) *api.SubnetRequest {
	if ip == "" {
		return nil
	}
	return &api.SubnetRequest{Ip: ip, List: black}
}

// PrepareIPListGrpcRequest .
func PrepareIPListGrpcRequest(list bool) *api.ListRequest {
	return &api.ListRequest{ListType: list}
}
