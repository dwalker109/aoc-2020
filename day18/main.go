package main

import (
	"fmt"
	"regexp"
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
	return solve(p, evalNormal)
}

func part2(p string) int {
	return solve(p, evalCrayCray)
}

func solve(p string, evalFn func(string) (int, string)) int {
	i := make(chan string, 64)
	go util.StreamInput(i, p)

	sum := 0
	exp := expand(i, evalFn)
	for _, e := range exp {
		res, _ := evalFn(e)
		sum += res
	}

	return sum
}

func expand(i <-chan string, evalFn func(string) (int, string)) []string {
	e := make([]string, 0)

	for s := range i {
		r := regexp.MustCompile(`\([^(|^)]*\)`)
		for {
			loc := r.FindIndex([]byte(s))

			if len(loc) == 0 {
				break
			}

			str := s[loc[0]:loc[1]]
			_, res := evalFn(str)
			s = strings.ReplaceAll(s, str, res)
		}
		e = append(e, s)
	}

	return e
}

func evalNormal(s string) (int, string) {
	s = strings.TrimPrefix(strings.TrimSuffix(s, ")"), "(")

	exp := strings.Fields(s)
	sum, _ := strconv.ParseInt(exp[0], 10, 64)

	for i := 1; i < len(exp); i += 2 {
		op, nxt := exp[i], exp[i+1]
		nxtN, _ := strconv.ParseInt(nxt, 10, 64)

		switch op {
		case "+":
			sum += nxtN
		case "*":
			sum *= nxtN
		default:
			panic("bad op")
		}
	}

	return int(sum), fmt.Sprint(sum)
}

func evalCrayCray(s string) (int, string) {
	s = strings.TrimPrefix(strings.TrimSuffix(s, ")"), "(")

	// Handle additions
	r := regexp.MustCompile(`\d+\s+\+\s+\d+`)
	for {
		loc := r.FindIndex([]byte(s))

		if len(loc) == 0 {
			break
		}

		str := s[loc[0]:loc[1]]
		exp := strings.Fields(str)

		lhs, _ := strconv.ParseInt(exp[0], 10, 64)
		rhs, _ := strconv.ParseInt(exp[2], 10, 64)
		res := fmt.Sprint(lhs + rhs)

		s = strings.ReplaceAll(s, str, res)
	}

	// Handle what's left, which will only be multiplications
	exp := strings.Fields(strings.ReplaceAll(s, "*", ""))
	sum, _ := strconv.ParseInt(exp[0], 10, 64)

	for _, nxt := range exp[1:] {
		nxtN, _ := strconv.ParseInt(nxt, 10, 64)
		sum *= nxtN
	}

	return int(sum), fmt.Sprint(sum)
}
