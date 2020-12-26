package sat

import "bytes"

type Img struct {
	GridSize         int
	Parts, Available map[int]*Tile
	Assembled        [][]*Tile
}

func NewImg(tiles []*Tile) *Img {
	sz := 2
	for {
		if len(tiles)%sz == 0 {
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

	return &Img{sz, parts, available, assembled}
}

func (i *Img) Assemble() {
	// Setup rows/cols
	for r := 0; r < len(i.Assembled); r++ {
		i.Assembled[r] = make([]*Tile, i.GridSize)
	}

	// Find and orient top left piece
CornerMatch:
	for {
		for _, t := range i.Available {
			for p := 0; p <= 8; p++ {
				mt := i.FindMatchingEdge(t.ID, t.Top())
				ml := i.FindMatchingEdge(t.ID, t.Lft())
				if mt == nil && ml == nil {
					i.Assembled[0][0] = t
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
			e := i.Assembled[r-1][0].Btm()
			m := i.FindMatchingEdge(-1, e)
			for p := 0; p <= 8; p++ {
				if bytes.Equal(m.Top(), e) {
					i.Assembled[r][0] = m
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
				e := i.Assembled[r][c-1].Rgt()
				m := i.FindMatchingEdge(-1, e)
				for p := 0; p <= 8; p++ {
					if bytes.Equal(m.Lft(), e) {
						i.Assembled[r][c] = m
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

func (i *Img) FindMatchingEdge(exclude int, e []byte) *Tile {
	for ID, t := range i.Available {
		if ID == exclude {
			continue
		}

		if bytes.Equal(t.Top(), e) ||
			bytes.Equal(t.TopRev(), e) ||
			bytes.Equal(t.Rgt(), e) ||
			bytes.Equal(t.RgtRev(), e) ||
			bytes.Equal(t.Btm(), e) ||
			bytes.Equal(t.BtmRev(), e) ||
			bytes.Equal(t.Lft(), e) ||
			bytes.Equal(t.LftRev(), e) {
			return t
		}
	}

	return nil
}
