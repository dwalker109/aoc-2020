package main

import "fmt"

type recitation map[uint]num

func (r *recitation) remember(k uint, v uint) {
	lookup, ok := (*r)[k]
	if ok {
		(*r)[k] = append(lookup, v)
	} else {
		(*r)[k] = num{v}
	}
}

type num []uint

func (n num) turnsApart() uint {
	if len(n) < 2 {
		return 0
	}
	return n[len(n)-1] - n[len(n)-2]
}

func main() {
	pt1 := solve([]uint{1, 0, 16, 5, 17, 4}, 2020)
	pt2 := solve([]uint{1, 0, 16, 5, 17, 4}, 30000000)

	fmt.Println("Part 1:", pt1)
	fmt.Println("Part 2:", pt2)
}

func solve(s []uint, t uint) int {
	var lastNum uint

	rec := make(recitation)
	for i, n := range s {
		lastNum = n
		rec.remember(n, uint(i+1))
	}

	for i := uint(len(rec) + 1); i <= t; i++ {
		lastHist := rec[lastNum]
		lastNum = lastHist.turnsApart()
		rec.remember(lastNum, i)
	}

	return int(lastNum)
}
