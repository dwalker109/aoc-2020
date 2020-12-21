package main

import (
	"fmt"

	"github.com/dwalker109/aoc-2020/day20/tile"

	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("./input.txt")
	pt2 := part2("./input_test.txt")

	fmt.Println("Part 1: ", pt1)
	fmt.Println("Part 2: ", pt2)
}

func part1(p string) int {
	tiles := parseInput(p)
	return findCornerSum(tiles)
}

func part2(p string) int {
	tiles := parseInput(p)
	corners := findCorners(tiles)
	img := makeImg(tiles, corners)
}

func parseInput(p string) tile.Tiles {
	i := make(chan string)
	go util.StreamInputCustomSplit(i, p, util.SplitOnString("\n\n"))

	tiles := make(tile.Tiles, 0, 100)
	for dat := range i {
		tiles = append(tiles, tile.NewTile(dat))
	}

	return tiles
}

func findCornerSum(t tile.Tiles) int {
	r := 1
	for _, tile := range findCorners(t) {
		r *= tile.ID
	}

	return r
}

func findCorners(t tile.Tiles) tile.Tiles {
	e := make(map[string][]string)
	for _, t := range t {
		t.Print()
		e[string(t.Top)] = append(e[string(t.Top)], t.Id)
		e[string(t.Btm)] = append(e[string(t.Btm)], t.Id)
		e[string(t.Lft)] = append(e[string(t.Lft)], t.Id)
		e[string(t.Rgt)] = append(e[string(t.Rgt)], t.Id)

		e[string(util.ReverseBytes(t.Top))] = append(e[string(util.ReverseBytes(t.Top))], t.Id)
		e[string(util.ReverseBytes(t.Btm))] = append(e[string(util.ReverseBytes(t.Btm))], t.Id)
		e[string(util.ReverseBytes(t.Lft))] = append(e[string(util.ReverseBytes(t.Lft))], t.Id)
		e[string(util.ReverseBytes(t.Rgt))] = append(e[string(util.ReverseBytes(t.Rgt))], t.Id)
	}

	c := make(map[string]int, len(t))
	for _, ids := range e {
		if len(ids) == 1 {
			c[ids[0]]++
		}
	}

	r := make(tile.Tiles, 0, 4)
	for id, c := range c {
		if c == 4 {
			r[id] = t[id]
		}
	}

	return r
}

func makeImage(t tile.Tiles, c tile.Tiles) [][]byte {

}
