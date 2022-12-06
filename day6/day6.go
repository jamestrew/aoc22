package day6

import (
	"fmt"

	"github.com/jamestrew/aoc22/utils"
)

func what(input string, length int) int {
	ret := 0

	for idx := range input[:len(input)-length+1] {
		hashmap := make(map[byte]bool)
		for i := idx; i < idx+length; i++ {
			_, ok := hashmap[input[i]]
			if ok {
				break
			}
			hashmap[input[i]] = true
			ret = i + 1
		}
		if len(hashmap) == length {
			break
		}
	}
	return ret
}

func part1(input string) int {
	return what(input, 4)
}

func part2(input string) int {
	return what(input, 14)
}

func Answers() {
	fmt.Println(part1(utils.GetInput(6)))
	fmt.Println(part2(utils.GetInput(6)))
}
