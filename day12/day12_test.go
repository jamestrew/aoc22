package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 31, part1(input))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 29, part2(input))
}
