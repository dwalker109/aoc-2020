package main

import "testing"

func Test_part1(t *testing.T) {
	exp := 16003257187056
	got := part1("./input.txt")
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}

func Test_part2(t *testing.T) {
	exp := 3219837697833
	got := part2("./input.txt")
	if exp != got {
		t.Errorf("Part 2: expected %d, got %d", exp, got)
	}
}
