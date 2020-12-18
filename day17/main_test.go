package main

import "testing"

func Test_part1(t *testing.T) {
	exp := 112
	got := part1("./input_test.txt")
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}

// func Test_part2(t *testing.T) {
// 	exp := 51240700105297
// 	got := part2("./input.txt")
// 	if exp != got {
// 		t.Errorf("Part 1: expected %d, got %d", exp, got)
// 	}
// }
