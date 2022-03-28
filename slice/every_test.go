package slice

import (
	"testing"
)

func TestEveryInt(t *testing.T) {
	l := []int{1, 4, -5, 2, 3}
	res, err := Every(l, isPositive)
	exp := false
	if err != nil {
		t.Error(err)
	}
	if res != exp {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestEveryFloat(t *testing.T) {
	l := []float64{-1.5, -4.3, -5.7842, -2.012938, -3.0}
	res, err := Every(l, isNegative)
	exp := true
	if err != nil {
		t.Error(err)
	}
	if res != exp {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestEveryString(t *testing.T) {
	l := []string{"hello world", "HI", "This is it", "WHERE IS IT"}
	res, err := Every(l, isUppercase)
	exp := false
	if err != nil {
		t.Error(err)
	}
	if res != exp {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
