package main

import (
	"bufio"
	"bytes"
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
	i := getInput()
	ui := uniqueInput(i)
	for _, s := range ui {
		sum += uint(len(s))
	}
	return
}
func part2() (sum uint) {
	i := getInput()
	ui := concensusInput(i)
	for _, s := range ui {
		sum += uint(len(s))
	}
	return
}

func uniqueInput(i [][]string) (u []string) {
	for _, groupAnswers := range i {
		uniqueAnswers := make(map[string]bool)
		for _, ans := range strings.Join(groupAnswers, "") {
			uniqueAnswers[string(ans)] = true
		}
		var b bytes.Buffer
		for ans := range uniqueAnswers {
			b.WriteString(ans)
		}
		u = append(u, b.String())
	}
	return
}

func concensusInput(i [][]string) (u []string) {
	for _, groupAnswers := range i {
		answersCount := make(map[string]uint)
		for _, ans := range groupAnswers {
			for _, c := range ans {
				curr := answersCount[string(c)]
				answersCount[string(c)] = curr + 1
			}
		}
		var b bytes.Buffer
		for ans, count := range answersCount {
			if int(count) == len(groupAnswers) {
				b.WriteString(ans)
			}
		}
		u = append(u, b.String())
	}
	return
}

func getInput() (i [][]string) {
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
		all := strings.Fields(s.Text())
		i = append(i, all)
	}

	if s.Err() != nil {
		panic("File parse error!")
	}

	return
}
