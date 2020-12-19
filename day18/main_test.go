package main

import "testing"

func Test_part1(t *testing.T) {
	exp := 45283905029161
	got := part1("./input.txt")
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}

func Test_part2(t *testing.T) {
	exp := 216975281211165
	got := part2("./input.txt")
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}
