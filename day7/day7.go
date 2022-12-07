package day7

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jamestrew/aoc22/utils"
)

var sum int

type Item interface {
	isLsItem()
}

type File struct {
	name string
	size int
}

func (f *File) isLsItem() {}

type Dir struct {
	name     string
	children []Item
	parent   *Dir
	size     int
}

func (d *Dir) isLsItem() {}
func (d *Dir) getSizes() {
	size := 0
	for _, child := range d.children {
		switch child := child.(type) {
		case *File:
			size += child.size
			// fmt.Printf("child file %v -- %v\n", child.name, child.size)
		case *Dir:
			child.getSizes()
			size += child.size
			// fmt.Printf("child dir %v -- %v\n", child.name, child.size)
		}
	}
	d.size = size
}

func buildFs(input *bufio.Scanner) *Dir {
	root := &Dir{name: "/"}
	d := root
	for input.Scan() {
		cmd := input.Text()

		switch {
		case cmd[0:4] == "$ cd":
			d = cdInto(d, cmd)
			// fmt.Println("\ncd'ed to", d.name)
		case cmd == "$ ls":
			// do nothing
			// fmt.Println("$ ls", d.name)
		default:
			d = handlels(d, cmd)
		}
	}
	root.getSizes()
	// fmt.Println(root.size)
	// spew.Dump(root)
	return root
}

func cdInto(currDir *Dir, cmd string) *Dir {
	split := strings.Split(cmd, " ")
	dir := split[2]

	switch dir {
	case "/":
		currDir = goToRoot(currDir)
	case "..":
		getDirSize(currDir)
		currDir = cdBack(currDir)
	default:
		for _, child := range currDir.children {
			childDir, ok := child.(*Dir)
			if !ok {
				continue
			}
			if childDir.name == dir {
				currDir = childDir
			}
		}
	}
	return currDir
}

func goToRoot(currDir *Dir) *Dir {
	for currDir.name != "/" {
		currDir = cdBack(currDir)
	}
	return currDir
}

func cdBack(currDir *Dir) *Dir {
	currDir = currDir.parent
	return currDir
}

var sums []int

func getDirSize(currDir *Dir) {
	for _, child := range currDir.children {
		switch child := child.(type) {
		case *File:
			currDir.size += child.size
		case *Dir:
			currDir.size += child.size
		}
	}
	// if currDir.size <= 100000 {
	// }
	// sum += currDir.size
	// fmt.Println(currDir.name, currDir.size)
	// sums = append(sums, currDir.size)
}

func getDirSize2(currDir *Dir) {
	for _, child := range currDir.children {
		switch child := child.(type) {
		case *File:
			currDir.size += child.size
		case *Dir:
			currDir.size += child.size
			getDirSize(child)
		}
	}
	fmt.Println(currDir.name, currDir.size)
	sums = append(sums, currDir.size)
}

var fileTotal int

func handlels(currDir *Dir, line string) *Dir {
	if line[0] == '$' {
		return currDir
	}
	if currDir.children == nil {
		currDir.children = []Item{}
	}

	split := strings.Split(line, " ")
	first, sec := split[0], split[1]

	if first == "dir" {
		currDir.children = append(currDir.children, &Dir{name: sec, parent: currDir})
		// fmt.Printf("adding %v %v to %v\n", first, sec, currDir.name)
	} else {
		size, _ := strconv.Atoi(first)
		currDir.children = append(currDir.children, &File{name: sec, size: size})
		// fmt.Printf("adding %v %v to %v\n", size, sec, currDir.name)
		fileTotal += size
	}
	return currDir
}

func part1(input *bufio.Scanner) int {
	root := buildFs(input)
	return getSmallSum(root, 0)
}

func getSmallSum(dir *Dir, sum int) int {
	for _, child := range dir.children {
		childDir, ok := child.(*Dir)
		if !ok {
			continue
		}
		if childDir.size <= 100000 {
			sum += childDir.size
			// fmt.Println(sum)
			sum = getSmallSum(childDir, sum)
		}
	}
	return sum
}

const AVAILABLE = 70000000
const NEED = 30000000

// 1272621

func part2(input *bufio.Scanner) int {
	root := buildFs(input)
	foo(root)
	sums = append(sums, root.size)
	sort.Ints(sums)
	fmt.Println(sums)
	ret := 0
	total := fileTotal
	mustClear := NEED - (AVAILABLE - total)

	for _, size := range sums {
		if mustClear-size < 0 {
			return size
		}
	}
	return ret
}

func foo(dir *Dir) {
	for _, child := range dir.children {
		childDir, ok := child.(*Dir)
		if !ok {
			continue
		}
		sums = append(sums, childDir.size)
		foo(childDir)
	}
}

func Answers() {
	// fmt.Println(part1(utils.GetInputScanner(7)))
	// fmt.Println(sum)

	fmt.Println(part2(utils.GetInputScanner(7)))

}
