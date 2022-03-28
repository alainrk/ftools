# ftools

Functional tools for Go based on [Generics](https://go.dev/blog/intro-generics).

Requires `go >= 1.18`

## Prerequisites
`GO111MODULE=on` has to be setup in your env.

## Example
```go
import "github.com/alainrk/ftools/slice"

// Filter
ints := []int{1, 2, 3, 4, 5}
even, _ := slice.Filter(ints, func(x int) (bool, error) {
  return x%2 == 0, nil
})
fmt.Println(even) // [2 4]

// Map
chars := []string{"a", "b", "c", "d", "e"}
upper, _ := slice.Map(chars, func(x string) (string, error) {
  return strings.ToUpper(x), nil
})
fmt.Println(upper) // [A B C D E]

// Reduce
products := []product{
  product{Name: "a", Price: 10},
  product{Name: "b", Price: 20},
  product{Name: "c", Price: 30},
}
sumPrice, _ := slice.Reduce(products, func(s float64, p product) (float64, error) {
  return s + p.Price, nil
}, 0)
fmt.Println(sumPrice) // 60

// Some & Every
isNegative := func(i float64) (bool, error) {
  return i < 0, nil
}

floats := []float64{-1.5, -4.3, 0, -5.7842, 2.012938, 3.0}
hasNegatives, _ := slice.Some(floats, isNegative)
allNegatives, _ := slice.Every(floats, isNegative)
fmt.Println(hasNegatives) // true
fmt.Println(allNegatives) // false
```