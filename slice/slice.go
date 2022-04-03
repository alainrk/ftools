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
