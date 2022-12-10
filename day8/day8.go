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
type Tree struct {
	x, y, height int
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
			if isVisible(trees, Tree{i + 1, j + 1, height}) {
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

func isVisible(trees TreeMap, tree Tree) bool {
	return checkLeft(trees, tree) || checkRight(trees, tree) ||
		checkTop(trees, tree) ||
		checkBottom(trees, tree)
}

func checkLeft(trees TreeMap, tree Tree) bool {
	for j := 0; j < tree.y; j++ {
		if trees[tree.x][j] >= tree.height {
			return false
		}
	}
	return true
}

func checkRight(trees TreeMap, tree Tree) bool {
	for j := len(trees[0]) - 1; j > tree.y; j-- {
		if trees[tree.x][j] >= tree.height {
			return false
		}
	}
	return true
}

func checkTop(trees TreeMap, tree Tree) bool {
	for i := 0; i < tree.x; i++ {
		if trees[i][tree.y] >= tree.height {
			return false
		}
	}
	return true
}

func checkBottom(trees TreeMap, tree Tree) bool {
	for i := len(trees) - 1; i > tree.x; i-- {
		if trees[i][tree.y] >= tree.height {
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
			tree := Tree{i, j, height}
			left := countTreesLeft(trees, tree)
			right := countTreesRight(trees, tree)
			up := countTreesUp(trees, tree)
			down := countTreesDown(trees, tree)
			score = left * right * up * down
			if score > ret {
				ret = score
			}
		}
	}
	return ret
}

func countTreesLeft(trees TreeMap, tree Tree) int {
	count := 0
	for j := tree.y - 1; j >= 0; j-- {
		count++
		if trees[tree.x][j] >= tree.height {
			break
		}
	}
	return count
}

func countTreesRight(trees TreeMap, tree Tree) int {
	count := 0
	for j := tree.y + 1; j <= len(trees[0])-1; j++ {
		count++
		if trees[tree.x][j] >= tree.height {
			break
		}
	}
	return count
}

func countTreesUp(trees TreeMap, tree Tree) int {
	count := 0
	for i := tree.x - 1; i >= 0; i-- {
		count++
		if trees[i][tree.y] >= tree.height {
			break
		}
	}
	return count
}

func countTreesDown(trees TreeMap, tree Tree) int {
	count := 0
	for i := tree.x + 1; i <= len(trees)-1; i++ {
		count++
		if trees[i][tree.y] >= tree.height {
			break
		}
	}
	return count
}

func Answers() {
	fmt.Println(part1(utils.GetInput(8)))
	fmt.Println(part2(utils.GetInput(8)))
}
