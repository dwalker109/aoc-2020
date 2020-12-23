package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	inputTest = "389125467"
	input     = "653427918"
)

func main() {
	pt1 := part1(input)
	pt2 := part2(inputTest)

	fmt.Println(pt1)
	fmt.Println(pt2)
}

func part1(in string) string {
	cups := make([]uint8, len(in))
	for i := 0; i < len(in); i++ {
		val, _ := strconv.Atoi(string(in[i]))
		cups[i] = uint8(val)
	}

	lblToPos := func(lbl uint8, pool []uint8) int {
		for i, v := range pool {
			if v == lbl {
				return i
			}
		}
		panic("bad label")
	}

	ccIdx := 0
	hi := uint8(0)
	for _, v := range cups {
		if v > hi {
			hi = v
		}
	}

	for t := 0; t < 100; t++ {

		// No wrap, no shift
		wip := make([]uint8, 0, len(cups))
		wip = cups[ccIdx:]
		if ccIdx > 0 {
			wip = append(wip, cups[0:ccIdx]...)
		}
		take := wip[1:4]
		rest := wip[4:]
		wip = append([]uint8{wip[0]}, rest...)

		insLbl := wip[0] - 1
	insLblChk:
		for {
			if insLbl < 1 {
				insLbl = hi
			}
			for _, v := range take {
				if insLbl == v {
					insLbl--
					continue insLblChk
				}
			}
			break
		}
		insIdx := lblToPos(insLbl, wip)

		wipCpy := make([]uint8, len(wip))
		copy(wipCpy, wip)
		next := append(wipCpy[0:insIdx+1], take...)
		next = append(next, wip[insIdx+1:]...)

		if ccIdx > 0 {
			rarr := make([]uint8, 0, len(cups))
			rarr = append(rarr, next[len(next)-ccIdx:]...)
			rarr = append(rarr, next[:len(next)-ccIdx]...)
			cups = rarr
		} else {
			cups = next
		}

		ccIdx = (ccIdx + 1) % len(cups)
	}

	sb := new(strings.Builder)
	cup1 := lblToPos(uint8(1), cups)
	for _, n := range cups[cup1+1:] {
		s := strconv.Itoa(int(n))
		sb.Write([]byte(s))
	}
	for _, n := range cups[:cup1] {
		s := strconv.Itoa(int(n))
		sb.Write([]byte(s))
	}

	return sb.String()
}

func part2(in string) int {
	cups := make(map[uint]*cup)
	for i := uint(1_000_000); i >= 1; i-- {
		c := cup{lbl: i}
		if nc, exists := cups[i+1]; exists {
			c.nxt = nc
		}
		cups[i] = &c
	}
	cups[1_000_000].nxt = cups[1]

	hi := 0
	for i := 0; i < len(in); i++ {
		n1, _ := strconv.Atoi(string(in[i]))
		var n2 int

		if n1 > hi {
			hi = n1
		}

		if i+1 == len(in) {
			n2 = hi + 1
		} else {
			n2, _ = strconv.Atoi(string(in[i+1]))
		}

		c1 := cups[uint(n1)]
		c2 := cups[uint(n2)]
		c1.nxt = c2
	}

	for i := 0; i < len(in); i++ {
		val, _ := strconv.Atoi(string(in[i]))
		c := cups[uint(val)]
		var nc *cup
		if c.lbl == uint(len(in)) {
			nc = cups[uint(len(in))]
		} else {
			nc = cups[uint(in[c.lbl+1])]
		}
		c.nxt = nc
	}

	return 0
}

type cup struct {
	lbl uint
	nxt *cup
}
