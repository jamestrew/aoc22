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
	Compare(Expr) comp
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

func (l *List) Compare(right Expr) comp {
	switch r := right.(type) {
	case *List:
		minLength := utils.Min(l.Length(), r.Length())

		for i := 0; i < minLength; i++ {
			left, right := l.elements[i], r.elements[i]
			if c := left.Compare(right); c != pass {
				return c
			}
		}

		delta := l.Length() - r.Length()
		switch {
		case delta < 0:
			return ok
		case delta > 0:
			return bad
		default:
			return pass
		}

	case *Int:
		right := &List{}
		right.elements = append(right.elements, r)
		return l.Compare(right)
	}
	panic(fmt.Sprintf("list compare shouldn't be here: %v-%v", l, right))
}

type Int struct {
	val int
}

func (i *Int) Expression()    {}
func (i *Int) String() string { return fmt.Sprintf("%d", i.val) }
func (i *Int) Compare(right Expr) comp {
	switch r := right.(type) {
	case *Int:
		switch {
		case i.val > r.val:
			return bad
		case i.val < r.val:
			return ok
		default:
			return pass
		}
	case *List:
		left := &List{}
		left.elements = append(left.elements, i)
		return left.Compare(r)
	}
	panic(fmt.Sprintf("int compare shouldn't be here: %v-%v", i, right))
}

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
