package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("./input.txt")
	pt2 := part2("./input.txt")

	fmt.Println("Part 1: ", pt1)
	fmt.Println("Part 2: ", pt2)
}

var (
	searchBytes []byte = []byte("mask = ")
	searchLen   int    = len(searchBytes)
)

func splitInput(data []byte, atEOF bool) (advance int, token []byte, err error) {
	dataLen := len(data)

	if atEOF && dataLen == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, searchBytes); i >= 0 {
		return i + searchLen, data[0:i], nil
	}

	if atEOF {
		return dataLen, data, nil
	}

	// Moar data
	return 0, nil, nil
}

func part1(p string) int {
	i := make(chan string, 128)
	go util.StreamInputCustomSplit(i, p, splitInput)
	init := initFerryV1(i)
	return int(init)
}

func part2(p string) int {
	i := make(chan string, 128)
	go util.StreamInputCustomSplit(i, p, splitInput)
	init := initFerryV2(i)
	return int(init)
}

func initFerryV1(i <-chan string) int {
	mem := make(map[uint64]uint64)

	for s := range i {
		l := strings.Split(s, "\n")
		strMask := l[0]
		orMask, _ := strconv.ParseUint(strings.ReplaceAll(strMask, "X", "0"), 2, 64)
		andMask, _ := strconv.ParseUint(strings.ReplaceAll(strMask, "X", "1"), 2, 64)

		for _, ops := range l[1:] {
			if ops == "" {
				continue
			}

			txt := regexp.MustCompile("[0-9]+").FindAllString(ops, -1)
			addr, _ := strconv.ParseUint(txt[0], 10, 64)
			val, _ := strconv.ParseUint(txt[1], 10, 64)
			mVal := (val | orMask) & andMask

			mem[addr] = mVal
		}
	}

	sum := uint64(0)
	for _, v := range mem {
		sum += v
	}
	return int(sum)
}

func initFerryV2(i <-chan string) int {
	mem := make(map[uint64]uint64)

	for s := range i {
		l := strings.Split(s, "\n")
		strMask := l[0]

		for _, ops := range l[1:] {
			if ops == "" {
				continue
			}

			txt := regexp.MustCompile("[0-9]+").FindAllString(ops, -1)
			addrInt, _ := strconv.ParseUint(txt[0], 10, 64)
			val, _ := strconv.ParseUint(txt[1], 10, 64)

			addrBytes := []byte(fmt.Sprintf("%036b", addrInt))
			for i, c := range strMask {
				if string(c) == "1" || string(c) == "X" {
					addrBytes[i] = byte(c)
				}
			}
			addrStr := string(addrBytes)

			permutations := make([][]int, 0)
			n := len(regexp.MustCompile("X").FindAllStringIndex(addrStr, -1))
			k := (1 << n)
			for i := 0; i < k; i++ {
				permutations = append(permutations, make([]int, n))

				for j := 0; j < n; j++ {
					permutations[i][j] = (i >> j) & 1
				}
			}

			for _, i := range permutations {
				build := addrStr
				for _, j := range i {
					build = strings.Replace(build, "X", strconv.FormatInt(int64(j), 2), 1)
				}
				decodedDec, _ := strconv.ParseUint(build, 2, 64)
				mem[decodedDec] = val
			}
		}
	}

	sum := uint64(0)
	for _, v := range mem {
		sum += v
	}
	return int(sum)
}
