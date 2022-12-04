package day4

import (
	"testing"

	"github.com/jamestrew/aoc22/utils"
	"github.com/stretchr/testify/assert"
)

const input = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

func TestPart1(t *testing.T) {
	assert.Equal(t, 2, part1(utils.StringScanner(input)))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 4, part2(utils.StringScanner(input)))
}
