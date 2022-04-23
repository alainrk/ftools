package slice

import (
	"reflect"
	"testing"
)

func TestPull1(t *testing.T) {
	l := []int{-1, 4, 0, -5, 2, 3}
	res := Pull(l, -1, 4, 0, 2)
	exp := []int{-5, 3}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestPull2(t *testing.T) {
	l := []int{-1, 4, 0, -5, 2, 3}
	res := Pull(l, 0)
	exp := []int{-1, 4, -5, 2, 3}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestPull3(t *testing.T) {
	l := []int{-1, 4, 0, -5, 2, 3}
	res := Pull(l, -1, 4, 0, -5, 2, 3)
	if len(res) != 0 {
		t.Errorf("Expected empty result, given %+v", res)
	}
}

func TestPull4(t *testing.T) {
	l := []int{}
	res := Pull(l, 100, 200, 300)
	if len(res) != 0 {
		t.Errorf("Expected empty result, given %+v", res)
	}
}
