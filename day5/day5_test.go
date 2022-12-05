package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestPart1(t *testing.T) {
	assert.Equal(t, "CMZ", part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "MCD", part2(input))
}

/*
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3
0123456789

       [D]
   [N] [C]
   [Z] [M] [P]
    1   2   3
0123456789


first guess: BQGTHFZHV
*/
