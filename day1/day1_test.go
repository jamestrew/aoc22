package day1

import (
	"testing"

	"github.com/jamestrew/aoc22/utils"
	"github.com/stretchr/testify/assert"
)

const input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 24000, part1(utils.StringScanner(input)))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 45000, part2(utils.StringScanner(input)))
}
