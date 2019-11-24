package grpc

import api "github.com/omerkaya1/abf-guard/internal/grpc/api"

// PrepareGRPCAuthorisationBody forms a GRPC AuthorisationRequest object
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

// PrepareFlushBucketsGrpcRequest forms a GRPC FlushBucketRequest object
func PrepareFlushBucketsGrpcRequest(login, ip string) *api.FlushBucketRequest {
	if login == "" || ip == "" {
		return nil
	}
	return &api.FlushBucketRequest{
		Login: login,
		Ip:    ip,
	}
}

// PreparePurgeBucketGrpcRequest forms a GRPC PurgeBucketRequest object
func PreparePurgeBucketGrpcRequest(name string) *api.PurgeBucketRequest {
	if name == "" {
		return nil
	}
	return &api.PurgeBucketRequest{
		Name: name,
	}
}

// PrepareSubnetGrpcRequest forms a GRPC SubnetRequest object
func PrepareSubnetGrpcRequest(ip string, black bool) *api.SubnetRequest {
	if ip == "" {
		return nil
	}
	return &api.SubnetRequest{Ip: ip, List: black}
}

// PrepareIPListGrpcRequest forms a GRPC ListRequest object
func PrepareIPListGrpcRequest(list bool) *api.ListRequest {
	return &api.ListRequest{ListType: list}
}
