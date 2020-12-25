package main

import (
	"fmt"
	"regexp"

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
	return f.countBlack()
}
func part2(p string) int {
	f := initFloor(p)
	for n := 1; n <= 100; n++ {
		grow(f)
		art(f)
	}
	return f.countBlack()
}

func initFloor(p string) *floor {
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
			t := newTile(l)
			t.flip()
			t.commit()
			f[l] = t
		}
	}
	return f
}

func grow(f *floor) {
	for _, t := range *f {
		for _, c := range t.adjacentCoords() {
			if _, exists := (*f)[c]; !exists {
				(*f)[c] = newTile(c)
			}
		}
	}
}

func art(f *floor) {
	for _, t := range *f {
		bc := 0
		for _, c := range t.adjacentCoords() {
			at := (*f)[c]
			if at != nil && at.col.curr == "b" {
				bc++
			}
		}
		if t.col.curr == "b" && (bc == 0 || bc > 2) {
			t.col.pend = "w"
		}
		if t.col.curr == "w" && bc == 2 {
			t.col.pend = "b"
		}
	}
	f.commitAll()
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

func newTile(c [2]int) *tile {
	return &tile{c, struct {
		pend string
		curr string
	}{"w", "w"}}
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
	if t.col.curr != t.col.pend {
		t.col.curr = t.col.pend
	}
}

func (t *tile) adjacentCoords() [6][2]int {
	qr := t.pos
	adj := [6][2]int{
		{qr[0] + 1, qr[1]},
		{qr[0], qr[1] + 1},
		{qr[0] - 1, qr[1] + 1},
		{qr[0] - 1, qr[1]},
		{qr[0], qr[1] - 1},
		{qr[0] + 1, qr[1] - 1},
	}
	return adj
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

func (f *floor) commitAll() {
	for _, t := range *f {
		t.commit()
	}
}
