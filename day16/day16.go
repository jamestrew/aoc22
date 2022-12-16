package day16

import (
	"fmt"
	"strings"

	"github.com/jamestrew/aoc22/utils"
)

type Valve struct {
	name     string
	flowrate int
	on       bool
	paths    []*Valve
}

func newValve(name string) *Valve {
	return &Valve{name: name, on: true, paths: make([]*Valve, 0)}
}

func (v *Valve) String() string {
	paths := []string{}
	for _, path := range v.paths {
		paths = append(paths, path.name)
	}
	return fmt.Sprintf("{%s %d [%s]}", v.name, v.flowrate, strings.Join(paths, " "))
}

func part1(input string) int {
	ret := 0
	valves := parseInput(input)
	for _, v := range valves {
		utils.NewQueue(v)

		// for queue.Size != 0 {
		// 	currv := queue.Dequeue()
		// }
	}
	return ret
}

func parseInput(input string) []*Valve {
	input = strings.TrimSpace(input)
	valves := make(map[string]*Valve)

	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, ";")
		first := strings.Fields(split[0])
		name := first[1]
		valve, ok := valves[name]
		if !ok {
			valve = newValve(name)
		}
		valve.flowrate = utils.Atoi(strings.Split(first[4], "=")[1])

		paths := strings.Fields(split[1])[4:]
		for _, path := range paths {
			path = strings.ReplaceAll(path, ",", "")
			v, ok := valves[path]
			if !ok {
				v = newValve(path)
				valves[path] = v
			}
			valve.paths = append(valve.paths, v)
		}
		valves[valve.name] = valve
	}

	ret := make([]*Valve, 0, len(valves))
	for _, v := range valves {
		ret = append(ret, v)
	}
	return ret
}

func resetValves(valves []*Valve) {
	for _, v := range valves {
		v.on = false
	}
}

func part2(input string) int {
	ret := 0
	return ret
}

func Answers() {
	fmt.Println(part1(utils.GetInput(16)))
	fmt.Println(part2(utils.GetInput(16)))
}
