package slice

import (
	"testing"
)

func TestFindIndexInt(t *testing.T) {
	l := []int{-1, 4, 0, -5, 2, 3}
	res, err := FindIndex(l, func(i int) (bool, error) { return i == 4, nil })
	if err != nil {
		t.Error(err)
	}
	if res != 1 {
		t.Errorf("Failed TestFindIndexInt")
	}
}

func TestFindIndexFloatNeg(t *testing.T) {
	l := []float64{1, 4, 0, -5, 2, 3}
	res, err := FindIndex(l, isNegative)
	if err != nil {
		t.Error(err)
	}
	if res != 3 {
		t.Errorf("Failed TestFindIndexFloatNeg")
	}
}

func TestFindIndexString(t *testing.T) {
	l := []string{"hello", "world", "this", "is", "it"}
	res, err := FindIndex(l, isEmptyString)
	if err != nil {
		t.Error(err)
	}
	if res != -1 {
		t.Errorf("Failed TestFindIndexString")
	}
}

func TestFindLastIndexInt(t *testing.T) {
	l := []int{-1, 4, 0, -5, 2, 4}
	res, err := FindLastIndex(l, func(i int) (bool, error) { return i == 4, nil })
	if err != nil {
		t.Error(err)
	}
	if res != 5 {
		t.Errorf("Failed TestFindLastIndexInt")
	}
}

func TestFindLastIndexFloatNeg(t *testing.T) {
	l := []float64{-1, -4, -0, -5, 2, 3}
	res, err := FindLastIndex(l, isNegative)
	if err != nil {
		t.Error(err)
	}
	if res != 3 {
		t.Errorf("Failed TestFindLastIndexFloatNeg")
	}
}

func TestFindLastIndexString(t *testing.T) {
	l := []string{"hello", "", "this", "", "it"}
	res, err := FindLastIndex(l, isEmptyString)
	if err != nil {
		t.Error(err)
	}
	if res != 3 {
		t.Errorf("Failed TestFindLastIndexString")
	}
}
