package day12

import (
	"fmt"
	"strings"

	"github.com/jamestrew/aoc22/utils"
)

// current position S as elevation a
// want to get to position E with elevation z
// a-z elevation lowest to highest

// take one step at a time
// max one level higher
// can be much lower

type Map [][]int

// fewest number of steps to get from S to E
func part1(input string) int {
	_map, start, end := parseInput(input)
	return shortestPath(_map, []utils.Pos{start}, end)
}

func part2(input string) int {
	_map, _, end := parseInput(input)
	starts := []utils.Pos{}
	for i, row := range _map {
		for j, elv := range row {
			if elv == 0 {
				starts = append(starts, utils.Pos{X: j, Y: i})
			}
		}
	}
	return shortestPath(_map, starts, end)
}

func parseInput(input string) (Map, utils.Pos, utils.Pos) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	_map := make(Map, 0, len(lines))
	start := utils.Pos{}
	end := utils.Pos{}
	for i, line := range lines {
		row := make([]int, 0, len(line))
		for j, heightLetter := range strings.Split(line, "") {
			heightRune := rune(heightLetter[0])
			switch heightRune {
			case 'S':
				start.X, start.Y = j, i
				heightRune = 'a'
			case 'E':
				end.X, end.Y = j, i
				heightRune = 'z'
			}
			row = append(row, int(heightRune-'a'))
		}
		_map = append(_map, row)
	}
	return _map, start, end
}

func shortestPath(_map Map, starts []utils.Pos, end utils.Pos) int {
	queue := utils.QueueFromSlice(starts)

	dist := utils.NewDefaultDict[utils.Pos](10000)
	for _, pos := range starts {
		dist.Set(pos, 0)
	}

	for queue.Size != 0 {
		pos := queue.Dequeue()
		if pos == end {
			return dist.Get(end)
		}

		for neighbor := range utils.GetNeighbors(_map, pos) {
			currElv, neighborElv := _map[pos.Y][pos.X], _map[neighbor.Y][neighbor.X]
			if currElv+1 >= neighborElv {
				newDist := dist.Get(pos) + 1
				if newDist < dist.Get(neighbor) {
					queue.Enqueue(neighbor)
					dist.Set(neighbor, newDist)
				}
			}
		}
	}
	return -1
}

func Answers() {
	fmt.Println(part1(utils.GetInput(12)))
	fmt.Println(part2(utils.GetInput(12)))
}
