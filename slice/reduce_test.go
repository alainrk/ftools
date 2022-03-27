package slice

import (
	"strconv"
	"testing"
)

func SumInt(i, j int) (int, error) {
	return i + j, nil
}

func Join(i string, j int) (string, error) {
	return i + strconv.Itoa(j), nil
}

func TestReduceSumInt(t *testing.T) {
	l := []int{2, 2, 2, 2}
	res, err := Reduce(l, SumInt, 0)
	exp := 8
	if err != nil {
		t.Error(err)
	}
	if res != exp {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestJoin(t *testing.T) {
	l := []int{1, 2, 3, 4}
	res, err := Reduce(l, Join, "")
	exp := "1234"
	if err != nil {
		t.Error(err)
	}
	if res != exp {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
