package slice

import (
	"reflect"
	"testing"
)

func isPositive(i int) bool {
	return i >= 0
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
