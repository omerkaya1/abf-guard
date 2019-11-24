package integration

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	status := 0
	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)

}
