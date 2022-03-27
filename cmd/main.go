package main

import (
	"fmt"
	"strings"

	"github.com/alainrk/ftools/slice"
)

type product struct {
	Name  string
	Price float64
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	chars := []string{"a", "b", "c", "d", "e"}
	products := []product{
		product{Name: "a", Price: 10},
		product{Name: "b", Price: 20},
		product{Name: "c", Price: 30},
	}

	// Filter
	even, _ := slice.Filter(ints, func(x int) (bool, error) {
		return x%2 == 0, nil
	})
	fmt.Println(even)

	// Map
	upper, _ := slice.Map(chars, func(x string) (string, error) {
		return strings.ToUpper(x), nil
	})
	fmt.Println(upper)

	// Reduce
	sumPrice, _ := slice.Reduce(products, func(s float64, p product) (float64, error) {
		return s + p.Price, nil
	}, 0)
	fmt.Println(sumPrice)
}
