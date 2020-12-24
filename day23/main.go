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
	pt2 := part2(input)

	fmt.Println("Part 1", pt1)
	fmt.Println("Part 2", pt2)
}

func part1(in string) string {
	numCups := len(in)
	numTurns := 100

	cups, initial := unpackCups(in, numCups, numTurns)
	cups = play(cups, initial, numCups, numTurns)

	sb := new(strings.Builder)
	c1 := (*cups)[uint(1)]
	tc := c1.nxt
	for tc != c1 {
		sb.Write([]byte(strconv.Itoa(int(tc.lbl))))
		tc = tc.nxt
	}

	return sb.String()
}

func part2(in string) int {
	numCups := 1_000_000
	numTurns := 10_000_000

	cups, initial := unpackCups(in, numCups, numTurns)
	cups = play(cups, initial, numCups, numTurns)

	c1 := (*cups)[uint(1)]
	return int(c1.nxt.lbl * c1.nxt.nxt.lbl)
}

type cup struct {
	lbl uint
	nxt *cup
}

type cupTardis map[uint]*cup

func unpackCups(in string, numCups, numTurns int) (*cupTardis, *cup) {
	// Create all cups, setting .nxt (final nil .nxt will be wrapped later)
	cups := make(cupTardis)
	for i := uint(numCups); i >= 1; i-- {
		cups[i] = &cup{lbl: i, nxt: cups[i+1]}
	}

	// Set .nxt on seed cups
	for i := 0; i < len(in); i++ {
		n1, _ := strconv.Atoi(string(in[i]))

		var n2 int
		if i+1 == len(in) {
			n2 = len(in) + 1
		} else {
			n2, _ = strconv.Atoi(string(in[i+1]))
		}

		c1 := cups[uint(n1)]
		c2 := cups[uint(n2)]
		c1.nxt = c2
	}

	// Get the current cup from seeds
	initLbl, _ := strconv.Atoi(string(in[0]))
	cc := cups[uint(initLbl)]

	// Setup wrapping from max->min
	if numCups > len(in) {
		cups[uint(numCups)].nxt = cc
	} else {
		lastLbl, _ := strconv.Atoi(string(in[len(in)-1]))
		cups[uint(lastLbl)].nxt = cc
	}
	return &cups, cc
}

func play(cupsPtr *cupTardis, cc *cup, numCups, numTurns int) *cupTardis {
	cups := *cupsPtr

	for t := 1; t <= numTurns; t++ {
		// Pick up next three
		puc1 := cc.nxt
		puc2 := cc.nxt.nxt
		puc3 := cc.nxt.nxt.nxt

		// Get the destination cup
		nxtLbl := cc.lbl - 1
		destc := cups[nxtLbl]
		for destc == nil || destc == puc1 || destc == puc2 || destc == puc3 {
			if nxtLbl == 0 {
				nxtLbl = uint(numCups)
				destc = cups[nxtLbl]
				continue
			}
			nxtLbl--
			destc = cups[nxtLbl]
		}

		cc.nxt = puc3.nxt    // Remove next three (picked up) cups
		puc3.nxt = destc.nxt // Insert final picked up cup after destination cup
		destc.nxt = puc1     // Point destination cup at first picked up cup

		cc = cc.nxt
	}

	return &cups
}
