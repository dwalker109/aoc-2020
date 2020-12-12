package main

import (
	"fmt"
	"strconv"

	"github.com/dwalker109/aoc-2020/day12/ship"
	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	p1 := part1("./input.txt")

	fmt.Println("Part 1: ", p1)
}

func part1(p string) int {
	i := make(chan string, 128)
	go util.StreamInput(i, p)
	md := navigate(i)
	return md
}

func navigate(i <-chan string) int {
	s := ship.NewSimpleShip(0, 0, 90)

	for l := range i {
		com := string(l[0])
		dis, _ := strconv.ParseInt(l[1:], 10, 0)

		switch com {
		case "F":
			s.Forward(int(dis))
		case "L":
			s.Left(int(dis))
		case "R":
			s.Right((int(dis)))
		default:
			s.MoveDir(com, int(dis))
		}
	}

	return s.Manhattan()
}
