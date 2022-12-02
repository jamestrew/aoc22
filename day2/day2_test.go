package day2

import (
	"testing"

	"github.com/jamestrew/aoc22/utils"
	"github.com/stretchr/testify/assert"
)

const input = `A Y
B X
C Z`

func TestPart1(t *testing.T) {
	assert.Equal(t, 15, part1(utils.StringScanner(input)))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 12, part2(utils.StringScanner(input)))
}
