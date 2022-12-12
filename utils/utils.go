package utils

import (
	"errors"
	"math"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IntDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero zero")
	}
	return int(math.Round(float64(a) / float64(b))), nil
}

func GCD(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func LCM(ints ...int) int {
	if len(ints) < 2 {
		return ints[0]
	}
	a, b, others := ints[0], ints[1], ints[2:]
	ret := a * b / GCD(a, b)
	for i := 0; i < len(others); i++ {
		ret = LCM(ret, others[i])
	}
	return ret
}

type Ordered interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

func Min[T Ordered](slice []T) T {
	min := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func Max[T Ordered](slice []T) T {
	max := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}
