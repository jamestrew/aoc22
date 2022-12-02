package day2

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/jamestrew/aoc22/utils"
)

// A X - rock
// B Y - paper
// C Z - scissor

type moves uint
type outcome uint

const (
	ROCK moves = iota
	PAPER
	SCISSOR
)

const (
	WIN outcome = iota
	LOSE
	TIE
)

var ELF_MOVES = map[string]moves{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSOR,
}

var MY_MOVES = map[string]moves{
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSOR,
}

var WANTED_OUTCOME = map[string]outcome{
	"X": LOSE,
	"Y": TIE,
	"Z": WIN,
}

var MOVE_POINTS = map[moves]int{
	ROCK:    1,
	PAPER:   2,
	SCISSOR: 3,
}

var OUTCOME = map[outcome]int{
	WIN:  6,
	LOSE: 0,
	TIE:  3,
}

var RULES = [][]outcome{
	{TIE, LOSE, WIN},
	{WIN, TIE, LOSE},
	{LOSE, WIN, TIE},
}

var PLAY_MOVES = [][]moves{
	{PAPER, SCISSOR, ROCK},
	{SCISSOR, ROCK, PAPER},
	{ROCK, PAPER, SCISSOR},
}

func part1(input *bufio.Scanner) int {
	var score int

	for input.Scan() {
		moves := strings.Split(input.Text(), " ")
		elf, me := ELF_MOVES[moves[0]], MY_MOVES[moves[1]]

		outcome := RULES[me][elf]
		score += OUTCOME[outcome] + MOVE_POINTS[me]
	}
	return score
}

func part2(input *bufio.Scanner) int {
	var score int

	for input.Scan() {
		strat := strings.Split(input.Text(), " ")
		elf, result := ELF_MOVES[strat[0]], WANTED_OUTCOME[strat[1]]

		myMove := PLAY_MOVES[elf][result]
		score += OUTCOME[result] + MOVE_POINTS[myMove]
	}
	return score
}

func Answers() {
	fmt.Println(part1(utils.GetInputScanner(2)))
	fmt.Println(part2(utils.GetInputScanner(2)))
}
