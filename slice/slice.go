package slice

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
