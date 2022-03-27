package slice

import (
	"reflect"
	"strings"
	"testing"
)

func isPositive(i int) bool {
	return i >= 0
}

func isNegative(i float64) bool {
	return i < 0
}

func isUppercase(s string) bool {
	return s == strings.ToUpper(s)
}

func TestInt(t *testing.T) {
	l := []int{-1, 4, 0, -5, 2, 3}
	res, err := Filter(l, isPositive)
	exp := []int{4, 0, 2, 3}
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestFloat(t *testing.T) {
	l := []float64{-1.5, 4.3, 0, -5.7842, 2.012938, 3.0}
	res, err := Filter(l, isNegative)
	exp := []float64{-1.5, -5.7842}
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestString(t *testing.T) {
	l := []string{"hello world", "HI", "This is it", "WHERE IS IT"}
	res, err := Filter(l, isUppercase)
	exp := []string{"HI", "WHERE IS IT"}
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
