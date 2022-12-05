package day5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jamestrew/aoc22/utils"
)

type Move struct {
	count, from, to int
}

type Moves []Move
type Stack map[int][]string

type CrateCol struct {
	col   int
	crate rune
}

func getMoves(lines string) Moves {
	ret := []Move{}
	lines = strings.TrimSpace(lines)

	pattern := regexp.MustCompile(`move ([0-9]+) from ([0-9]+) to ([0-9]+)`)
	for _, move := range strings.Split(lines, "\n") {
		loc := pattern.FindStringSubmatch(move)
		count, _ := strconv.Atoi(loc[1])
		from, _ := strconv.Atoi(loc[2])
		to, _ := strconv.Atoi(loc[3])
		ret = append(ret, Move{count, from, to})
	}
	return ret
}

func getStack(stacksStr string) Stack {
	stacks := strings.Split(stacksStr, "\n")
	height := len(stacks) - 1

	ret := make(Stack)
	for i := height - 1; i >= 0; i-- {
		stack := stacks[i]
		for i := 1; i < len(stack); i += 4 {
			ch := stack[i]
			if ch == ' ' {
				continue
			}
			col := (i + 5) / 4
			ret[col] = append(ret[col], string(ch))
		}
	}
	return ret
}

func playMove1(move Move, stack Stack) Stack {
	for i := 0; i < move.count; i++ {
		moveHeight := len(stack[move.from]) - 1
		stack[move.to] = append(stack[move.to], stack[move.from][moveHeight])
		stack[move.from] = stack[move.from][0:moveHeight]
	}
	return stack
}

func playMove2(move Move, stack Stack) Stack {
	top := len(stack[move.from])
	moving := stack[move.from][top-move.count : top]
	stack[move.to] = append(stack[move.to], moving...)
	stack[move.from] = stack[move.from][0 : top-move.count]
	return stack
}

func movesAndStacks(input string) (Moves, Stack) {
	input = strings.Trim(input, "\n")
	parts := strings.Split(input, "\n\n")
	return getMoves(parts[1]), getStack(parts[0])
}

func getStackTops(stack Stack) string {
	ret := make([]string, 0, len(stack))
	for i := 1; i <= len(stack); i++ {
		s := stack[i]
		ret = append(ret, s[len(s)-1])
	}
	return strings.Join(ret, "")
}

func part1(input string) string {
	moves, stack := movesAndStacks(input)
	for _, move := range moves {
		stack = playMove1(move, stack)
	}
	return getStackTops(stack)
}

func part2(input string) string {
	moves, stack := movesAndStacks(input)
	for _, move := range moves {
		stack = playMove2(move, stack)
	}
	return getStackTops(stack)
}

func Answers() {
	fmt.Println(part1(utils.GetInput(5)))
	fmt.Println(part2(utils.GetInput(5)))
}
