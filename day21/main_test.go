package main

import "testing"

func Test_part1(t *testing.T) {
	exp := 2485
	got := part1("./input.txt")
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}

func Test_part2(t *testing.T) {
	exp := "bqkndvb,zmb,bmrmhm,snhrpv,vflms,bqtvr,qzkjrtl,rkkrx"
	got := part2("./input.txt")
	if exp != got {
		t.Errorf("Part 1: expected %s, got %s", exp, got)
	}
}
