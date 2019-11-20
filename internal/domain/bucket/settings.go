package bucket

import "time"

// Settings .
type Settings struct {
	LoginLimit    int
	PasswordLimit int
	IPLimit       int
	Expire        time.Duration
}
