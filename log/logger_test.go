package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			if l, err := InitLogger(c.level); assert.Error(t, err) {
				assert.Nil(t, l)
			}
		})
	}
	t.Run("Correct log level", func(t *testing.T) {
		if l, err := InitLogger(1); assert.NoError(t, err) {
			assert.NotNil(t, l)
		}
	})
}
