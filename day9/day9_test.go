package day9

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

func TestPart1(t *testing.T) {
	input := strings.TrimSpace(input)
	assert.Equal(t, 13, part1(input))
}

const input2 = `
R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`

func TestPart2(t *testing.T) {
	input := strings.TrimSpace(input)
	assert.Equal(t, 1, part2(input))

	input = strings.TrimSpace(input2)
	assert.Equal(t, 36, part2(input))
}

func TestGetMotions(t *testing.T) {
	input := strings.TrimSpace(input)
	expected := []motion{
		{1, 0}, // R 4
		{1, 0},
		{1, 0},
		{1, 0},
		{0, 1}, // U 4
		{0, 1},
		{0, 1},
		{0, 1},
		{-1, 0}, // L 3
		{-1, 0},
		{-1, 0},
		{0, -1}, // D 1
		{1, 0},  // R 4
		{1, 0},
		{1, 0},
		{1, 0},
		{0, -1}, // D 1
		{-1, 0}, // L 5
		{-1, 0},
		{-1, 0},
		{-1, 0},
		{-1, 0},
		{1, 0}, // R 2
		{1, 0},
	}

	motions := getMotions(input)
	assert.Equal(t, len(expected), len(motions))
	for i, m := range motions {
		assert.Equal(t, expected[i], m)
	}
}

func TestTailFollow(t *testing.T) {
	tests := []struct {
		loc      locations
		expected pos
	}{
		// touching
		{locations{head: pos{0, 0}, tail: pos{0, 0}}, pos{0, 0}},
		{locations{head: pos{0, 0}, tail: pos{1, 0}}, pos{1, 0}},
		{locations{head: pos{0, 0}, tail: pos{0, 1}}, pos{0, 1}},
		{locations{head: pos{0, 0}, tail: pos{1, 1}}, pos{1, 1}},
		{locations{head: pos{0, 0}, tail: pos{-1, 0}}, pos{-1, 0}},
		{locations{head: pos{0, 0}, tail: pos{0, -1}}, pos{0, -1}},
		{locations{head: pos{0, 0}, tail: pos{-1, -1}}, pos{-1, -1}},

		// rook moves
		{locations{head: pos{0, 0}, tail: pos{2, 0}}, pos{1, 0}},
		{locations{head: pos{0, 0}, tail: pos{0, 2}}, pos{0, 1}},
		{locations{head: pos{0, 0}, tail: pos{-2, 0}}, pos{-1, 0}},
		{locations{head: pos{0, 0}, tail: pos{0, -2}}, pos{0, -1}},

		// biship moves
		{locations{head: pos{0, 0}, tail: pos{-2, -1}}, pos{-1, 0}},
		{locations{head: pos{0, 0}, tail: pos{-1, -2}}, pos{0, -1}},
	}

	for _, tc := range tests {
		dx, dy := tailFollow(&tc.loc)
		assert.Equal(t, tc.expected, tc.loc.tail, fmt.Sprintf("(%v,%v)", dx, dy))
	}
}
