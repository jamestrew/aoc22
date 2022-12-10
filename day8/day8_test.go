package day8

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `
30373
25512
65332
33549
35390
`

func TestPart1(t *testing.T) {
	input := strings.TrimSpace(input)
	assert.Equal(t, 21, part1(input))
}

func TestPart2(t *testing.T) {
	input := strings.TrimSpace(input)
	assert.Equal(t, 8, part2(input))
}

func TestPart2LookLeft(t *testing.T) {
	input := strings.TrimSpace(input)
	trees := makeTreeMap(input)

	tests := []struct {
		height, i, j, expected int
	}{
		{5, 3, 2, 2},
		{5, 1, 2, 1},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expected, countTreesLeft(trees, Tree{tc.i, tc.j, tc.height}))
	}
}

func TestPart2LookRight(t *testing.T) {
	input := strings.TrimSpace(input)
	trees := makeTreeMap(input)

	tests := []struct {
		height, i, j, expected int
	}{
		{5, 3, 2, 2},
		{5, 1, 2, 2},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expected, countTreesRight(trees, Tree{tc.i, tc.j, tc.height}))
	}
}

func TestPart2LookUp(t *testing.T) {
	input := strings.TrimSpace(input)
	trees := makeTreeMap(input)

	tests := []struct {
		height, i, j, expected int
	}{
		{5, 3, 2, 2},
		{5, 1, 2, 1},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expected, countTreesUp(trees, Tree{tc.i, tc.j, tc.height}))
	}
}

func TestPart2LookDown(t *testing.T) {
	input := strings.TrimSpace(input)
	trees := makeTreeMap(input)

	tests := []struct {
		height, i, j, expected int
	}{
		{5, 3, 2, 1},
		{5, 1, 2, 2},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expected, countTreesDown(trees, Tree{tc.i, tc.j, tc.height}))
	}
}
