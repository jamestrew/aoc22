package day8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jamestrew/aoc22/utils"
)

type TreeMap [][]int
type Pos struct {
	x, y int
}

func makeTreeMap(input string) TreeMap {
	rows := strings.Split(input, "\n")

	trees := make([][]int, 0, len(rows))
	for _, row := range rows {
		cols := strings.Split(row, "")

		colInts := make([]int, 0, len(cols))
		for _, col := range cols {
			height, _ := strconv.Atoi(col)
			colInts = append(colInts, height)
		}
		trees = append(trees, colInts)
	}
	return trees
}

func part1(input string) int {
	ret := 0
	trees := makeTreeMap(input)
	for i, row := range trees[1 : len(trees)-1] {
		for j, height := range row[1 : len(trees[0])-1] {
			if isVisible(trees, height, Pos{i + 1, j + 1}) {
				ret++
			}
		}
	}

	ret += countEdges(trees)
	return ret
}

func countEdges(trees TreeMap) int {
	rows, cols := len(trees), len(trees[0])
	return 2*(rows+cols) - 4
}

func isVisible(trees TreeMap, height int, pos Pos) bool {
	return checkLeft(trees, height, pos) || checkRight(trees, height, pos) ||
		checkTop(trees, height, pos) ||
		checkBottom(trees, height, pos)
}

func checkLeft(trees TreeMap, height int, pos Pos) bool {
	for j := 0; j < pos.y; j++ {
		if trees[pos.x][j] >= height {
			return false
		}
	}
	return true
}

func checkRight(trees TreeMap, height int, pos Pos) bool {
	for j := len(trees[0]) - 1; j > pos.y; j-- {
		if trees[pos.x][j] >= height {
			return false
		}
	}
	return true
}

func checkTop(trees TreeMap, height int, pos Pos) bool {
	for i := 0; i < pos.x; i++ {
		if trees[i][pos.y] >= height {
			return false
		}
	}
	return true
}

func checkBottom(trees TreeMap, height int, pos Pos) bool {
	for i := len(trees) - 1; i > pos.x; i-- {
		if trees[i][pos.y] >= height {
			return false
		}
	}
	return true
}

func part2(input string) int {
	ret := 0
	trees := makeTreeMap(input)
	for i, rows := range trees {
		score := 0
		for j, height := range rows {
			pos := Pos{i, j}
			left := lookLeft(trees, height, pos)
			right := lookRight(trees, height, pos)
			up := lookUp(trees, height, pos)
			down := lookDown(trees, height, pos)
			score = left * right * up * down
			if score > ret {
				ret = score
			}
		}
	}
	return ret
}

func lookLeft(trees TreeMap, height int, pos Pos) int {
	count := 0
	for j := pos.y - 1; j >= 0; j-- {
		count++
		if trees[pos.x][j] >= height {
			break
		}
	}
	return count
}

func lookRight(trees TreeMap, height int, pos Pos) int {
	count := 0
	for j := pos.y + 1; j <= len(trees[0])-1; j++ {
		count++
		if trees[pos.x][j] >= height {
			break
		}
	}
	return count
}

func lookUp(trees TreeMap, height int, pos Pos) int {
	count := 0
	for i := pos.x - 1; i >= 0; i-- {
		count++
		if trees[i][pos.y] >= height {
			break
		}
	}
	return count
}

func lookDown(trees TreeMap, height int, pos Pos) int {
	count := 0
	for i := pos.x + 1; i <= len(trees)-1; i++ {
		count++
		if trees[i][pos.y] >= height {
			break
		}
	}
	return count
}

func Answers() {
	fmt.Println(part1(utils.GetInput(8)))
	fmt.Println(part2(utils.GetInput(8)))
}
