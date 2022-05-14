package slice

// Filter returns a new slice containing all elements that satisfy the predicate.
func Filter[T any](list []T, f func(T) (bool, error)) ([]T, error) {
	var result []T
	for _, v := range list {
		ok, err := f(v)
		if err != nil {
			return nil, err
		}
		if ok {
			result = append(result, v)
		}
	}
	return result, nil
}

// Map returns a new slice applying the function to each element.
func Map[T any, U any](list []T, f func(T) (U, error)) ([]U, error) {
	var result []U
	for _, v := range list {
		if x, err := f(v); err == nil {
			result = append(result, x)
		} else {
			return nil, err
		}
	}
	return result, nil
}

// Reduce returns a single value by applying the function to each element in the list
// and uses the base as the initial value.
func Reduce[T any, U any](list []T, f func(U, T) (U, error), base U) (U, error) {
	for _, v := range list {
		if b, err := f(base, v); err != nil {
			return base, err
		} else {
			base = b
		}
	}
	return base, nil
}

// Some returns true if any element in the list satisfies the predicate.
func Some[T any](list []T, f func(T) (bool, error)) (bool, error) {
	for _, v := range list {
		ok, err := f(v)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}
	return false, nil
}

// Every returns true if all the elements in the list satisfies the predicate.
func Every[T any](list []T, f func(T) (bool, error)) (bool, error) {
	for _, v := range list {
		ok, err := f(v)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil
}

// Chunks returns a slice of slices of size n, keeping the difference at the end if
// not equally divisible.
func Chunk[T any](list []T, size int) [][]T {
	var result [][]T
	if size <= 0 {
		return result
	}

	for i := 0; i < len(list); i += size {
		end := i + size
		if end > len(list) {
			end = len(list)
		}
		result = append(result, list[i:end])
	}
	return result
}

// FindIndex returns the index of the first element in the list that satisfies the predicate.
func FindIndex[T any](list []T, f func(T) (bool, error)) (int, error) {
	for i, v := range list {
		ok, err := f(v)
		if err != nil {
			return -1, err
		}
		if ok {
			return i, nil
		}
	}
	return -1, nil
}

// FindLastIndex returns the index of the last element in the list that satisfies the predicate.
func FindLastIndex[T any](list []T, f func(T) (bool, error)) (int, error) {
	for i := len(list) - 1; i >= 0; i-- {
		ok, err := f(list[i])
		if err != nil {
			return -1, err
		}
		if ok {
			return i, nil
		}
	}
	return -1, nil
}

// Intersection returns a slice with the common elements of the given slices.
func intersect[T any](l1, l2 []T) []T {
	s := map[any]bool{}
	for _, v := range l1 {
		s[v] = true
	}
	var intersection []T
	for _, v := range l2 {
		if _, ok := s[v]; ok {
			intersection = append(intersection, v)
		}
	}
	return intersection
}

// Intersection returns a slice with the common elements of the given slices.
func Intersection[T any](l ...[]T) []T {
	if len(l) == 0 {
		return []T{}
	}

	intersection := l[0]
	for i := 1; i < len(l); i++ {
		intersection = intersect(intersection, l[i])
	}

	return intersection
}

// Pull returns a slice without the specified elements.
// The given type has to be a comparable (https://pkg.go.dev/golang.org/x/tools/internal/typesinternal?utm_source=gopls#IncomparableMapKey).
func Pull[T any](list []T, removeList ...T) []T {
	m := make(map[any]bool, len(removeList))
	var res []T

	for _, r := range removeList {
		m[r] = true
	}
	for _, l := range list {
		if _, ok := m[l]; !ok {
			res = append(res, l)
		}
	}
	return res
}

// Union returns a slice where all the elements are unique and
// each element in each slice is present at most once in the resulting one.
func Union[T any](list ...[]T) []T {
	m := make(map[any]bool)
	res := []T{}
	for _, l := range list {
		for _, v := range l {
			if _, ok := m[v]; !ok {
				res = append(res, v)
				m[v] = true
			}
		}
	}
	return res
}
