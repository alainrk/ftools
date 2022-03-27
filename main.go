package main

import "fmt"

func Filter[T any](list []T, f func(T) bool) ([]T, error) {
	var result []T
	for _, v := range list {
		if f(v) {
			result = append(result, v)
		}
	}
	return result, nil
}

func isPositive(i int) bool {
	return i > 0
}

func main() {
	l := []int{-1, 4, 0, -5, 2, 3}
	fmt.Println(Filter(l, isPositive))
}
