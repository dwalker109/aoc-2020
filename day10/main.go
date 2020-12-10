package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/dwalker109/aoc-2020/util"
)

type adapters []int

func (ax adapters) contains(x int) bool {
	for _, v := range ax {
		if v == x {
			return true
		}
	}
	return false
}

func (ax adapters) outletJolts() int {
	return 0
}

func (ax adapters) builtinJolts() int {
	return ax[len(ax)-1]
}

func main() {
	p1 := part1("./input.txt")
	p2 := part2("./input.txt")

	fmt.Println("Part 1: ", p1)
	fmt.Println("Part 2: ", p2)
}

func part1(p string) int {
	i := make(chan string, 128)
	go util.StreamInput(i, p)
	a := getSortedAdapters(i)
	d := calcJoltageDifference(a)
	return d
}

func part2(p string) int {
	i := make(chan string, 128)
	go util.StreamInput(i, p)
	a := getSortedAdapters(i)
	n := calcAdapterArrangements(a)
	return n
}

func getSortedAdapters(i <-chan string) adapters {
	a := make([]int, 0)
	for s := range i {
		j, _ := strconv.ParseInt(s, 0, 0)
		a = append(a, int(j))
	}
	sort.Ints(a)
	a = append(a, a[len(a)-1]+3)
	return a
}

func calcJoltageDifference(ax adapters) int {
	j := ax.outletJolts()
	s := make(map[int]int)
	for _, a := range ax {
		d := a - j
		s[d]++
		j += d
	}
	return s[1] * s[3]
}

func calcAdapterArrangements(ax adapters) int {
	routes := make([]int, 1, len(ax))
	routes[ax.outletJolts()] = 1
	ax = append(ax, ax.builtinJolts())
	for x := ax.outletJolts(); x <= ax.builtinJolts(); x++ {
		// Check each possble adapter within range (+3) of x
		for y := x + 1; y <= x+3; y++ {
			// If y (the child) is an adapter, increment it's count
			// with the qty of routes to x (the parent)
			if ax.contains(y) {
				routes[y] += routes[x]
			}
		}
	}
	return routes[ax.builtinJolts()]
}
