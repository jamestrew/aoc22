package day4

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/jamestrew/aoc22/utils"
)

type Section struct {
	start, end int
}

func NewSection(input string) Section {
	split := strings.Split(input, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])
	return Section{start, end}
}

func part1(input *bufio.Scanner) int {
	ret := 0
	for input.Scan() {
		pair := strings.Split(input.Text(), ",")
		first, second := NewSection(pair[0]), NewSection(pair[1])

		if (first.start <= second.start && first.end >= second.end) ||
			(first.start >= second.start && first.end <= second.end) {
			ret++
		}
	}
	return ret
}

func part2(input *bufio.Scanner) int {
	ret := 0
	for input.Scan() {
		pair := strings.Split(input.Text(), ",")
		first, second := NewSection(pair[0]), NewSection(pair[1])

		if (first.start >= second.start && first.start <= second.end) ||
			(first.end >= second.start && first.end <= second.end) ||
			(second.start >= first.start && second.start <= first.end) ||
			(second.end >= first.start && second.end <= first.end) {
			ret++
		}
	}
	return ret
}

func Answers() {
	fmt.Println(part1(utils.GetInputScanner(4)))
	fmt.Println(part2(utils.GetInputScanner(4)))
}
