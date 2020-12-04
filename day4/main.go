package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	input := getInput()
	p1 := part1(input)
	p2 := part2(input)

	fmt.Println("Part 1: ", p1)
	fmt.Println("Part 2: ", p2)
}

func part1(i []map[string]string) (v int) {
	v = 0
	for _, s := range i {
		if isValidSeq(s) {
			v++
		}
	}
	return
}

func part2(i []map[string]string) (v int) {
	v = 0
	for _, s := range i {
		if isValidData(s) {
			v++
		}
	}
	return
}

func getInput() (i []map[string]string) {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic("No file!")
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	onBlankLine := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data)-1; i++ {
			if data[i] == '\n' && data[i+1] == '\n' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	}
	s.Split(onBlankLine)
	for s.Scan() {
		entry := strings.ReplaceAll(s.Text(), "\n", " ")
		items := strings.Fields(entry)
		mapped := make(map[string]string, 0)
		for _, v := range items {
			pair := strings.Split(v, ":")
			k, v := string(pair[0]), string(pair[1])
			mapped[k] = v
		}
		i = append(i, mapped)
	}

	if s.Err() != nil {
		panic("File parse error!")
	}

	return
}

func isValidSeq(s map[string]string) bool {
	req := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	for _, field := range req {
		if _, ok := s[field]; ok == false {
			return false
		}

	}

	return true
}

func isValidData(s map[string]string) bool {
	if isValidSeq(s) == false {
		return false
	}

	rules := map[string]*regexp.Regexp{
		"byr": regexp.MustCompile("^(19[2-9]\\d|200[0-2])$"),
		"iyr": regexp.MustCompile("^(201[0-9]|2020)$"),
		"eyr": regexp.MustCompile("^(202[0-9]|2030)$"),
		"hgt": regexp.MustCompile("^((1[5-8][0-9])cm|(19[0-3])cm|59in|(6[0-9])in|(7[0-6])in)$"),
		"hcl": regexp.MustCompile("^(#[0-9a-f]{6})$"),
		"ecl": regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$"),
		"pid": regexp.MustCompile("^(\\d{9})$"),
		"cid": regexp.MustCompile(".*"),
	}

	for k, v := range s {
		res := rules[k].FindAllStringIndex(v, -1)
		if len(res) == 0 {
			return false
		}
	}

	return true
}
