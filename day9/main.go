package main

import (
	"fmt"
	"github.com/dwalker109/aoc-2020/util"
	"math"
	"sort"
	"strconv"
)

const (
	apertureSize = 25
)

func main() {
	p1 := part1()
	p2 := part2(p1)

	fmt.Println("Part 1: ", p1)
	fmt.Println("Part 2: ", p2)
}

func part1() int64 {
	i := make(chan string, 128)
	a := make(chan []int64, 128)

	go util.StreamInput(i, "./input.txt")
	go streamAperture(i, a)
	return notPrevApertureSum(a)
}

func part2(target int64) int64 {
	i := make(chan string, 128)
	go util.StreamInput(i, "./input.txt")
	return findEncWeakness(i, target)
}

func streamAperture(i <-chan string, a chan<- []int64) {
	ap := make([]int64, 0)
	for str := range i {
		next, _ := strconv.ParseInt(str, 0, 0)
		from := int(math.Max(0, float64(len(ap)-apertureSize)))
		ap = append(ap[from:], next)
		a <- ap
	}
	close(a)
}

func notPrevApertureSum(a <-chan []int64) int64 {
	for nums := range a {
		if len(nums) < apertureSize+1 {
			continue
		}

		pool := nums[0:apertureSize]
		target := nums[apertureSize]
		found := false

		for i1, n1 := range pool {
			for i2, n2 := range pool {
				if i1 != i2 && n1+n2 == target {
					found = true
				}
			}
		}

		if !found {
			return target
		}
	}

	panic("ruh roh 1!")
}

func findEncWeakness(i <-chan string, target int64) int64 {
	nums := make([]int64, len(i))
	for str := range i {
		num, _ := strconv.ParseInt(str, 0, 0)
		nums = append(nums, num)
	}

	for n := 0; n < len(nums); n++ {
		total := nums[n]
		tally := []int{int(total)}
		for n2 := n + 1; total < target; n2++ {
			total += nums[n2]
			tally = append(tally, int(nums[n2]))
			if total == target {
				sort.Ints(tally)
				return int64(tally[0] + tally[len(tally)-1])
			}
		}
	}

	panic("ruh roh 2!")
}
