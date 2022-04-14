package slice

import (
	"reflect"
	"testing"
)

func TestIntersectionInt(t *testing.T) {
	l1 := []int{1, 2, 3, 4}
	l2 := []int{1, 2, 3, 4}

	res := Intersection(l1, l2)
	exp := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}

	l1 = []int{4, 5, 6, 7, 8}
	l2 = []int{1, 2, 3, 4}

	res = Intersection(l1, l2)
	exp = []int{4}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}

	l1 = []int{5, 6, 7, 8}
	l2 = []int{1, 2, 3, 4}

	res = Intersection(l1, l2)

	if len(res) != 0 {
		t.Errorf("Expected to be empty, got %v", res)
	}
}

func TestIntersectionStr(t *testing.T) {
	res := Intersection(
		[]string{"a", "b", "c"},
		[]string{"b", "b", "c"},
		[]string{"b", "b", "b"},
		[]string{"a", "b", "c"},
		[]string{"d", "d", "b"},
	)
	exp := []string{"b"}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestIntersectionStrEmpty(t *testing.T) {
	res := Intersection(
		[]string{"a", "b", "c"},
		[]string{"b", "b", "c"},
		[]string{"b", "b", "b"},
		[]string{"a", "b", "c"},
		[]string{"d", "d", "d"},
	)

	if len(res) != 0 {
		t.Errorf("Expected empty slice")
	}
}

func TestIntersectionStruct(t *testing.T) {
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

	res := Intersection(l1, l2, l3)
	exp := []Person{
		{"John", 30},
	}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
