package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pt1 := part1()
	pt2 := part2()

	fmt.Println("Part 1: ", pt1)
	fmt.Println("Part 2: ", pt2)
}

func part1() int {
	i := make(chan string, 4096)
	go streamInput(i)
	chains := parseInput(i)
	count := countPart1(chains)
	return count
}

func part2() int {
	i := make(chan string, 4096)
	go streamInput(i)
	chains := parseInput(i)
	count := countPart2(chains, "shiny gold", 0)
	return int(count)
}

func streamInput(i chan<- string) {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic("No file!")
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		i <- s.Text()
	}

	close(i)

	if s.Err() != nil {
		panic("File parse error!")
	}
}

type bagChk struct {
	col string
	qty uint
}

type bagChain []bagChk

func (bc bagChain) contains(s string) bool {
	for _, bag := range bc {
		if bag.col == s {
			return true
		}
	}
	return false
}

func parseInput(i <-chan string) map[string]bagChain {
	chains := make(map[string]bagChain)

	for s := range i {
		words := strings.Split(s, " ")

		parent := strings.Join(words[:2], " ")

		children := make([]bagChk, 0)
		for n := 4; n <= len(words)-3; n = n + 4 {
			qty, _ := strconv.ParseUint(words[n], 0, 0)
			if qty == 0 {
				continue
			}
			col := strings.Join(words[n+1:n+3], " ")
			bc := bagChk{col, uint(qty)}
			children = append(children, bc)
		}

		chains[parent] = children
	}

	return chains
}

func countPart1(chains map[string]bagChain) (count int) {
	counts := make(map[string]uint)
	toCheck := []string{"shiny gold"}
	for len(toCheck) != 0 {
		nextToCheck := make([]string, 0)
		for parent, bags := range chains {
			for _, col := range toCheck {
				if bags.contains(col) {
					counts[parent]++
					nextToCheck = append(nextToCheck, parent)
				}
			}
		}
		toCheck = nextToCheck
	}
	return len(counts)
}

func countPart2(chains map[string]bagChain, col string, lvl uint) uint {
	count := uint(0)

	bag := chains[col]
	for _, iBag := range bag {
		count += iBag.qty * countPart2(chains, iBag.col, lvl+1)
	}

	if lvl == 0 {
		return count // Don't add on the top level container bag
	}

	return count + 1 // Deeper level, so add on the container bag as well
}
