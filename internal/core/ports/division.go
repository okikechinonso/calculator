package ports

import (
	"calculator/internal/core/domain"
	"github.com/gin-gonic/gin"
)

// DivisionService is a service interface for the core to communicate with the adapters' user handlers .
type DivisionService interface {
	Divide(payload domain.Divison) (interface{}, error)
}

// DivisionHandler is a user handler interface for request and response handler .
type DivisionHandler interface {
	Division(c *gin.Context)
}
