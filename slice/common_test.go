package slice

import (
	"reflect"
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

func haveSameElements[T any](x, y []T) bool {
	if len(x) != len(y) {
		return false
	}
	m := make(map[any]int)
	n := make(map[any]int)
	for i := 0; i < len(x); i++ {
		c, ok := m[x[i]]
		if !ok {
			m[c] = 0
		}
		m[c]++

		c, ok = n[x[i]]
		if !ok {
			n[c] = 0
		}
		n[c]++
	}
	return reflect.DeepEqual(m, n)
}
