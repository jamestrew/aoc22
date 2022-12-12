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
type Pos struct {
	x, y int
}

// fewest number of steps to get from S to E
func part1(input string) int {
	_map, start, end := parseInput(input)
	return shortestPath(_map, start, end)
}

func part2(input string) int {
	_map, _, end := parseInput(input)

	starts := []Pos{}
	for i, row := range _map {
		for j, elv := range row {
			if elv == 0 {
				starts = append(starts, Pos{j, i})
			}
		}
	}

	pathLens := []int{}
	for _, start := range starts {
		distance := shortestPath(_map, start, end)
		if distance != -1 {
			pathLens = append(pathLens, shortestPath(_map, start, end))
		}
	}

	return utils.Min(pathLens)
}

func shortestPath(_map Map, start, end Pos) int {
	queue := &utils.Queue[Pos]{}
	queue.Enqueue(start)
	visited := make(map[Pos]bool)
	visited[start] = true
	distance := make(map[Pos]int)
	distance[start] = 0

	for queue.Size != 0 {
		pos := queue.Dequeue()
		if pos == end {
			return distance[end]
		}

		for neighbor := range getNeighbors(_map, pos) {
			currElv, neighborElv := _map[pos.y][pos.x], _map[neighbor.y][neighbor.x]
			_, ok := visited[neighbor]
			if currElv+1 >= neighborElv && !ok {
				queue.Enqueue(neighbor)
				visited[neighbor] = true
				distance[neighbor] = distance[pos] + 1
			}
		}
	}
	return -1
}

func parseInput(input string) (Map, Pos, Pos) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	_map := make(Map, 0, len(lines))
	start := Pos{}
	end := Pos{}
	for i, line := range lines {
		row := make([]int, 0, len(line))
		for j, heightLetter := range strings.Split(line, "") {
			heightRune := rune(heightLetter[0])
			switch heightRune {
			case 'S':
				start.x, start.y = j, i
				heightRune = 'a'
			case 'E':
				end.x, end.y = j, i
				heightRune = 'z'
			}
			row = append(row, int(heightRune-'a'))
		}
		_map = append(_map, row)
	}
	return _map, start, end
}

func getNeighbors(_map Map, pos Pos) <-chan Pos {
	ret := make(chan Pos)

	go func() {
		yMax, xMax := len(_map), len(_map[0])
		if xMax == 0 || yMax == 0 {
			close(ret)
			return
		}

		dx, dy := [4]int{1, -1, 0, 0}, [4]int{0, 0, 1, -1}
		for i := 0; i < 4; i++ {
			newx, newy := pos.x+dx[i], pos.y+dy[i]
			if newx >= 0 && newx < xMax && newy >= 0 && newy < yMax {
				ret <- Pos{newx, newy}
			}
		}
		close(ret)
	}()
	return ret
}

func Answers() {
	fmt.Println(part1(utils.GetInput(12)))
	fmt.Println(part2(utils.GetInput(12)))
}
