package hexx

type Floor map[[2]int]*Tile

func (f Floor) CountBlack() int {
	n := 0
	for _, t := range f {
		if t.col.curr == "b" {
			n++
		}
	}
	return n
}

func (f *Floor) CommitAll() {
	for _, t := range *f {
		t.Commit()
	}
}
