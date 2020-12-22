package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("./input.txt")

	fmt.Println(pt1)
}

func part1(p string) int {
	i := make(chan string)
	p1, p2 := getDecks(i, p)
	winDeck := play(&p1, &p2)

	return winDeck.score()
}

func getDecks(i chan string, p string) (p1, p2 deck) {
	go util.StreamInputCustomSplit(i, p, util.SplitOnString("\n\n"))
	d1 := <-i
	s1 := strings.SplitN(d1, ":", 2)
	p1.id = s1[0]
	c1 := strings.Fields(s1[1])
	for _, x := range c1 {
		n, _ := strconv.ParseUint(x, 10, 64)
		p1.cards = append(p1.cards, n)
	}
	d2 := <-i
	s2 := strings.SplitN(d2, ":", 2)
	p2.id = s2[0]
	c2 := strings.Fields(s2[1])
	for _, x := range c2 {
		n, _ := strconv.ParseUint(x, 10, 64)
		p2.cards = append(p2.cards, n)
	}

	return
}

func play(p1, p2 *deck) *deck {
	for {
		c1, c2 := p1.draw(), p2.draw()
		if c1 > c2 {
			p1.confiscate(c1, c2)
		}
		if c2 > c1 {
			p2.confiscate(c2, c1)
		}
		if len(p1.cards) == 0 || len(p2.cards) == 0 {
			break
		}
	}

	if len(p1.cards) > 1 {
		return p1
	}
	return p2
}

type deck struct {
	id    string
	cards []uint64
}

func (d *deck) draw() uint64 {
	c := d.cards[0]
	d.cards = d.cards[1:]

	return c
}

func (d *deck) confiscate(c ...uint64) {
	d.cards = append(d.cards, c...)
}

func (d *deck) score() (n int) {
	for i := 1; i <= len(d.cards); i++ {
		n += (i * int(d.cards[len(d.cards)-i]))
	}
	return
}
