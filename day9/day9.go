package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jamestrew/aoc22/utils"
)

type pos struct {
	x, y int
}

type motion struct {
	x, y int
}

type locations struct {
	head, tail pos
}

func getMotions(input string) []motion {
	motions := []motion{}
	splitInput := strings.Split(input, "\n")
	for _, input := range splitInput {
		split := strings.Split(input, " ")
		dir, length := split[0], split[1]
		l, _ := strconv.Atoi(length)
		var m motion
		for i := 0; i < l; i++ {
			switch dir {
			case "R":
				m = motion{1, 0}
			case "L":
				m = motion{-1, 0}
			case "U":
				m = motion{0, 1}
			case "D":
				m = motion{0, -1}
			}
			motions = append(motions, m)
		}
	}

	return motions
}

func tailFollow(loc *locations) (int, int) {
	dx := loc.head.x - loc.tail.x
	dy := loc.head.y - loc.tail.y
	if noMove(loc, dx, dy) {
		return dx, dy
	}

	if dx == 0 || dy == 0 {
		rookMove(loc, dx, dy)
	} else {
		diagonalMove(loc, dx, dy)
	}
	return dx, dy
}

func locationDelta(loc *locations) (int, int) {
	dx := loc.head.x - loc.tail.x
	dy := loc.head.y - loc.tail.y
	return dx, dy
}

func noMove(loc *locations, dx, dy int) bool {
	if (dx <= 1 && dx >= -1) && (dy <= 1 && dy >= -1) {
		return true
	}
	return false
}

func rookMove(loc *locations, dx, dy int) {
	if dx == 0 {
		alignDir(loc, dy, 'y')
	} else {
		alignDir(loc, dx, 'x')
	}
}

func diagonalMove(loc *locations, dx, dy int) {
	alignDir(loc, dx, 'x')
	alignDir(loc, dy, 'y')
}

func alignDir(loc *locations, delta int, dir rune) {
	if delta == 0 {
		return
	}
	unitMagnitude := utils.Abs(delta) / delta

	switch dir {
	case 'x':
		loc.tail.x += unitMagnitude
	case 'y':
		loc.tail.y += unitMagnitude
	}
}

func part1(input string) int {
	visited := utils.NewSet[pos]()
	loc := &locations{head: pos{0, 0}, tail: pos{0, 0}}

	for _, m := range getMotions(input) {
		loc.head.x += m.x
		loc.head.y += m.y
		tailFollow(loc)
		visited.Add(loc.tail)
	}

	return visited.Length()
}

func part2(input string) int {
	motions := getMotions(input)
	for i := 0; i < 8; i++ {
		loc := &locations{head: pos{0, 0}, tail: pos{0, 0}}
		var newMotions []motion
		for _, m := range motions {
			loc.head.x += m.x
			loc.head.y += m.y
			oldTail := loc.tail
			tailFollow(loc)
			newMotions = append(
				newMotions,
				motion{x: loc.tail.x - oldTail.x, y: loc.tail.y - oldTail.y},
			)
		}
		copy(motions, newMotions)
	}

	visited := utils.NewSet[pos]()
	loc := &locations{head: pos{0, 0}, tail: pos{0, 0}}
	for _, m := range motions {
		loc.head.x += m.x
		loc.head.y += m.y
		tailFollow(loc)
		visited.Add(loc.tail)
	}

	return visited.Length()
}

func Answers() {
	fmt.Println(part1(utils.GetInput(9)))
	fmt.Println(part2(utils.GetInput(9)))
}
