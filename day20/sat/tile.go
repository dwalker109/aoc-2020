package sat

import (
	"strconv"
	"strings"

	"github.com/dwalker109/aoc-2020/util"
)

const (
	units = 10
)

type Tile struct {
	ID  int
	raw []byte
	img [][]byte
}

// NewTile inits a new tile from a string - MUST BE 10x10
func NewTile(s string) *Tile {
	dat := (strings.ReplaceAll(s, "\n", ""))
	id, raw := dat[5:9], dat[10:]
	img := make([][]byte, 0, units)
	for i := 0; i < len(raw); i += units {
		img = append(img, []byte(raw[i:i+units]))
	}

	ID, _ := strconv.Atoi(id)
	t := Tile{ID, []byte(raw), img}

	return &t
}

func (t *Tile) Top() []byte {
	return t.img[0]
}

func (t *Tile) TopRev() []byte {
	return util.ReverseBytes(t.Top())
}

func (t *Tile) Rgt() []byte {
	e := make([]byte, units)
	for i, row := range t.img {
		e[i] = row[units-1]
	}
	return e
}

func (t *Tile) RgtRev() []byte {
	return util.ReverseBytes(t.Rgt())
}

func (t *Tile) Btm() []byte {
	return t.img[units-1]
}

func (t *Tile) BtmRev() []byte {
	return util.ReverseBytes(t.Btm())
}

func (t *Tile) Lft() []byte {
	e := make([]byte, units)
	for i, row := range t.img {
		e[i] = row[0]
	}
	return e
}

func (t *Tile) LftRev() []byte {
	return util.ReverseBytes(t.Lft())
}

func (t *Tile) Edges() [8][]byte {
	e := [8][]byte{
		t.Top(),
		t.TopRev(),
		t.Rgt(),
		t.RgtRev(),
		t.Btm(),
		t.BtmRev(),
		t.Lft(),
		t.LftRev(),
	}
	return e
}

func (t *Tile) rotate() {
	m := make([][]byte, len(t.img))
	for i := 0; i < len(m); i++ {
		m[i] = make([]byte, len(t.img[0]))
	}

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			m[x][len(m)-1-y] = t.img[y][x]
		}
	}
	t.img = m
}

func (t *Tile) flip() {
	m := make([][]byte, len(t.img))
	for i := 0; i < len(m); i++ {
		m[i] = make([]byte, len(t.img[0]))
	}

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			m[len(m)-1-y][x] = t.img[y][x]
		}
	}
	t.img = m
}
