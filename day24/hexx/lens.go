package hexx

type Lens [2]int

func (l *Lens) Move(d string) {
	switch d {
	case "e":
		l[0]++
	case "se":
		l[1]++
	case "sw":
		l[0]--
		l[1]++
	case "w":
		l[0]--
	case "nw":
		l[1]--
	case "ne":
		l[0]++
		l[1]--
	default:
		panic("bad dir")
	}
}
