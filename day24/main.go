package main

import (
	"fmt"
	"regexp"

	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("./input.txt")

	fmt.Println("Part 1:", pt1)
}

func part1(p string) int {
	i := make(chan string, 128)
	go util.StreamInput(i, p)
	instr := parseInput(i)
	floor := flipd(instr)
	return floor.countBlack()
}

func parseInput(i <-chan string) [][]string {
	instr := make([][]string, 0)
	r := regexp.MustCompile(`e|se|sw|w|nw|ne`)
	for i := range i {
		instr = append(instr, r.FindAllString(i, -1))
	}
	return instr
}

func flipd(instr [][]string) floor {
	f := make(floor)
	for _, i := range instr {
		l := lens{0, 0}
		for _, d := range i {
			l.move(d)
		}
		if t, exists := f[l]; exists {
			t.flip()
			t.commit()
		} else {
			f[l] = &tile{l, struct {
				pend string
				curr string
			}{"b", "b"}}
		}
	}
	return f
}

type lens [2]int

func (l *lens) move(d string) {
	switch d {
	case "e":
		l[0], l[1] = l[0]+1, l[1]
	case "se":
		l[0], l[1] = l[0], l[1]+1
	case "sw":
		l[0], l[1] = l[0]-1, l[1]+1
	case "w":
		l[0], l[1] = l[0]-1, l[1]
	case "nw":
		l[0], l[1] = l[0], l[1]-1
	case "ne":
		l[0], l[1] = l[0]+1, l[1]-1
	default:
		panic("bad dir")
	}
}

type tile struct {
	pos [2]int
	col struct {
		pend string
		curr string
	}
}

func (t *tile) flip() {
	if t.col.curr == "b" {
		t.col.pend = "w"
		return
	}
	if t.col.curr == "w" {
		t.col.pend = "b"
		return
	}
}

func (t *tile) commit() {
	t.col.curr = t.col.pend
}

type floor map[[2]int]*tile

func (f floor) countBlack() int {
	n := 0
	for _, t := range f {
		if t.col.curr == "b" {
			n++
		}
	}
	return n
}
