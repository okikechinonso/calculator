package services

import (
	"calculator/internal/core/domain"
	"calculator/internal/core/ports"
	"errors"
	"math"
)

type divisionService struct{}

func (d divisionService) Divide(payload domain.Divison) (interface{}, error) {
	if payload.Denominator == 0 || (payload.Numerator == 0 && payload.Denominator == 0) {
		return nil, errors.New("Infinity or interdeminate")
	}
	ans := float64(payload.Numerator) / payload.Denominator
	return roundFloat(ans, 4), nil
}

func NewDivisionService() ports.DivisionService {
	return &divisionService{}
}
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
