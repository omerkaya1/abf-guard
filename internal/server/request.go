package server

import abfg "github.com/omerkaya1/abf-guard/internal/server/api"

// Authorisation .
type Authorisation struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	IP       string `json:"ip"`
}

// PrepareGRPCAuthorisationBody .
func PrepareGRPCAuthorisationBody(login, password, ip string) *abfg.AuthorisationRequest {
	if login == "" || password == "" || ip == "" {
		return nil
	}
	return &abfg.AuthorisationRequest{
		Login:    login,
		Password: password,
		Ip:       ip,
	}
}

// PrepareFlushBucketGrpcRequest .
func PrepareFlushBucketGrpcRequest(login, ip string) *abfg.FlushBucketRequest {
	if login == "" || ip == "" {
		return nil
	}
	return &abfg.FlushBucketRequest{
		Login: login,
		Ip:    ip,
	}
}

// PrepareSubnetGrpcRequest .
func PrepareSubnetGrpcRequest(ip string, black bool) *abfg.SubnetRequest {
	if ip == "" {
		return nil
	}
	return &abfg.SubnetRequest{Ip: ip, List: black}
}

// PrepareIPListGrpcRequest .
func PrepareIPListGrpcRequest(list bool) *abfg.ListRequest {
	return &abfg.ListRequest{ListType: list}
}
