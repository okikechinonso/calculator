package services

import (
	"calculator/internal/core/domain"
	"testing"
)

func TestDivison(t *testing.T) {
	service := divisionService{}
	param := domain.Divison{10, 5}
	ans, _ := service.Divide(param)

	if ans != 2 {
		t.Fatalf("Answer %v is should be %v", ans, 2)
	}

	param.Denominator = 0
	param.Numerator = 0

	ans, err := service.Divide(param)
	if err == nil {
		t.Fatal("Should Be infinity or Indeterminate")
	}

	param.Numerator = 3
	_, err = service.Divide(param)
	if err == nil {
		t.Fatal("Should Be infinity or Indeterminate")
	}
}
