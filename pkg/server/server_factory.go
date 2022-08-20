package server

import (
	"calculator/internal/core/ports"
	"calculator/pkg/router"
	"errors"
)

type server struct{}

var (
	errInvalidWebServerInstance = errors.New("invalid router server instance")
)

const (
	InstanceGin int = iota
)

// NewServerFactory sets up the router based of the instance router provided
func NewServerFactory(instance int, log ports.Logger, router *router.Router) (ports.Router, error) {
	switch instance {
	case InstanceGin:
		return newGinServer(log, router), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
