package main

import (
	"fmt"
	"strconv"

	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("input.txt")

	fmt.Println("Part 1:", pt1)
}

func part1(p string) int {
	i := make(chan string)
	go util.StreamInput(i, p)
	cardPubKey, doorPubKey := getKeys(i)
	cardLs, doorLs := calcLs(cardPubKey), calcLs(doorPubKey)
	encKeyA := calcEncKey(cardPubKey, doorLs)
	encKeyB := calcEncKey(doorPubKey, cardLs)
	if encKeyA != encKeyB {
		panic("bad enc key")
	}
	return encKeyA
}

func getKeys(i <-chan string) (int, int) {
	s1, s2 := <-i, <-i

	cardPubKey, _ := strconv.Atoi(s1)
	doorPubKey, _ := strconv.Atoi(s2)

	return int(cardPubKey), int(doorPubKey)
}

func calcLs(n int) int {
	val := 1
	ls := 0
	for val != n {
		val = val * 7
		val = val % 20201227
		ls++
	}
	return ls
}

func calcEncKey(key, ls int) int {
	val := 1
	for i := 0; i < ls; i++ {
		val = val * key
		val = val % 20201227
	}
	return val
}
