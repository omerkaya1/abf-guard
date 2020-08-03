package log

import (
	"fmt"

	"go.uber.org/zap"
)

// InitLogger method sets up a server logger according to the specified log level
func InitLogger(level int) (*zap.Logger, error) {
	switch level {
	case 0:
		return zap.NewExample(), nil
	case 1:
		return zap.NewProduction()
	case 2:
		return zap.NewDevelopment()
	default:
		return nil, fmt.Errorf("incorrect logging level: %v", level)
	}
}
