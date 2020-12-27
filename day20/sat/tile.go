package sat

import (
	"strconv"
	"strings"

	"github.com/dwalker109/aoc-2020/util"
)

const (
	defaultPxGrid = 10
)

type Tile struct {
	ID  int
	Raw []byte
	Img [][]byte
}

// NewGarbledTile inits a new tile from a string - MUST BE 10x10
func NewGarbledTile(s string) *Tile {
	dat := (strings.ReplaceAll(s, "\n", ""))
	id, raw := dat[5:9], dat[10:]
	img := make([][]byte, 0, defaultPxGrid)
	for i := 0; i < len(raw); i += defaultPxGrid {
		img = append(img, []byte(raw[i:i+defaultPxGrid]))
	}

	ID, _ := strconv.Atoi(id)
	t := Tile{ID, []byte(raw), img}

	return &t
}

func NewActualImageTile(s string, pxGrid int) *Tile {
	img := make([][]byte, 0, pxGrid)
	for i := 0; i < len(s); i += pxGrid {
		img = append(img, []byte(s[i:i+pxGrid]))
	}

	t := Tile{0, []byte(s), img}

	return &t
}

func (t *Tile) top() []byte {
	return t.Img[0]
}

func (t *Tile) topRev() []byte {
	return util.ReverseBytes(t.top())
}

func (t *Tile) rgt() []byte {
	e := make([]byte, defaultPxGrid)
	for i, row := range t.Img {
		e[i] = row[defaultPxGrid-1]
	}
	return e
}

func (t *Tile) rgtRev() []byte {
	return util.ReverseBytes(t.rgt())
}

func (t *Tile) btm() []byte {
	return t.Img[defaultPxGrid-1]
}

func (t *Tile) btmRev() []byte {
	return util.ReverseBytes(t.btm())
}

func (t *Tile) lft() []byte {
	e := make([]byte, defaultPxGrid)
	for i, row := range t.Img {
		e[i] = row[0]
	}
	return e
}

func (t *Tile) lftRev() []byte {
	return util.ReverseBytes(t.lft())
}

func (t *Tile) edges() [8][]byte {
	e := [8][]byte{
		t.top(),
		t.topRev(),
		t.rgt(),
		t.rgtRev(),
		t.btm(),
		t.btmRev(),
		t.lft(),
		t.lftRev(),
	}
	return e
}

func (t *Tile) rotate() {
	m := make([][]byte, len(t.Img))
	for i := 0; i < len(m); i++ {
		m[i] = make([]byte, len(t.Img[0]))
	}

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			m[x][len(m)-1-y] = t.Img[y][x]
		}
	}
	t.Img = m
}

func (t *Tile) flip() {
	m := make([][]byte, len(t.Img))
	for i := 0; i < len(m); i++ {
		m[i] = make([]byte, len(t.Img[0]))
	}

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			m[len(m)-1-y][x] = t.Img[y][x]
		}
	}
	t.Img = m
}
