package log

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitLogger(t *testing.T) {
	testCases := []struct {
		header string
		level  int
	}{
		{"Negative level", -1},
		{"Above allowed", 16419},
	}
	for _, c := range testCases {
		t.Run(c.header, func(t *testing.T) {
			l, err := InitLogger(c.level)
			require.Error(t, err)
			require.Nil(t, l)
		})
	}
	for i := 0; i < 3; i++ {
		t.Run("Correct log level", func(t *testing.T) {
			l, err := InitLogger(i)
			require.NoError(t, err)
			require.NotNil(t, l)
		})
	}
}
