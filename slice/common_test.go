package slice_test

import (
	"strings"
)

func isPositive(i int) (bool, error) {
	return i >= 0, nil
}

func isNegative(i float64) (bool, error) {
	return i < 0, nil
}

func isEmptyString(i string) (bool, error) {
	return len(i) == 0, nil
}

func isUppercase(s string) (bool, error) {
	return s == strings.ToUpper(s), nil
}

func Abs(i int) (int, error) {
	if i < 0 {
		return -i, nil
	}
	return i, nil
}
