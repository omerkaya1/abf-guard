package log

import (
	"fmt"

	"go.uber.org/zap"
)

// Logger is the base interface for logging
type Logger interface {
	Fatal(args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
}

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
