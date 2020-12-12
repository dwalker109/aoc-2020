package main

import (
	"fmt"
	"strconv"

	"github.com/dwalker109/aoc-2020/day12/ship"
	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	p1 := part1("./input.txt")
	p2 := part2("./input.txt")

	fmt.Println("Part 1: ", p1)
	fmt.Println("Part 2: ", p2)
}

func part1(p string) int {
	i := make(chan string, 128)
	go util.StreamInput(i, p)
	s := ship.NewSimpleShip(0, 0, 90)
	md := navigate(i, s)
	return md
}

func part2(p string) int {
	i := make(chan string, 128)
	go util.StreamInput(i, p)
	s := ship.NewWaypointShip(0, 0, 10, 1)
	md := navigate(i, s)
	return md
}

func navigate(i <-chan string, s ship.Ship) int {
	for instr := range i {
		comm := string(instr[0])
		dist, err := strconv.ParseInt(instr[1:], 10, 0)
		if err != nil {
			panic("invalid instruction")
		}

		switch comm {
		case "F":
			s.Forward(int(dist))
		case "L":
			s.Left(int(dist))
		case "R":
			s.Right((int(dist)))
		default:
			s.Move(comm, int(dist))
		}
	}

	return s.Manhattan()
}
