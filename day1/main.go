package main

import (
	"fmt"
)

const target = 2020

func part1(c chan <-int) {
	for _, in1 := range MyInput {
		for _, in2 := range MyInput {
			if in1+in2 == target {
				c <- in1 * in2
			}
		}
	}
}

func part2(c chan <-int) {
	for _, in1 := range MyInput {
		for _, in2 := range MyInput {
			for _, in3 := range MyInput {
				if in1+in2+in3 == target {
					c <- in1 * in2 * in3
				}
			}
		}
	}
}

func main() {
	r1 := make(chan int, 1)
	r2 := make(chan int, 1)

	go part1(r1)
	go part2(r2)

	fmt.Println("Part 1 result: ", <-r1)
	fmt.Println("Part 2 result: ", <-r2)
}