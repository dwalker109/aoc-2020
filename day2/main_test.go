package main

import "testing"

func Test_part1(t *testing.T) {
	exp := 477
	got := part1(getPasswords())

	if got != exp {
		t.Errorf("Answer wrong, expected %d and received %d", exp, got)
	}
}

func Test_part2(t *testing.T) {
	exp := 686
	got := part2(getPasswords())

	if got != exp {
		t.Errorf("Answer wrong, expected %d and received %d", exp, got)
	}
}
