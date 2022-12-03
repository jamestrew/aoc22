package day3

import (
	"bufio"
	"fmt"

	"github.com/jamestrew/aoc22/utils"
)

type groupBags struct {
	first, second, third string
}

func part1(input *bufio.Scanner) int {
	priority := 0
	for input.Scan() {
		rucksack := input.Text()
		left, right := rucksack[0:len(rucksack)/2], rucksack[len(rucksack)/2:]
		common := commonItems(left, right)
		priority += itemPoint(rune(common[0]))
	}
	return priority
}

func part2(input *bufio.Scanner) int {
	priority := 0
	for bags := range getGroups(input) {
		common := commonItems(bags.first, bags.second)
		common = commonItems(common, bags.third)
		priority += itemPoint(rune(common[0]))
	}
	return priority
}

func getGroups(input *bufio.Scanner) chan groupBags {
	bags := make(chan groupBags)

	go func() {
		for input.Scan() {
			first := input.Text()
			input.Scan()
			second := input.Text()
			input.Scan()
			third := input.Text()
			bags <- groupBags{first, second, third}
		}
		close(bags)
	}()
	return bags
}

func commonItems(first, second string) string {
	hashmap := make(map[rune]bool)
	ret := []rune{}
	for _, ch := range first {
		hashmap[ch] = true
	}

	for _, ch := range second {
		_, ok := hashmap[ch]
		if ok {
			ret = append(ret, ch)
			delete(hashmap, ch)
		}
	}
	return string(ret)
}

func itemPoint(ch rune) int {
	if ch >= 65 && ch <= 90 {
		return int(ch-'A') + 27
	}
	return int(ch-'a') + 1
}

func Answers() {
	fmt.Println(part1(utils.GetInputScanner(3)))
	fmt.Println(part2(utils.GetInputScanner(3)))
}
