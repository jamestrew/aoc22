package day1

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/jamestrew/aoc22/util"
)

func part1(scanner *bufio.Scanner) int {
	calorieSums := []int{}
	sum := 0
	for scanner.Scan() {
		val, ok := strconv.Atoi(scanner.Text())
		if ok == nil {
			sum += val
		} else {
			calorieSums = append(calorieSums, sum)
			sum = 0
		}
	}
	calorieSums = append(calorieSums, sum)

	max := 0
	for _, cal := range calorieSums {
		if cal > max {
			max = cal
		}
	}
	return max
}

func Answers() {
	fmt.Println(part1(util.FileScanner("./day1/input1")))
}
