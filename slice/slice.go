package slice

func Filter[T any](list []T, f func(T) bool) ([]T, error) {
	var result []T
	for _, v := range list {
		if f(v) {
			result = append(result, v)
		}
	}
	return result, nil
}
