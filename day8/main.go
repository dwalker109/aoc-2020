package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/dwalker109/aoc-2020/util"
)

const (
	acc = uint8(255)
	jmp = uint8(254)
	nop = uint8(253)
)

type instruction struct {
	op  uint8
	val int
}

type log []int

func (l log) used(id int) bool {
	for _, v := range l {
		if v == id {
			return true
		}
	}
	return false
}

func main() {
	p1 := part1()
	p2 := part2()

	fmt.Println("Part 1: ", p1)
	fmt.Println("Part 2: ", p2)
}

func part1() int {
	i := make(chan string, 4096)

	go util.StreamInput(i)
	n, _ := maxAccumulator(parseInput(i))

	return n
}

func part2() int {
	i := make(chan string, 4096)
	variants := make(chan []instruction, 4096)

	go util.StreamInput(i)
	go fixOps(parseInput(i), variants)

	for variant := range variants {
		n, err := maxAccumulator(variant)
		if err == nil {
			close(variants)
			return n
		}
	}
	close(variants)

	panic("ruh roh")
}

func parseInput(i <-chan string) []instruction {
	ops := make([]instruction, 0)
	for l := range i {
		field := strings.Fields(l)
		val, _ := strconv.ParseInt(field[1], 0, 0)

		var op uint8
		switch field[0] {
		case "acc":
			op = acc
		case "jmp":
			op = jmp
		case "nop":
			op = nop
		}

		ops = append(ops, instruction{op, int(val)})
	}

	return ops
}

func maxAccumulator(ops []instruction) (int, error) {
	accu := 0
	hist := make(log, 0)
	next := 0

	for hist.used(next) == false {
		if next == len(ops) {
			return accu, nil
		}

		if next > len(ops) {
			return accu, errors.New("out of range")
		}

		cur := ops[next]
		hist = append(hist, next)

		switch cur.op {
		case acc:
			accu += cur.val
			next++
		case jmp:
			next += cur.val
		case nop:
			next++
		}
	}

	return accu, errors.New("infinite loop")
}

func fixOps(i []instruction, c chan<- []instruction) {
	for n, curr := range i {
		switch curr.op {
		case nop:
			curr.op = jmp
		case jmp:
			curr.op = nop
		}

		cp := append([]instruction(nil), i...)
		cp[n] = curr
		c <- cp
	}
}
