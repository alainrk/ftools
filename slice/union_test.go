package slice_test

import (
	"testing"

	"github.com/alainrk/ftools/slice"
)

func TestUnionInt(t *testing.T) {
	l1 := []int{1, 2, 3, 4}
	l2 := []int{1, 2, 3, 4}

	res := slice.Union(l1, l2)
	exp := []int{1, 2, 3, 4}

	if !slice.SameElements(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}

	l1 = []int{4, 5, 6, 7, 8, 8, 8, 8, 8, 8}
	l2 = []int{1, 2, 3, 4}

	res = slice.Union(l1, l2)
	exp = []int{1, 2, 3, 4, 5, 6, 7, 8}

	if !slice.SameElements(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestUnionStr(t *testing.T) {
	res := slice.Union(
		[]string{"a", "b", "c"},
		[]string{"b", "b", "c"},
		[]string{"b", "b", "b"},
		[]string{"a", "b", "c"},
		[]string{"d", "d", "b"},
	)
	exp := []string{"a", "b", "c", "d"}

	if !slice.SameElements(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestUnionStrEmpty(t *testing.T) {
	res := slice.Union(
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
	)

	if len(res) != 0 {
		t.Errorf("Expected empty slice")
	}
}

func TestUnionStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	l1 := []Person{
		{"John", 30},
		{"Jane", 25},
		{"Jack", 40},
	}
	l2 := []Person{
		{"John", 30},
		{"Chris", 25},
		{"Jack", 40},
	}
	l3 := []Person{
		{"Antony", 12},
		{"Jack", 45},
		{"John", 30},
	}

	res := slice.Union(l1, l2, l3)
	exp := []Person{
		{"John", 30},
		{"Jane", 25},
		{"Chris", 25},
		{"Jack", 40},
		{"Antony", 12},
		{"Jack", 45},
	}

	if !slice.SameElements(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
