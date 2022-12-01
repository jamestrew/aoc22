package day1

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"

	"github.com/jamestrew/aoc22/util"
)

func part1(scanner *bufio.Scanner) int {
	calorieSums := calorieSums(scanner)

	max := 0
	for _, cal := range calorieSums {
		if cal > max {
			max = cal
		}
	}
	return max
}

func part2(scanner *bufio.Scanner) int {
	calorieSums := calorieSums(scanner)
	sort.Ints(calorieSums)

	ret := 0
	for i := len(calorieSums) - 1; i > len(calorieSums)-4; i-- {
		ret += calorieSums[i]
	}
	return ret
}

func calorieSums(scanner *bufio.Scanner) []int {
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
	return calorieSums
}

func Answers() {
	fmt.Println(part1(util.FileScanner("./day1/input1")))
	fmt.Println(part2(util.FileScanner("./day1/input1")))
}
