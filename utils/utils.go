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

// func LCM(a, b int, others ...int) int {
// 	ret := a * b / GCD(a, b)
// 	for i := 0; i < len(others); i++ {
// 		ret = LCM(ret, others[i])
// 	}
// 	return ret
// }

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
