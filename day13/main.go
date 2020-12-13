package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("input.txt")
	pt2 := part2("input.txt")

	fmt.Println("Part 1:", pt1)
	fmt.Println("Part 2:", pt2)
}

func part1(p string) int {
	i := make(chan string)
	go util.StreamInput(i, p)
	d := calcDeptSum(i)
	return d
}

func part2(p string) int {
	i := make(chan string)
	go util.StreamInput(i, p)
	t := winCoinComp(i)
	return t
}

func calcDeptSum(i <-chan string) int {
	target, _ := strconv.ParseUint(<-i, 10, 0)

	buses := make(map[uint64][]uint64)
	for l := range i {
		for _, b := range strings.Split(l, ",") {
			if b != "x" {
				id, _ := strconv.ParseUint(b, 10, 0)
				intervals := make([]uint64, 0)
				for i := id; i <= target+id; i += id {
					intervals = append(intervals, i)
				}
				buses[id] = intervals
			}
		}
	}

	tally := struct{ id, wait uint64 }{0, target}
	for id, buses := range buses {
		b := buses[len(buses)-1]
		if b-target <= tally.wait {
			tally.id = id
			tally.wait = b - target
		}
	}

	return (int(tally.id) * int(tally.wait))
}

func winCoinComp(i <-chan string) int {
	<-i // discard first line
	buses := make([]uint64, 0)
	for l := range i {
		for _, b := range strings.Split(l, ",") {
			if b == "x" {
				buses = append(buses, 0)
			} else {
				id, _ := strconv.ParseUint(b, 10, 0)
				buses = append(buses, id)
			}
		}
	}

	done := make(chan int)
	go func() {
		var ts uint64

		for {
			found := true
			step := uint64(1)

			for o, v := range buses {
				if v == 0 {
					continue
				}
				if (ts+uint64(o))%v != 0 {
					found = false
					break
				}
				step *= v
			}

			if found {
				done <- int(ts)
			}

			ts += step
		}
	}()

	select {
	case res := <-done:
		return res
	case <-time.After(10 * time.Second):
		panic("timed out")
	}
}
