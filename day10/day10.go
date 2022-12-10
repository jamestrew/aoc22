package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jamestrew/aoc22/utils"
)

type instruction struct {
	cmd       string
	arg       int
	cycleIncr int
	regIncr   int
}

func programLines(input string) <-chan instruction {
	ch := make(chan instruction)
	go func() {
		for _, line := range strings.Split(input, "\n") {
			split := strings.Split(line, " ")
			if len(split) == 1 {
				ch <- instruction{cmd: split[0], cycleIncr: 1, regIncr: 0}
			} else {
				cmd, arg := split[0], split[1]
				argInt, _ := strconv.Atoi(arg)
				ch <- instruction{cmd: cmd, arg: argInt, cycleIncr: 2, regIncr: argInt}
			}
		}
		close(ch)
	}()
	return ch
}

func part1(input string) int {
	var ret int
	reg := 1
	readCycle := 20
	currCycle := 1
	for instruction := range programLines(input) {
		for i := 0; i < instruction.cycleIncr; i++ {
			if currCycle == readCycle {
				ret += readCycle * reg
				readCycle += 40
			}
			currCycle++
		}
		reg += instruction.regIncr
	}
	return ret
}

func part2(input string) []string {
	const WIDTH = 40

	reg := 1
	currCycle := 1
	lines := [][]string{}
	for instruction := range programLines(input) {
		for i := 0; i < instruction.cycleIncr; i++ {
			spriteLoc := reg - 1
			pixelLoc := currCycle - 1

			row := pixelLoc / WIDTH
			pixelLoc = pixelLoc % WIDTH
			var line []string
			if row >= len(lines) {
				line = make([]string, WIDTH, WIDTH)
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
		reg += instruction.regIncr
	}

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
