package slice

import (
	"reflect"
	"strconv"
	"testing"
)

func Abs(i int) (int, error) {
	if i < 0 {
		return -i, nil
	}
	return i, nil
}

func TestMapAbs(t *testing.T) {
	l := []int{-1, 4, 0, -5, 2, 3}
	res, err := Map(l, Abs)
	exp := []int{1, 4, 0, 5, 2, 3}
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestMapEmpty(t *testing.T) {
	l := []int{}
	res, err := Map(l, Abs)
	if err != nil {
		t.Error(err)
	}
	if len(res) != 0 {
		t.Errorf("Expected empty slice, got %v", res)
	}
}

func TestMapToString(t *testing.T) {
	l := []int{-1, 4, 0, -5, 2, 3}
	res, err := Map(l, func(i int) (string, error) { return strconv.Itoa(i), nil })
	exp := []string{"-1", "4", "0", "-5", "2", "3"}
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
