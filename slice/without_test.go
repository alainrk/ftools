package slice_test

import (
	"reflect"
	"testing"

	"github.com/alainrk/ftools/slice"
)

func TestWithoutInt(t *testing.T) {
	l := []int{1, 2, 3, 4}

	res := slice.Without(l, 1)
	exp := []int{2, 3, 4}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}

	res = slice.Without(l, 1, 2)
	exp = []int{3, 4}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestWithoutEmpty(t *testing.T) {
	l := []int{1, 2, 3, 4}

	res := slice.Without(l, 1, 2, 3, 4)

	if len(res) != 0 {
		t.Errorf("Expected empty slice, got %v", res)
	}
}
