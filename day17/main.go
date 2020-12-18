package main

import (
	"fmt"

	"github.com/dwalker109/aoc-2020/day17/pocket"
	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("./input.txt")

	fmt.Println(pt1)
}

func part1(p string) int {
	i := make(chan string)
	go util.StreamInput(i, p)
	dim := parseInput(i)

	for dim.Cycle < 6 {
		dim.Simulate()
	}

	return dim.CountEndActive()
}

func parseInput(i <-chan string) *pocket.Dimension {
	state := make([]string, 0)
	for l := range i {
		state = append(state, l)
	}
	return pocket.NewDimension(state)
}
