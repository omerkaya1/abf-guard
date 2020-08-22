package domain

import "errors"

var (
	// ErrEmptyIP .
	ErrEmptyIP = errors.New("empty IP is received")
	// ErrEmptyIP .
	ErrEmptyLogin = errors.New("empty login is received")
	// ErrEmptyIP .
	ErrEmptyPWD = errors.New("empty password is received")
)

func validateAuthorisationParams(login, password, ip string) error {
	if login == "" {
		return ErrEmptyLogin
	}
	if password == "" {
		return ErrEmptyPWD
	}
	if ip == "" {
		return ErrEmptyIP
	}
	return nil
}

func validateFlashParams(login, ip string) error {
	if login == "" {
		return ErrEmptyLogin
	}
	if ip == "" {
		return ErrEmptyIP
	}
	return nil
}

func prepareAuthorisationMap(login, password, ip string) map[string]int {
	return map[string]int{login: 0, password: 1, ip: 2}
}
