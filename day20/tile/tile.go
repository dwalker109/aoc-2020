package tile

import (
	"fmt"
	"strings"
)

type Tiles map[string]tile

type tile struct {
	ID  string
	Raw string
	Img [][]byte
	Top []byte
	Btm []byte
	Lft []byte
	Rgt []byte
}

// NewTile inits a new tile from a string - MUST BE 10x10
func NewTile(s string) tile {
	dat := (strings.ReplaceAll(s, "\n", ""))
	id, raw := dat[5:9], dat[10:]
	img := make([][]byte, 0, 10)
	for i := 0; i < len(raw); i += 10 {
		img = append(img, []byte(raw[i:i+10]))
	}

	t := tile{
		ID:  id,
		Raw: raw,
		Img: img,
	}

	t.cacheEdges()

	return t
}

func (t *tile) cacheEdges() {
	t.Top = t.Img[0]
	t.Btm = t.Img[len(t.Img)-1]
	var lft, rgt []byte
	for _, row := range t.Img {
		lft = append(lft, row[0])
		rgt = append(rgt, row[len(row)-1])
	}
	t.Lft = lft
	t.Rgt = rgt
}

func (t *tile) Print() {
	fmt.Println(t.Id)
	for _, l := range t.Img {
		fmt.Println(string(l))
	}
	fmt.Println()
}

func (t *tile) rotate() {
	m := make([][]byte, len(t.Img))
	for i := 0; i < len(m); i++ {
		m[i] = make([]byte, len(t.Img[0]))
	}

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			m[x][len(m)-1-y] = t.Img[y][x]
		}
	}
}

func (t *tile) flip() {
	m := make([][]byte, len(t.Img))
	for i := 0; i < len(m); i++ {
		m[i] = make([]byte, len(t.Img[0]))
	}

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			m[len(m)-1-y][x] = t.Img[y][x]
		}
	}
}
