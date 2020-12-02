package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PR struct {
	n1   int
	n2   int
	char string
	pw   string
}

func getPasswords() []PR {
	data := make([]PR, 0)
	for _, s := range MyInput {
		parts := strings.Replace(s, ":", "", 1)
		split := strings.SplitN(parts, " ", 3)
		numRange, char, pass := split[0], split[1], split[2]

		nx := strings.SplitN(numRange, "-", 2)
		sn1, sn2 := nx[0], nx[1]
		n1, _ := strconv.ParseInt(sn1, 10, 0)
		n2, _ := strconv.ParseInt(sn2, 10, 0)

		data = append(data, PR{int(n1), int(n2), char, pass})
	}

	return data
}

func part1(data []PR) (vq int) {
	vq = 0
	for _, pr := range data {
		r := regexp.MustCompile(pr.char)
		m := r.FindAllStringIndex(pr.pw, -1)
		if len(m) >= pr.n1 && len(m) <= pr.n2 {
			vq = vq + 1
		}
	}
	return
}

func part2(data []PR) (vq int) {
	vq = 0
	for _, pr := range data {
		pw := strings.Split(pr.pw, "")
		p1, p2 := pr.n1-1, pr.n2-1

		m1 := pw[p1] == pr.char
		m2 := pw[p2] == pr.char

		if (m1 || m2) && (m1 != m2) {
			vq = vq + 1
		}
	}
	return
}

func main() {
	input := getPasswords()
	sol1 := part1(input)
	fmt.Println("Part1: ", sol1)

	sol2 := part2(input)
	fmt.Println("Part2: ", sol2)
}
