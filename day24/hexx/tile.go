package hexx

func NewTile(c [2]int) *Tile {
	return &Tile{c, struct {
		pend string
		curr string
	}{"w", "w"}}
}

type Tile struct {
	pos [2]int
	col struct {
		pend string
		curr string
	}
}

func (t *Tile) IsBlack() bool {
	return t.col.curr == "b"
}
func (t *Tile) IsWhite() bool {
	return t.col.curr == "w"
}

func (t *Tile) Flip() {
	if t.col.curr == "b" {
		t.col.pend = "w"
	} else {
		t.col.pend = "b"
	}
}

func (t *Tile) Commit() {
	t.col.curr = t.col.pend
}

func (t *Tile) AdjacentCoords() [6][2]int {
	qr := t.pos
	adj := [6][2]int{
		{qr[0] + 1, qr[1]},
		{qr[0], qr[1] + 1},
		{qr[0] - 1, qr[1] + 1},
		{qr[0] - 1, qr[1]},
		{qr[0], qr[1] - 1},
		{qr[0] + 1, qr[1] - 1},
	}
	return adj
}
