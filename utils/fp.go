package utils

func Map[T any, S any](fn func(T) S, s []T) []S {
	ret := make([]S, len(s))

	for _, v := range s {
		ret = append(ret, fn(v))
	}
	return ret
}
