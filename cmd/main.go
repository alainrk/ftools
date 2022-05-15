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

type log struct {
	ServerID int
	Text     string
	Level    string
}

func main() {
	// Test
	logs := []log{
		{1, "1648449331 Proc started", "INFO"},
		{2, "1648449331 Proc started", "INFO"},
		{2, "1648449331 Missing log folder", "WARNING"},
		{3, "1648449331 Proc started", "INFO"},
		{1, "1648449331 Missing log folder", "WARNING"},
		{3, "1648449331 Missing log folder", "WARNING"},
		{1, "1648449331 Creating log folder", "DEBUG"},
		{2, "1648449331 Peer not started", "WARNING"},
		{1, "1648449331 Peer not started", "WARNING"},
		{3, "1648449331 Creating log folder", "DEBUG"},
	}

	filterLevel := func(level string) func(log) (bool, error) {
		return func(l log) (bool, error) {
			return l.Level == level, nil
		}
	}

	removeTimestamp := func(l log) (log, error) {
		chunks := strings.Split(l.Text, " ")
		l.Text = strings.Join(chunks[1:], " ")
		return l, nil
	}

	countOccurrence := func(count map[string]int, l log) (map[string]int, error) {
		if _, ok := count[l.Text]; ok {
			count[l.Text] += 1
		} else {
			count[l.Text] = 1
		}
		return count, nil
	}

	var count = map[string]int{}

	logs, _ = slice.Filter(logs, filterLevel("WARNING"))
	logs, _ = slice.Map(logs, removeTimestamp)
	count, _ = slice.Reduce(logs, countOccurrence, count)
	fmt.Println(count)

	ints := []int{1, 2, 3, 4, 5}
	chars := []string{"a", "b", "c", "d", "e"}
	products := []product{
		{Name: "a", Price: 10},
		{Name: "b", Price: 20},
		{Name: "c", Price: 30},
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

	// Some
	isNegative := func(i float64) (bool, error) {
		return i < 0, nil
	}

	floats := []float64{-1.5, -4.3, 0, -5.7842, 2.012938, 3.0}
	hasNegatives, _ := slice.Some(floats, isNegative)
	allNegatives, _ := slice.Every(floats, isNegative)
	fmt.Println(hasNegatives) // true
	fmt.Println(allNegatives) // false

	// Chunk
	ints = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := slice.Chunk(ints, 4)
	fmt.Println(res) // [[1 2 3 4] [5 6 7 8] [9 10]]

	// FindIndex
	fi := []int{-4, 3, 1, 5, 5, 6}
	idx, _ := slice.FindIndex(fi, func(i int) (bool, error) {
		return i == 5, nil
	})
	fmt.Println(idx) // 3

	// SameElements
	l := []int{3, 4, 2, 2}
	p := []int{2, 3, 4, 2}
	q := []int{3, 2, 2, 4}

	r := []int{2, 3, 4, 1}

	fmt.Println(slice.SameElements(l, p, q))    // true
	fmt.Println(slice.SameElements(l, p, q, r)) // false
}
