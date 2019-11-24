package manager

import "github.com/omerkaya1/abf-guard/internal/domain/errors"

func validateAuthorisationParams(login, password, ip string) error {
	if login == "" {
		return errors.ErrEmptyIP
	}
	if password == "" {
		return errors.ErrEmptyPWD
	}
	if ip == "" {
		return errors.ErrEmptyLogin
	}
	return nil
}

func validateFlashParams(login, ip string) error {
	if login == "" {
		return errors.ErrEmptyIP
	}
	if ip == "" {
		return errors.ErrEmptyIP
	}
	return nil
}

func prepareAuthorisationMap(login, password, ip string) map[string]int {
	return map[string]int{login: 0, password: 1, ip: 2}
}
