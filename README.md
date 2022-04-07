# ftools

Functional tools for Go based on [Generics](https://go.dev/blog/intro-generics).

Requires `go >= 1.18`

## Prerequisites
`GO111MODULE=on` has to be setup in your env.

## How to

### Install
```sh
go get github.com/alainrk/ftools
```

```go
import "github.com/alainrk/ftools/slice"
```

### Filter
```go
ints := []int{1, 2, 3, 4, 5}
even, _ := slice.Filter(ints, func(x int) (bool, error) {
  return x%2 == 0, nil
})
fmt.Println(even) // [2 4]

```

### Map
```go
chars := []string{"a", "b", "c", "d", "e"}
upper, _ := slice.Map(chars, func(x string) (string, error) {
  return strings.ToUpper(x), nil
})
fmt.Println(upper) // [A B C D E]

```

### Reduce
```go
products := []product{
  product{Name: "a", Price: 10},
  product{Name: "b", Price: 20},
  product{Name: "c", Price: 30},
}
sumPrice, _ := slice.Reduce(products, func(s float64, p product) (float64, error) {
  return s + p.Price, nil
}, 0)
fmt.Println(sumPrice) // 60

```

### Some & Every
```go
isNegative := func(i float64) (bool, error) {
  return i < 0, nil
}

floats := []float64{-1.5, -4.3, 0, -5.7842, 2.012938, 3.0}
hasNegatives, _ := slice.Some(floats, isNegative)
allNegatives, _ := slice.Every(floats, isNegative)
fmt.Println(hasNegatives) // true
fmt.Println(allNegatives) // false
```

### Chunk
```go
ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
res := slice.Chunk(ints, 4)
fmt.Println(res) // [[1 2 3 4] [5 6 7 8] [9 10]]
```

### FindIndex
```go
ints := []int{-4, 3, 1, 5, 5, 6}
idx, _ := slice.FindIndex(ints, func(i int) (bool, error) {
  return i == 5, nil
})
fmt.Println(idx) // 3
```

### FindLastIndex
```go
ints := []int{-4, 3, 1, 5, 5, 6}
idx, _ := slice.FindLastIndex(ints, func(i int) (bool, error) {
  return i == 5, nil
})
fmt.Println(idx) // 4
```

### Intersection
```go
l1 := []int{4, 5, 6, 7, 8}
l2 := []int{1, 2, 3, 4}
i := Intersection(l1, l2)
fmt.Println(i) // [4]
```

## Example
```go
type Log struct {
	ServerID int
	Text     string
	Level    string
}

logs := []Log{
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

filterLevel := func(level string) func(Log) (bool, error) {
  return func(l Log) (bool, error) {
    return l.Level == level, nil
  }
}

removeTimestamp := func(l Log) (Log, error) {
  chunks := strings.Split(l.Text, " ")
  l.Text = strings.Join(chunks[1:], " ")
  return l, nil
}

countOccurrence := func(count map[string]int, l Log) (map[string]int, error) {
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

fmt.Println(count) // map["Peer not started": 2, "Missing log folder": 3]
```