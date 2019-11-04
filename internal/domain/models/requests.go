package models

import abfg "github.com/omerkaya1/abf-guard/internal/grpc/api"

// Authorisation .
type Authorisation struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	IP       string `json:"ip"`
}

// NewGRPCAuthorisationBody .
func NewGRPCAuthorisationBody(login, password, ip string) *abfg.AuthorisationRequest {
	return &abfg.AuthorisationRequest{
		Login:    login,
		Password: password,
		Ip:       ip,
	}
}

// GrpcToInternalModel .
func GrpcToInternalModel() {
	return
}
