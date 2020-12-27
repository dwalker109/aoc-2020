package sat

import (
	"bytes"
	"regexp"
	"strings"
)

type Img struct {
	GridSize         int
	Parts, Available map[int]*Tile
	Matched          [][]*Tile
	Actual           *Tile
}

func NewImg(tiles []*Tile) *Img {
	sz := 1
	for {
		if len(tiles) == sz*sz {
			break
		}
		sz++
	}

	parts := make(map[int]*Tile)
	available := make(map[int]*Tile)
	assembled := make([][]*Tile, sz)

	for _, t := range tiles {
		parts[t.ID] = t
		available[t.ID] = t
	}

	return &Img{sz, parts, available, assembled, nil}
}

func (i *Img) Assemble() {
	// Setup rows/cols
	for r := 0; r < len(i.Matched); r++ {
		i.Matched[r] = make([]*Tile, i.GridSize)
	}

	// Find and orient top left piece
CornerMatch:
	for {
		for _, t := range i.Available {
			for p := 0; p <= 8; p++ {
				mt := i.findMatchingEdge(t.ID, t.top())
				ml := i.findMatchingEdge(t.ID, t.lft())
				if mt == nil && ml == nil {
					i.Matched[0][0] = t
					delete(i.Available, t.ID)
					break CornerMatch
				}
				if p == 3 {
					t.flip()
				} else {
					t.rotate()
				}
			}
		}
	}

	// Find and orient LHS pieces
LHSMatch:
	for r := 1; r < i.GridSize; r++ {
		for _, t := range i.Available {
			e := i.Matched[r-1][0].btm()
			m := i.findMatchingEdge(-1, e)
			for p := 0; p <= 8; p++ {
				if bytes.Equal(m.top(), e) {
					i.Matched[r][0] = m
					delete(i.Available, m.ID)
					continue LHSMatch
				}
				if p == 3 {
					t.flip()
				} else {
					t.rotate()
				}
			}
		}
	}

	// Find and orient remaining rows
	for r := 0; r < i.GridSize; r++ {
	RestMatch:
		for c := 1; c < i.GridSize; c++ {
			for _, t := range i.Available {
				e := i.Matched[r][c-1].rgt()
				m := i.findMatchingEdge(-1, e)
				for p := 0; p <= 8; p++ {
					if bytes.Equal(m.lft(), e) {
						i.Matched[r][c] = m
						delete(i.Available, m.ID)
						continue RestMatch
					}
					if p == 3 {
						t.flip()
					} else {
						t.rotate()
					}
				}
			}
		}
	}
}

func (i *Img) findMatchingEdge(exclude int, eA []byte) *Tile {
	for ID, t := range i.Available {
		if ID == exclude {
			continue
		}

		for _, eB := range t.edges() {
			if bytes.Equal(eA, eB) {
				return t
			}
		}
	}

	return nil
}

func (i *Img) BuildActual() {
	dat := new(strings.Builder)
	for _, tileRow := range i.Matched {
		for dataRow := 1; dataRow < defaultPxGrid-1; dataRow++ {
			for _, t := range tileRow {
				dat.Write(t.Img[dataRow][1 : defaultPxGrid-1])
			}
		}
	}

	imgPxGrid := (len(i.Matched[0][0].Img[0]) - 2) * len(i.Matched)
	i.Actual = NewActualImageTile(dat.String(), imgPxGrid)
}

func (i *Img) SurfaceSeaMonsters() {
	hereBeMonsters := false

	rx := [3]*regexp.Regexp{
		regexp.MustCompile(`.{18}(#{1}).{1}`),
		regexp.MustCompile(`(#{1}).{4}(#{2}).{4}(#{2}).{4}(#{3})`),
		regexp.MustCompile(`.{1}(#{1}).{2}(#{1}).{2}(#{1}).{2}(#{1}).{2}(#{1}).{2}(#{1}).{3}`),
	}

	w := 20
	h := 3
	for p := 0; p < 8; p++ {
		for r := 0; r < len(i.Actual.Img)-h; r++ {
			for c := 0; c < len(i.Actual.Img[0])-w; c++ {
				lx := [][]byte{
					i.Actual.Img[r][c : c+w],
					i.Actual.Img[r+1][c : c+w],
					i.Actual.Img[r+2][c : c+w],
				}

				mx := [3][]int{
					rx[0].FindSubmatchIndex(lx[0]),
					rx[1].FindSubmatchIndex(lx[1]),
					rx[2].FindSubmatchIndex(lx[2]),
				}

				if len(mx[0]) > 0 && len(mx[1]) > 0 && len(mx[2]) > 0 {
					hereBeMonsters = true

					for rIdx := 0; rIdx < 3; rIdx++ {
						for mIdx := 2; mIdx < len(mx[rIdx]); mIdx += 2 {
							starts := mx[rIdx][mIdx]
							ends := mx[rIdx][mIdx+1]
							for cIdx := starts; cIdx < ends; cIdx++ {
								i.Actual.Img[r+rIdx][c+cIdx] = 79
							}
						}
					}
				}
			}
		}

		if !hereBeMonsters {
			if p == 3 {
				i.Actual.flip()
			} else {
				i.Actual.rotate()
			}
		}
	}

	sb := new(strings.Builder)
	for _, b := range i.Actual.Img {
		sb.Write(b)
	}
	i.Actual.Raw = []byte(sb.String())
}

func (i *Img) CalcWaterRoughness() int {
	r := regexp.MustCompile(`#`)
	m := r.FindAllStringIndex(string(i.Actual.Raw), -1)
	return len(m)
}
