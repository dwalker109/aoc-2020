package main

import "testing"

func Test_part1(t *testing.T) {
	exp := 190
	got := part1(getInput())
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}

func Test_part2(t *testing.T) {
	exp := 121
	got := part2(getInput())
	if exp != got {
		t.Errorf("Part 2: expected %d, got %d", exp, got)
	}
}
