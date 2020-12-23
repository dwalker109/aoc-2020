package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"

	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("./input.txt")
	pt2 := part2("./input.txt")

	fmt.Println(pt1)
	fmt.Println(pt2)
}

func part1(p string) int {
	i := make(chan string)
	p1, p2 := getDecks(i, p)
	winner := combat(&p1, &p2)

	return winner.score()
}

func part2(p string) int {
	i := make(chan string)
	p1, p2 := getDecks(i, p)
	winDeck := recursiveCombat(&p1, &p2)

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

func combat(p1, p2 *deck) *deck {
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

	if len(p1.cards) > 0 {
		return p1
	}
	return p2
}

func recursiveCombat(p1, p2 *deck) *deck {
	infLog := make(map[string]interface{})

	for {
		infHash := p1.hash() + p2.hash()
		if _, exists := infLog[infHash]; exists {
			return p1
		}
		infLog[infHash] = new(interface{})

		c1, c2 := p1.draw(), p2.draw()

		// Play a subgame if necessary
		if len(p1.cards) >= int(c1) && len(p2.cards) >= int(c2) {
			p1sub := p1.clone(c1)
			p2sub := p2.clone(c2)
			winner := recursiveCombat(&p1sub, &p2sub)

			if winner.id == "Player 1" {
				p1.confiscate(c1, c2)
			} else {
				p2.confiscate(c2, c1)
			}

			continue
		}

		// No subgame, direct compare
		if c1 > c2 {
			p1.confiscate(c1, c2)
		} else {
			p2.confiscate(c2, c1)
		}

		if len(p1.cards) == 0 || len(p2.cards) == 0 {
			break
		}
	}

	if len(p1.cards) > 0 {
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

func (d *deck) hash() string {
	s := fmt.Sprint(d.cards)
	h := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", h)
}

func (d *deck) clone(n uint64) deck {
	c := make([]uint64, n)
	if int(n) == len(d.cards) {
		copy(c, d.cards)
	} else {
		copy(c, d.cards[:n+1])
	}
	return deck{d.id, c}
}
