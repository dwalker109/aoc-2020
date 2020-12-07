package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	p1 := part1()
	p2 := part2()

	fmt.Println("Part 1: ", p1)
	fmt.Println("Part 2: ", p2)
}

func part1() (sum uint) {
	i := make(chan []string, 4096)
	pi := make(chan answers, 4096)
	go streamInput(i)
	go processInput(i, pi)
	for a := range pi {
		sum += a.unique
	}
	return
}

func part2() (sum uint) {
	i := make(chan []string, 4096)
	pi := make(chan answers, 4096)
	go streamInput(i)
	go processInput(i, pi)
	for a := range pi {
		sum += a.concencus
	}
	return
}

func streamInput(i chan<- []string) {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic("No file!")
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	onBlankLine := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data)-1; i++ {
			if data[i] == '\n' && data[i+1] == '\n' {
				return i + 2, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	}
	s.Split(onBlankLine)
	for s.Scan() {
		i <- strings.Fields(s.Text())
	}

	close(i)

	if s.Err() != nil {
		panic("File parse error!")
	}
}

type answers struct {
	data      []string
	unique    uint
	concencus uint
}

func processInput(i <-chan []string, ci chan<- answers) {
	for groupAnswers := range i {
		answersCount := make(map[string]uint)

		for _, ans := range groupAnswers {
			for _, c := range ans {
				curr := answersCount[string(c)]
				answersCount[string(c)] = curr + 1
			}
		}

		unique, concensus := uint(len(answersCount)), uint(0)
		for _, count := range answersCount {
			if int(count) == len(groupAnswers) {
				concensus++
			}
		}

		ci <- answers{
			groupAnswers,
			unique,
			concensus,
		}
	}
	close(ci)
}
