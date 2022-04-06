package slice

import (
	"fmt"
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

func TestIntersectionStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	l1 := []Person{
		{"John", 30},
		{"Jane", 25},
		{"Jack", 41},
	}
	l2 := []Person{
		{"John", 30},
		{"Chris", 25},
		{"Jack", 40},
	}

	res := Intersection(l1, l2)
	exp := []Person{
		{"John", 30},
		{"Jack", 40},
	}

	fmt.Printf("Expected %v, got %v", exp, res)

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
