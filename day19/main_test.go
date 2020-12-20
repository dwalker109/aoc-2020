package main

import "testing"

func Test_part1(t *testing.T) {
	exp := 136
	got := solve("./input.txt")
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}

func Test_part2(t *testing.T) {
	exp := 136
	got := solve("./input_part2.txt")
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}
