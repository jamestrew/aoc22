package utils

import "strconv"

func Map[T any, S any](fn func(T) S, slice []T) []S {
	ret := make([]S, 0, len(slice))

	for _, v := range slice {
		ret = append(ret, fn(v))
	}
	return ret
}

func MapStrInt(slice []string) []int {
	fn := func(x string) int {
		i, _ := strconv.Atoi(x)
		return i
	}

	return Map(fn, slice)
}

func Filter[T any](fn func(T) bool, slice []T) []T {
	ret := make([]T, 0, len(slice))

	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
