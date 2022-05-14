package slice_test

import (
	"reflect"
	"testing"

	"github.com/alainrk/ftools/slice"
)

func TestFilterInt(t *testing.T) {
	l := []int{-1, 4, 0, -5, 2, 3}
	res, err := slice.Filter(l, isPositive)
	exp := []int{4, 0, 2, 3}
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestFilterFloat(t *testing.T) {
	l := []float64{-1.5, 4.3, 0, -5.7842, 2.012938, 3.0}
	res, err := slice.Filter(l, isNegative)
	exp := []float64{-1.5, -5.7842}
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestFilterString(t *testing.T) {
	l := []string{"hello world", "HI", "This is it", "WHERE IS IT"}
	res, err := slice.Filter(l, isUppercase)
	exp := []string{"HI", "WHERE IS IT"}
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestFilterEmpty(t *testing.T) {
	l := []string{"hello world", "asd", "This is it", "WHeRE IS IT"}
	res, err := slice.Filter(l, isUppercase)
	if err != nil {
		t.Error(err)
	}
	if len(res) != 0 {
		t.Errorf("Expected empty result, got %v", res)
	}
}
