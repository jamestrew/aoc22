package day13

import (
	"fmt"
	"strings"

	"github.com/jamestrew/aoc22/utils"
)

// packets out of order
// pairs of packets signified by blank line
// packet is always a list, one per line

// (left,right)
// if both ints -> low,high
// if both list -> compare ints -> small len, high len - compare values
// if one int   -> convert int to list & retry comparison

type comp int

const (
	ok comp = iota
	pass
	bad
)

type PacketPair struct {
	left, right *List
}

func part1(input string) int {
	ret := 0
	packets := parseInput1(input)
	for idx, pair := range packets {
		if c := compare(pair.left, pair.right); c == ok {
			ret += idx + 1
		}
	}
	return ret
}

func parseInput1(input string) []PacketPair {
	ret := []PacketPair{}
	input = strings.TrimSpace(input)

	for _, pair := range strings.Split(input, "\n\n") {
		packets := strings.Split(pair, "\n")
		p := PacketPair{
			left:  parseList(packets[0]),
			right: parseList(packets[1]),
		}
		ret = append(ret, p)
	}
	return ret
}

func compare(left, right *List) comp {
	minLength := utils.Min(left.Length(), right.Length())

	for i := 0; i < minLength; i++ {
		l, r := left.elements[i], right.elements[i]
		var c comp
		if c = intOrdered(l, r); c != pass {
			return c
		}
		if c = listOrdered(l, r); c != pass {
			return c
		}
		if c = asymmetricOrdered(l, r); c != pass {
			return c
		}
	}

	if left.Length() < right.Length() {
		return ok
	} else if left.Length() > right.Length() {
		return bad
	}

	return pass
}

func intOrdered(left, right Expr) comp {
	l, lOk := left.(*Int)
	r, rOk := right.(*Int)

	if (!lOk || !rOk) || l.val == r.val {
		return pass
	}

	if l.val < r.val {
		return ok
	}
	return bad
}

func listOrdered(left, right Expr) comp {
	l, lOk := left.(*List)
	r, rOk := right.(*List)

	if !lOk || !rOk {
		return pass
	}
	return compare(l, r)
}

func asymmetricOrdered(left, right Expr) comp {
	lInt, lOk := left.(*Int)
	rList, rOk := right.(*List)
	if lOk && rOk {
		lList := &List{}
		lList.elements = append(lList.elements, lInt)
		return compare(lList, rList)
	}

	lList, lOk := left.(*List)
	rInt, rOk := right.(*Int)
	if lOk && rOk {
		rList := &List{}
		rList.elements = append(rList.elements, rInt)
		return compare(lList, rList)
	}
	return pass
}

var MARK_1 = &List{[]Expr{&List{[]Expr{&Int{2}}}}}
var MARK_2 = &List{[]Expr{&List{[]Expr{&Int{6}}}}}

func part2(input string) int {
	ret := 1
	packets := parseInput2(input)
	compFn := func(a, b *List) bool {
		if c := compare(a, b); c == ok {
			return false
		}
		return true
	}
	utils.Sort(packets, compFn)

	for idx, list := range packets {
		if list == MARK_1 || list == MARK_2 {
			ret *= (idx + 1)
		}
	}
	return ret
}

func parseInput2(input string) []*List {
	ret := []*List{MARK_1, MARK_2}

	input = strings.TrimSpace(input)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		ret = append(ret, parseList(line))
	}

	return ret
}

func Answers() {
	fmt.Println(part1(utils.GetInput(13))) // 5316 -> too low, 5478 -> too low
	fmt.Println(part2(utils.GetInput(13)))
}
