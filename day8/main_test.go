package main

import "testing"

func Test_part1(t *testing.T) {
	exp := 1563
	got := part1()
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}

func Test_part2(t *testing.T) {
	exp := 767
	got := part2()
	if exp != got {
		t.Errorf("Part 2: expected %d, got %d", exp, got)
	}
}
