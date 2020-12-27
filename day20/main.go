package main

import (
	"fmt"

	"github.com/dwalker109/aoc-2020/day20/sat"
	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("./input.txt")
	pt2 := part2("./input.txt")

	fmt.Println("Part 1:", pt1)
	fmt.Println("Part 2:", pt2)
}

func part1(p string) int {
	img := parseInput(p)
	img.Assemble()
	return img.Matched[0][0].ID *
		img.Matched[0][img.GridSize-1].ID *
		img.Matched[img.GridSize-1][0].ID *
		img.Matched[img.GridSize-1][img.GridSize-1].ID
}

func part2(p string) int {
	img := parseInput(p)
	img.Assemble()
	img.BuildActual()
	img.SurfaceSeaMonsters()
	return img.CalcWaterRoughness()
}

func parseInput(p string) *sat.Img {
	i := make(chan string)
	go util.StreamInputCustomSplit(i, p, util.SplitOnString("\n\n"))

	tiles := make([]*sat.Tile, 0)
	for t := range i {
		tiles = append(tiles, sat.NewGarbledTile(t))
	}

	return sat.NewImg(tiles)
}
