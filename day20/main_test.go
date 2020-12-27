package main

import "testing"

func Test_part1(t *testing.T) {
	exp := 20899048083289
	got := part1("./input_test.txt")
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}

func Test_part2(t *testing.T) {
	exp := 273
	got := part2("./input_test.txt")
	if exp != got {
		t.Errorf("Part 2: expected %d, got %d", exp, got)
	}
}
