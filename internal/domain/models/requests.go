package models

import abfg "github.com/omerkaya1/abf-guard/internal/grpc/api"

type Authorisation struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	IP       string `json:"ip"`
}

type AuthRequestConverter struct {
}

func NewGRPCAuthorisationBody(login, password, ip string) *abfg.AuthorisationRequest {
	return &abfg.AuthorisationRequest{
		Login:    login,
		Password: password,
		Ip:       ip,
	}
}

func (arc *AuthRequestConverter) GrpcToInternalModel() {

}
