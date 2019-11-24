package manager

import "github.com/omerkaya1/abf-guard/internal/domain/errors"

// TODO: Error handling!!!
// ValidateAuthorisationParams .
func ValidateAuthorisationParams(login, password, ip string) error {
	if login == "" {
		return errors.ErrEmptyIP
	}
	if password == "" {
		return errors.ErrEmptyIP
	}
	if ip == "" {
		return errors.ErrEmptyIP
	}
	return nil
}

// ValidateFlashParams .
func ValidateFlashParams(login, ip string) error {
	if login == "" {
		return errors.ErrEmptyIP
	}
	if ip == "" {
		return errors.ErrEmptyIP
	}
	return nil
}

// PrepareAuthorisationMap .
func PrepareAuthorisationMap(login, password, ip string) map[string]int {
	return map[string]int{login: 0, password: 1, ip: 2}
}
