package log

import (
	"calculator/internal/core/ports"
	"errors"
)

const (
	InstanceLogrusLogger int = iota
)

var (
	errInvalidLoggerInstance = errors.New("invalid log instance")
)

// NewLoggerFactory returns a logger interface type based on the provided logger instance
func NewLoggerFactory(instance int) (ports.Logger, error) {
	switch instance {
	case InstanceLogrusLogger:
		return NewLogrusLogger(), nil
	default:
		return nil, errInvalidLoggerInstance
	}
}
