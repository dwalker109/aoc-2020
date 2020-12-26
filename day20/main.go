package main

import (
	"fmt"

	"github.com/dwalker109/aoc-2020/day20/sat"
	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("./input_test.txt")

	fmt.Println("Part 1:", pt1)
}

func part1(p string) int {
	img := parseInput(p)
	img.Assemble()
	return img.Assembled[0][0].ID *
		img.Assembled[0][img.GridSize-1].ID *
		img.Assembled[img.GridSize-1][0].ID *
		img.Assembled[img.GridSize-1][img.GridSize-1].ID
}

func parseInput(p string) *sat.Img {
	i := make(chan string)
	go util.StreamInputCustomSplit(i, p, util.SplitOnString("\n\n"))

	tiles := make([]*sat.Tile, 0)
	for t := range i {
		tiles = append(tiles, sat.NewTile(t))
	}

	return sat.NewImg(tiles)
}
