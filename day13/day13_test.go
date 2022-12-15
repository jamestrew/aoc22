package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`

// func TestPart1Example(t *testing.T) {
// 	assert.Equal(t, 13, part1(input))
// }

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 0, part2(input))
}

func TestParser(t *testing.T) {
	tests := []struct {
		input    string
		expected *List
	}{
		{"[]", &List{}},
		{"[[]]", &List{[]Expr{&List{}}}},
		{"[1]", &List{[]Expr{&Int{1}}}},
		{"[[1]]", &List{[]Expr{&List{[]Expr{&Int{1}}}}}},
		{"[1,2]", &List{[]Expr{&Int{1}, &Int{2}}}},
		{"[1,2,3]", &List{[]Expr{&Int{1}, &Int{2}, &Int{3}}}},
		{"[1,[2]]", &List{[]Expr{&Int{1}, &List{[]Expr{&Int{2}}}}}},
		{
			"[1,[[2],3]]",
			&List{[]Expr{&Int{1}, &List{[]Expr{&List{[]Expr{&Int{2}}}, &Int{3}}}}},
		},
		{
			"[1,[[2],3],4]",
			&List{[]Expr{&Int{1}, &List{[]Expr{&List{[]Expr{&Int{2}}}, &Int{3}}}, &Int{4}}},
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expected, parse(tc.input))
	}
}

func TestExprString(t *testing.T) {
	tests := []string{
		"[]",
		"[[]]",
		"[1]",
		"[[1]]",
		"[1,2]",
		"[1,2,3]",
		"[1,[2]]",
		"[[1],[2,3,4]]",
		"[1,[[2],3]]",
		"[1,[[2],3]]",
		"[1,[[2],3],4]", // [1,[[2],3,4]]
		"[1,[2,[3,[4,[5,6,7]]]],8,9]",
		// "[[[],6,9,[2,[10,4,6,9,1]]],[[[8,6,4,5],1,[5,5,3,7,8],[]]],[[[5,5,10,7]],[[2,4],0,[2,2],10,[]],[[1],[0],[1,1,7,3,8],[4,10,0,0]]]]",
	}

	for _, tc := range tests {
		expr := parse(tc)
		assert.Equal(t, tc, expr.String())
	}
}