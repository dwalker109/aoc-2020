package main

import "testing"

func Test_part1(t *testing.T) {
	r := make(chan int, 1)
	go part1(r)

	exp := 270144
	got := <-r

	if got != exp {
		t.Errorf("Answer wrong, expected %d and received %d", exp, got)
	}
}

func Test_part2(t *testing.T) {
	r := make(chan int, 1)
	go part2(r)

	exp := 261342720
	res := <-r

	if res != exp {
		t.Errorf("Answer wrong, expected %d and received %d", exp, res)
	}
}

