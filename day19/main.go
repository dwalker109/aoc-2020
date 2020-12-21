package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("./input_test.txt")
	pt2 := part2("./input_test.txt")

	fmt.Println("Part 1: ", pt1)
	fmt.Println("Part 2: ", pt2)
}

func part1(p string) int {
	i := make(chan string)
	go util.StreamInputCustomSplit(i, p, util.SplitOnString("\n\n"))

	rules, msgs := parseInput(i)
	compiled := compileRules(rules)

	n := 0
	for _, m := range msgs {
		valid, _ := validate([]string{m}, "0", &compiled, 0)
		if valid {
			n++
		}
	}

	return n
}

func part2(p string) int {
	i := make(chan string)
	go util.StreamInputCustomSplit(i, p, util.SplitOnString("\n\n"))

	rules, msgs := parseInput(i)
	compiled := compileRules(rules)
	perms, _ := build("0", &compiled)

	fmt.Println(perms, msgs)

	return 0
}

func parseInput(i <-chan string) ([]string, []string) {
	rules := strings.Split(strings.Trim(<-i, "\n"), "\n")
	msgs := strings.Split(strings.Trim(<-i, "\n"), "\n")

	return rules, msgs
}

type rule struct {
	id     string
	class  string
	abs    string
	chains [][]string
}

func compileRules(rules []string) map[string]rule {
	compiled := make(map[string]rule)

	for _, r := range rules {
		tmp := strings.SplitN(r, ":", 2)
		id, rest := tmp[0], tmp[1]

		// Terminator rule
		abs := regexp.MustCompile("(a|b)").FindString(rest)
		if abs != "" {
			compiled[id] = rule{id, "term", abs, make([][]string, 0)}
			continue
		}

		// Chain rule
		tmp = strings.Split(rest, "|")
		or := make([][]string, 0, len(tmp))
		for _, o := range tmp {
			or = append(or, strings.Fields(o))
		}

		compiled[id] = rule{id, "chain", "", or}
	}

	return compiled
}

func validate(msg []string, id string, rules *map[string]rule, lvl int) (bool, []string) {
	r := (*rules)[id]

	// End of a chain
	if r.class == "term" {
		remainders := make([]string, 0)
		for _, s := range msg {
			if s != "" && r.abs == string(s[0]) {
				remainders = append(remainders, s[1:])
			}
		}

		return len(remainders) > 0, remainders
	}

	// Traverse a chain
	hits := make([]string, 0)
Outer:
	for _, set := range r.chains {
		rem := msg
		for _, id := range set {
			res, newRem := validate(rem, id, rules, lvl+1)
			if !res {
				continue Outer
			}
			rem = newRem
		}

		hits = append(hits, rem...)
	}

	if lvl == 0 && len(hits) > 0 && hits[0] != "" {
		return false, hits
	}

	return len(hits) != 0, hits
}

func build(id string, rules *map[string]rule) ([]string, bool) {
	z := make([]map[int]map[int]string, 0)

	var recurse func(string, int) string
	recurse = func(id string, lvl int) string {
		r := (*rules)[id]
		if r.class == "term" {
			return r.abs
		}

		x := make(map[int]map[int]string)
		arr := make([]string, len(r.chains))
		var rr string
		for _, or := range r.chains {
			for i, and := range or {
				res := recurse(and, lvl+1)
				arr[i] += res
				fmt.Println(lvl, i, "Rule", id, "or", or, "and", and, ":", res)
			}
			rr = strings.Join(arr, "|")
		}

		fmt.Println(rr)

		z = append(z, x)

		return ""
	}
	recurse(id, 0)
	fmt.Println(z)
	return []string{}, false
}
