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
