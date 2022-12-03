package day3

import (
	"testing"

	"github.com/jamestrew/aoc22/utils"
	"github.com/stretchr/testify/assert"
)

const input = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestPart1(t *testing.T) {
	assert.Equal(t, 157, part1(utils.StringScanner(input)))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 70, part2(utils.StringScanner(input)))
}
