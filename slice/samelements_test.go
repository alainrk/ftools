package slice_test

import (
	"testing"

	"github.com/alainrk/ftools/slice"
)

func TestSameElementsInt(t *testing.T) {
	l1 := []int{1, 2, 3, 4}
	l2 := []int{4, 3, 2, 1}
	l3 := []int{2, 1, 3, 4}
	l4 := []int{1, 1, 3, 4}

	res := slice.SameElements(l1)
	exp := true

	if res != exp {
		t.Errorf("Expected %v, got %v", exp, res)
	}

	res = slice.SameElements(l1, l2, l3)
	exp = true

	if res != exp {
		t.Errorf("Expected %v, got %v", exp, res)
	}

	res = slice.SameElements(l1, l2, l3, l4)
	exp = false

	if res != exp {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
