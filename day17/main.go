package main

import (
	"fmt"

	"github.com/dwalker109/aoc-2020/day17/dim3"
	"github.com/dwalker109/aoc-2020/day17/dim4"
	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("./input.txt")
	pt2 := part2("./input.txt")

	fmt.Println(pt1)
	fmt.Println(pt2)
}

func part1(p string) int {
	i := make(chan string)
	go util.StreamInput(i, p)
	state := parseInput(i)

	dim := dim3.NewDimension(state)
	for dim.Cycle < 6 {
		dim.Simulate()
	}

	return dim.CountEndActive()
}

func part2(p string) int {
	i := make(chan string)
	go util.StreamInput(i, p)
	state := parseInput(i)

	dim := dim4.NewDimension(state)
	for dim.Cycle < 6 {
		dim.Simulate()
	}

	return dim.CountEndActive()
}

func parseInput(i <-chan string) []string {
	state := make([]string, 0)
	for l := range i {
		state = append(state, l)
	}
	return state
}
