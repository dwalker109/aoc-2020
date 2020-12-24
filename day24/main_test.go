package main

import "testing"

func Test_part1(t *testing.T) {
	exp := 10
	got := part1("./input_test.txt")
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}

// func Test_part2(t *testing.T) {
// 	exp := 149245887792
// 	got := part2(inputTest)
// 	if exp != got {
// 		t.Errorf("Part 1: expected %d, got %d", exp, got)
// 	}
// }
