package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jamestrew/aoc22/utils"
)

type MonkAttr struct {
	number    int
	items     *utils.Queue[int]
	operation struct {
		right    string
		operator string
	}
	testDivisor int
	ifTrue      int
	ifFalse     int
	inspected   int
}

// before a monkey tests worry level, worry level is divided by 3 (rounded to nearest int)

func parseInput(input string) map[int]*MonkAttr {
	input = strings.TrimSpace(input)

	ret := map[int]*MonkAttr{}
	for monkeyNum, monkey := range strings.Split(input, "\n\n") {
		monkAttr := &MonkAttr{number: monkeyNum, inspected: 0}
		split := strings.Split(monkey, "\n")

		itemLine := strings.Split(split[1], ":")
		itemStr := strings.Split(itemLine[1], ", ")
		monkAttr.items = utils.QueueFromSlice(utils.MapStrInt(itemStr))

		opLine := strings.Split(split[2], ":")
		opStr := strings.Split(strings.TrimSpace(opLine[1]), " ")
		monkAttr.operation.right = opStr[4]
		monkAttr.operation.operator = opStr[3]

		testLine := strings.Split(split[3], ":")
		testStr := strings.Split(strings.TrimSpace(testLine[1]), " ")[2]
		testDivisor, _ := strconv.Atoi(testStr)
		monkAttr.testDivisor = testDivisor

		trueLine := strings.Split(split[4], ":")
		trueStr := strings.Split(strings.TrimSpace(trueLine[1]), " ")[3]
		trueInt, _ := strconv.Atoi(trueStr)
		monkAttr.ifTrue = trueInt

		falseLine := strings.Split(split[5], ":")
		falseStr := strings.Split(strings.TrimSpace(falseLine[1]), " ")[3]
		falseInt, _ := strconv.Atoi(falseStr)
		monkAttr.ifFalse = falseInt

		ret[monkAttr.number] = monkAttr
	}

	return ret
}

func (m *MonkAttr) doOperation(wLevel int) int {
	var rInt int
	if m.operation.right == "old" {
		rInt = wLevel
	} else {
		i, _ := strconv.Atoi(m.operation.right)
		rInt = int(i)
	}

	var ret int
	switch m.operation.operator {
	case "+":
		ret = wLevel + rInt
	case "-":
		ret = wLevel - rInt
	case "*":
		ret = wLevel * rInt
	case "/":
		ret = wLevel / rInt
	}
	return ret
}

func getRelieved(item int) int {
	return item / 3
}

func (m *MonkAttr) testItem(item int) int {
	if item%int(m.testDivisor) == 0 {
		return m.ifTrue
	}
	return m.ifFalse
}

func (m *MonkAttr) throwCurrentItem() {
	m.items.Dequeue()
}

// func (m *MonkAttr) receiveItem(item int) {
// 	m.items.Enqueue(item)
// }

// total number of times a monkey inspected an item over 20 rounds (top 2)
// multiplied together
func part1(input string) int {
	monkeys := parseInput(input)
	for i := 0; i < 20; i++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := monkeys[i]
			itemSlice := monkey.items.ToSlice()
			for _, oldLvl := range itemSlice {
				midLvl := monkey.doOperation(oldLvl)
				monkey.inspected++
				newLvl := getRelieved(midLvl)
				throwTo := monkey.testItem(newLvl)
				monkey.items.Dequeue()
				monkeys[throwTo].items.Enqueue(newLvl)
			}
		}
	}

	counts := []int{}
	for i := 0; i < len(monkeys); i++ {
		m := monkeys[i]
		// m.items.DebugPrint()
		counts = append(counts, m.inspected)
	}
	fmt.Println(counts)
	sort.Ints(counts)

	return counts[len(counts)-1] * counts[len(counts)-2]
}

func part2(input string) int {
	monkeys := parseInput(input)

	divisors := []int{}
	for _, mon := range monkeys {
		divisors = append(divisors, mon.testDivisor)
	}
	mod := utils.LCM(divisors...)

	for i := 0; i < 10000; i++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := monkeys[i]
			itemSlice := monkey.items.ToSlice()
			for _, oldLvl := range itemSlice {
				newLvl := monkey.doOperation(oldLvl) % mod
				monkey.inspected++
				throwTo := monkey.testItem(newLvl)
				monkey.items.Dequeue()
				monkeys[throwTo].items.Enqueue(newLvl)
			}
		}
	}

	counts := []int{}
	for i := 0; i < len(monkeys); i++ {
		m := monkeys[i]
		// m.items.DebugPrint()
		counts = append(counts, m.inspected)
	}
	sort.Ints(counts)
	fmt.Println(counts)

	return counts[len(counts)-1] * counts[len(counts)-2]
}

func Answers() {
	fmt.Println(part1(utils.GetInput(11)))
	fmt.Println(part2(utils.GetInput(11)))
}
