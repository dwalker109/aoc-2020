package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() (i [][]string) {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic("No file!")
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		i = append(i, strings.Split(s.Text(), ""))
	}

	if s.Err() != nil {
		panic("File parse error!")
	}

	return
}

func traverse(sx, sy int) (t int) {
	data := getInput()
	x, y := 0, 0
	maxx, maxy := len(data[0]), len(data)-1

	for y < maxy {
		y += sy
		x = (x + sx) % maxx
		row := data[y]

		if row[x] == "#" {
			t++
		}
	}

	return
}

func part1() int {
	return traverse(3, 1)
}

func part2() (x int) {
	routes := []struct {
		x int
		y int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	x = 1
	for _, r := range routes {
		x = x * traverse(r.x, r.y)
	}

	return
}

func main() {
	p1 := part1()
	p2 := part2()
	fmt.Println("Part1: ", p1)
	fmt.Println("Part2: ", p2)
}
