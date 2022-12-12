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
	if len(slice) == 0 {
		panic("can't find Min of a zero length slice")
	}
	min := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func Max[T Ordered](slice []T) T {
	if len(slice) == 0 {
		panic("can't find Max of a zero length slice")
	}
	max := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

type Pos struct {
	X, Y int
}

var dirs = []Pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func GetNeighbors[T any](matrix [][]T, pos Pos) <-chan Pos {
	ret := make(chan Pos)

	go func() {
		yMax, xMax := len(matrix), len(matrix[0])
		if xMax == 0 || yMax == 0 {
			close(ret)
			return
		}

		for _, dir := range dirs {
			newx, newy := pos.X+dir.X, pos.Y+dir.Y
			if newx >= 0 && newx < xMax && newy >= 0 && newy < yMax {
				ret <- Pos{newx, newy}
			}
		}
		close(ret)
	}()
	return ret
}
