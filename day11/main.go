package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/dwalker109/aoc-2020/util"
	"regexp"
	"strings"
)

func main() {
	p1 := part1("./input.txt")

	fmt.Println("Part 1:", p1)
}

func part1(p string) int {
	i := make(chan string, 128)
	go util.StreamInput(i, p)
	grid := parseInput(i)
	c := fillSeats(grid)
	return c
}

func fillSeats(grid seatgrid) (c int) {
	orig := cloneGridVals(grid)
	for {
		orig = cloneGridVals(grid)
		for row := 0; row <= grid.lastY(); row++ {
			for seat := 0; seat <= grid.lastX(); seat++ {
				if orig.shouldToggleOccupancy(row, seat) {
					grid.toggle(row, seat)
				}
			}
		}
		if orig.sum() == grid.sum() {
			break
		}
	}

	return grid.occupied()
}

func cloneGridVals(source seatgrid) seatgrid {
	clone := make(seatgrid, len(source))
	for i, row := range source {
		rowclone := make([]string, len(row))
		copy(rowclone, row)
		clone[i] = rowclone
	}
	return clone
}

type seatgrid [][]string

func (sg seatgrid) shouldToggleOccupancy(y, x int) bool {
	curr := sg[y][x]

	sx := util.MaxInt(x-1, 0)
	ex := util.MinInt(x+1, sg.lastX())
	sy := util.MaxInt(y-1, 0)
	ey := util.MinInt(y+1, sg.lastY())

	var sb strings.Builder
	for j := sy; j <= ey; j++ {
		for i := sx; i <= ex; i++ {
			if i == x && j == y {
				sb.WriteString("+")
			} else {
				sb.WriteString(sg[j][i])
			}
		}
	}

	check := sb.String()
	matches := regexp.MustCompile("#").FindAllStringIndex(check, -1)

	if curr == "L" && len(matches) == 0 {
		return true
	}

	if curr == "#" && len(matches) >= 4 {
		return true
	}

	return false
}

func (sg *seatgrid) toggle(y, x int) {
	curr := (*sg)[y][x]
	if curr == "L" {
		(*sg)[y][x] = string('#')
	}
	if curr == "#" {
		(*sg)[y][x] = string('L')
	}

}

func (sg seatgrid) lastY() int {
	return len(sg) - 1
}

func (sg seatgrid) lastX() int {
	return len(sg[0]) - 1
}

func (sg seatgrid) sum() [32]byte {
	var sb strings.Builder
	for _, r := range sg {
		sb.WriteString(strings.Join(r, "-"))
	}
	return sha256.Sum256([]byte(sb.String()))
}

func (sg seatgrid) occupied() (o int) {
	for _, row := range sg {
		check := strings.Join(row, "")
		matches := regexp.MustCompile("#").FindAllStringIndex(check, -1)
		o += len(matches)
	}
	return
}

func parseInput(i <-chan string) seatgrid {
	grid := make(seatgrid, 0)
	for l := range i {
		row := make([]string, len(l), len(l))
		for i, s := range strings.Split(l, "") {
			row[i] = s
		}
		grid = append(grid, row)
	}
	return grid
}
