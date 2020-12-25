package main

import (
	"fmt"
	"regexp"

	"github.com/dwalker109/aoc-2020/day24/hexx"
	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("./input.txt")
	pt2 := part2("./input.txt")

	fmt.Println("Part 1:", pt1)
	fmt.Println("Part 2:", pt2)
}

func part1(p string) int {
	f := initFloor(p)
	return f.CountBlack()
}
func part2(p string) int {
	f := initFloor(p)
	for n := 1; n <= 100; n++ {
		grow(f)
		art(f)
	}
	return f.CountBlack()
}

func initFloor(p string) *hexx.Floor {
	i := make(chan string, 128)
	go util.StreamInput(i, p)
	instr := parseInput(i)
	floor := flipd(instr)
	return &floor
}

func parseInput(i <-chan string) [][]string {
	instr := make([][]string, 0)
	r := regexp.MustCompile(`e|se|sw|w|nw|ne`)
	for i := range i {
		instr = append(instr, r.FindAllString(i, -1))
	}
	return instr
}

func flipd(instr [][]string) hexx.Floor {
	f := make(hexx.Floor)
	for _, i := range instr {
		l := hexx.Lens{0, 0}
		for _, d := range i {
			l.Move(d)
		}
		if t, exists := f[l]; exists {
			t.Flip()
			t.Commit()
		} else {
			t := hexx.NewTile(l)
			t.Flip()
			t.Commit()
			f[l] = t
		}
	}
	return f
}

func grow(f *hexx.Floor) {
	// Stop growing beyond 10x the num of black tiles - works
	// with this input, might not with other input!
	if len(*f) > f.CountBlack()*10 {
		return
	}
	for _, t := range *f {
		for _, c := range t.AdjacentCoords() {
			if _, exists := (*f)[c]; !exists {
				(*f)[c] = hexx.NewTile(c)
			}
		}
	}
}

func art(f *hexx.Floor) {
	for _, t := range *f {
		bc := 0
		for _, c := range t.AdjacentCoords() {
			at := (*f)[c]
			if at != nil && at.IsBlack() {
				bc++
			}
		}
		if (t.IsBlack() && (bc == 0 || bc > 2)) || t.IsWhite() && bc == 2 {
			t.Flip()
		}
	}
	f.CommitAll()
}
