package handler

import (
	"calculator/internal/core/domain"
	"calculator/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type divisionHandler struct {
	userService ports.DivisionService
	logger      ports.Logger
}

func NewDivsionHandler(divisionServe ports.DivisionService, l ports.Logger) ports.DivisionHandler {
	return &divisionHandler{
		userService: divisionServe,
		logger:      l,
	}
}

func (h *divisionHandler) Division(c *gin.Context) {
	param := &domain.Divison{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		h.logger.Errorf("%v", err.Error())
		c.JSON(500, gin.H{"message": "unable to decode payload"})
		return
	}
	ans, err := h.userService.Divide(*param)
	if err != nil {
		h.logger.Errorf("Invalid input", err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"answer": ans})
	h.logger.Infof("The answer to %v / %v is %v", param.Numerator, param.Denominator, ans)

}
