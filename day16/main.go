package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("./input.txt")
	pt2 := part2("./input.txt")

	fmt.Println("Part 1: ", pt1)
	fmt.Println("Part 2: ", pt2)
}

func part1(p string) int {
	rules, _, others := parseNotes(p)
	tser := calcTSER(rules, others)

	return int(tser)
}

func part2(p string) int {
	rules, my, others := parseNotes(p)
	valid := withoutInvalid(rules, others)
	im := calcIndexMap(rules, append(valid, my))
	n := calcDepartureVal(im, my)
	return int(n)
}

type rulesValidator map[string]func(n int64) bool
type ticket []int64

func parseNotes(p string) (rulesValidator, ticket, []ticket) {
	i := make(chan string)
	go util.StreamInputCustomSplit(i, p, util.SplitOnString("\n\n"))

	// Rules
	rules := make(rulesValidator)
	for _, s := range strings.Split(<-i, "\n") {
		kv := strings.Split(s, ":")
		k, v := kv[0], strings.Fields(kv[1])
		r1, r2 := v[0], v[2]
		rules[k] = func(n int64) bool {
			for _, r := range []string{r1, r2} {
				p := strings.Split(r, "-")
				lo, _ := strconv.ParseInt(p[0], 10, 64)
				hi, _ := strconv.ParseInt(p[1], 10, 64)
				if n >= lo && n <= hi {
					return true
				}
			}
			return false
		}
	}

	extractTicket := func(s string) ticket {
		t := make(ticket, 0)
		for _, i := range strings.Split(s, ",") {
			j, _ := strconv.ParseInt(i, 10, 64)
			t = append(t, j)
		}
		return t
	}

	// My ticket
	mt := strings.Split(strings.Trim(<-i, "\n"), "\n")[1]
	myTicket := extractTicket(mt)

	// Other tickets
	otherTickets := make([]ticket, 0)
	for _, t := range strings.Split(strings.Trim(<-i, "\n"), "\n")[1:] {
		otherTickets = append(otherTickets, extractTicket(t))
	}

	return rules, myTicket, otherTickets
}

func calcTSER(validator rulesValidator, tickets []ticket) (tser int64) {
	for _, t := range tickets {
	LoopTicketVals:
		for _, n := range t {
			for _, fn := range validator {
				if fn(n) {
					continue LoopTicketVals
				}
			}
			tser += n
		}
	}
	return
}

func withoutInvalid(validator rulesValidator, tickets []ticket) []ticket {
	valid := make([]ticket, 0)
	for _, t := range tickets {
		tValid := true
		for _, n := range t {
			nValid := false
			for _, fn := range validator {
				if fn(n) {
					nValid = true
				}
			}
			if nValid == false {
				tValid = false
			}
		}
		if tValid {
			valid = append(valid, t)
		}
	}
	return valid
}

func calcIndexMap(validator rulesValidator, tickets []ticket) map[string]int {
	indexMap := make(map[string]int)

	ticketsByCol := make([][]int64, len(validator))
	for i := 0; i < len(ticketsByCol); i++ {
		for _, ticket := range tickets {
			ticketsByCol[i] = append(ticketsByCol[i], ticket[i])
		}
	}

	for len(indexMap) != len(validator) {
	NextCol:
		for colNum, colTickets := range ticketsByCol {
			matches := make(map[string]bool)

			// This col already mapped, move to next
			for _, fc := range indexMap {
				if fc == colNum {
					continue NextCol
				}
			}

			// Test each field
			for key, fn := range validator {
				// This key already mapped, move to next
				if _, found := indexMap[key]; found {
					continue
				}

				// Test each value in this column against the validator
				for _, v := range colTickets {
					if fn(v) {
						// Val valid - signal true exactly once, if nothing else has already signaled something
						if _, exists := matches[key]; !exists {
							matches[key] = true
						}
					} else {
						// Invalid value encountered - this key isn't for this column
						matches[key] = false
					}
				}
			}

			// Discount any keys we know aren't for this column
			for mk, mv := range matches {
				if mv == false {
					delete(matches, mk)
				}
			}

			// Exactly one key remains - bingo
			if len(matches) == 1 {
				for key := range matches {
					indexMap[key] = colNum
					continue NextCol
				}
			}
		}
	}

	return indexMap
}

func calcDepartureVal(indexMap map[string]int, t ticket) int64 {
	n := int64(1)
	for k, i := range indexMap {
		if strings.Contains(k, "departure ") {
			n *= t[i]
		}
	}
	return n
}
