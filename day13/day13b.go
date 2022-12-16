package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jamestrew/aoc22/utils"
)

type Expr interface {
	Expression()
	String() string
}

type List struct {
	elements []Expr
}

func (l *List) Expression() {}
func (l *List) Length() int {
	if l.elements == nil {
		return 0
	}
	return len(l.elements)
}

func (l *List) String() string {
	var out strings.Builder
	out.WriteString("[")

	elems := make([]string, 0, len(l.elements))
	for _, elem := range l.elements {
		elems = append(elems, elem.String())
	}
	out.WriteString(strings.Join(elems, ","))
	out.WriteString("]")
	return out.String()
}

type Int struct {
	val int
}

func (i *Int) Expression()    {}
func (i *Int) String() string { return fmt.Sprintf("%d", i.val) }

func parseList(input string) *List {
	input = strings.TrimSpace(input)
	input = input[1:]

	root := &List{}
	curr := root
	parent := utils.Stack[*List]{}
	parent.Push(root)

	var num strings.Builder
	for _, ch := range input {
		switch ch {
		case '[':
			list := &List{}
			curr.elements = append(curr.elements, list)
			parent.Push(curr)
			curr = list
		case ']':
			if val, ok := makeNum(num); ok {
				curr.elements = append(curr.elements, &Int{val})
				num.Reset()
			}
			curr = parent.Pop()
		case ',':
			if val, ok := makeNum(num); ok {
				curr.elements = append(curr.elements, &Int{val})
				num.Reset()
			}
		default:
			fmt.Fprintf(&num, "%v", ch-'0')
		}
	}
	return root
}

func makeNum(sb strings.Builder) (int, bool) {
	if sb.Len() <= 0 {
		return 0, false
	}

	i, _ := strconv.Atoi(sb.String())
	return i, true
}
