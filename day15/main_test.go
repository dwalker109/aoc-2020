package main

import "testing"

func Test_part1(t *testing.T) {
	exp := 438
	got := solve([]uint{3, 2, 1}, uint(2020))
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}

func Test_part2(t *testing.T) {
	exp := 18
	got := solve([]uint{3, 2, 1}, uint(30000000))
	if exp != got {
		t.Errorf("Part 1: expected %d, got %d", exp, got)
	}
}

func Benchmark_part1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solve([]uint{3, 2, 1}, uint(2020))
	}
}
