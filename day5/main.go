package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	p1 := part1()
	p2 := part2()
	fmt.Println("Part 1: ", p1)
	fmt.Println("Part 2: ", p2)
}

type pass struct {
	raw  string
	row  uint
	col  uint
	code uint
}

type passes []pass

func (ps passes) codeExists(c uint) bool {
	for _, p := range ps {
		if p.code == c {
			return true
		}
	}
	return false
}

func part1() uint {
	dec := decodeAllPasses(getInput())
	return dec[0].code
}

func part2() uint {
	dec := decodeAllPasses(getInput())
	for n := uint(0); n < dec[0].code; n++ {
		if !dec.codeExists(n) && dec.codeExists(n-1) && dec.codeExists(n+1) {
			return n
		}
	}
	panic("Not found")
}

func getInput() (i []string) {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic("No file!")
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		i = append(i, s.Text())
	}

	if s.Err() != nil {
		panic("File parse error!")
	}

	return
}

func decodeAllPasses(i []string) passes {
	dec := make(passes, 0)
	for _, p := range i {
		dec = append(dec, decodePass(p))
	}
	sort.Slice(dec, func(i, j int) bool {
		return dec[i].code > dec[j].code
	})
	return dec
}

func decodePass(s string) pass {
	rbin, cbin := toBinary(s[:7], 'F', 'B'), toBinary(s[7:], 'L', 'R')
	row, col := extract(rbin, cbin)
	code := row*8 + col
	pass := pass{
		s,
		row,
		col,
		code,
	}
	return pass
}

func toBinary(s string, l byte, u byte) string {
	s = strings.ReplaceAll(s, string(l), "0")
	s = strings.ReplaceAll(s, string(u), "1")
	return s
}

func extract(rbin string, cbin string) (uint, uint) {
	row := narrowPool(128, rbin)
	col := narrowPool(8, cbin)
	return row, col
}

func narrowPool(s uint, bin string) uint {
	pool := makePool(s)
	for _, bin := range bin {
		i, _ := strconv.ParseInt(string(bin), 0, 0)
		mid := len(pool) / 2
		opts := [2][]uint{pool[:mid], pool[mid:]}
		pool = opts[i]
	}
	if len(pool) != 1 {
		panic("pool should have been narrowed to one!")
	}
	return pool[0]
}

func makePool(s uint) []uint {
	pool := make([]uint, s)
	for i := range pool {
		pool[i] = uint(i)
	}
	return pool
}
