package main

import "testing"

func Test_part1(t *testing.T) {
	exp := 2
	got := solve("./input_test.txt")
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}

func Test_part2(t *testing.T) {
	exp := 12
	got := solve("./input_test_part2.txt")
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}
