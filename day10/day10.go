package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jamestrew/aoc22/utils"
)

type instruction struct {
	cmd string
	arg int
}

func programLines(input string) <-chan instruction {
	ch := make(chan instruction)
	go func() {
		for _, line := range strings.Split(input, "\n") {
			split := strings.Split(line, " ")
			if len(split) == 1 {
				ch <- instruction{cmd: split[0]}
			} else {
				cmd, arg := split[0], split[1]
				argInt, _ := strconv.Atoi(arg)
				ch <- instruction{cmd, argInt}
			}
		}
		close(ch)
	}()
	return ch
}

func foo(inst instruction) (int, int) {
	var cycleIncr, regIncr int
	switch inst.cmd {
	case "noop":
		cycleIncr = 1
		regIncr = 0
	case "addx":
		cycleIncr = 2
		regIncr = inst.arg
	}
	return cycleIncr, regIncr
}

func part1(input string) int {
	var ret int
	reg := 1
	readCycle := 20
	currCycle := 1
	for instruction := range programLines(input) {
		var cycleIncr, regIncr int
		switch instruction.cmd {
		case "noop":
			cycleIncr = 1
			regIncr = 0
		case "addx":
			cycleIncr = 2
			regIncr = instruction.arg
		}
		for i := 0; i < cycleIncr; i++ {
			if currCycle == readCycle {
				ret += readCycle * reg
				readCycle += 40
			}
			currCycle++
		}
		reg += regIncr
	}

	return ret
}

func part2(input string) []string {
	// pixel location is reg - 1
	// broken down into rows of 40
	// 0-39
	// 40-79
	// 80-119
	// 120...
	reg := 1
	currCycle := 1
	lines := [][]string{}
	for instruction := range programLines(input) {
		cycleIncr, regIncr := foo(instruction)
		for i := 0; i < cycleIncr; i++ {
			spriteLoc := reg - 1
			pixelLoc := currCycle - 1

			row := pixelLoc / 40
			pixelLoc = pixelLoc % 40
			var line []string
			if row >= len(lines) {
				line = make([]string, 40, 40)
				lines = append(lines, line)
			} else {
				line = lines[row]
			}

			if pixelLoc >= spriteLoc && pixelLoc <= spriteLoc+2 {
				line[pixelLoc] = "#"
			} else {
				line[pixelLoc] = "."
			}
			currCycle++
		}
		reg += regIncr
	}

	// fmt.Println(lines)

	var ret []string
	for _, v := range lines {
		line := strings.Join(v, "")
		ret = append(ret, line)
	}
	return ret
}

func Answers() {
	fmt.Println(part1(utils.GetInput(10)))
	ret := part2(utils.GetInput(10))
	fmt.Println(strings.Join(ret, "\n"))
}
