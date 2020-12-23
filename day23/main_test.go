package main

import "testing"

func Test_part1(t *testing.T) {
	exp := "67384529"
	got := part1(inputTest)
	if exp != got {
		t.Errorf("Part 1: expected %s, got %s", exp, got)
	}
}

// func Test_part2(t *testing.T) {
// 	exp := 291
// 	got := part2("./input_test.txt")
// 	if exp != got {
// 		t.Errorf("Part 1: expected %d, got %d", exp, got)
// 	}
// }