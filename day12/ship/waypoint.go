package ship

import "math"

type waypointShip struct {
	x  int
	y  int
	wp struct {
		x int
		y int
	}
}

func NewWaypointShip(x, y, wpX, wpY int) *waypointShip {
	return &waypointShip{x, y, struct{ x, y int }{wpX, wpY}}
}

func (s *waypointShip) Move(o string, d int) {
	switch o {
	case "N":
		s.wp.y += d
	case "E":
		s.wp.x += d
	case "S":
		s.wp.y -= d
	case "W":
		s.wp.x -= d
	default:
		panic("only NSEW cardinal dirs supported")
	}
}

func (s *waypointShip) Forward(q int) {
	s.x += s.wp.x * q
	s.y += s.wp.y * q
}

func (s *waypointShip) Left(d int) {
	rad := float64(d) * (math.Pi / 180.0)
	cos := math.Cos(rad)
	sin := math.Sin(rad)

	x := float64(s.wp.x)*cos - float64(s.wp.y)*sin
	y := float64(s.wp.x)*sin + float64(s.wp.y)*cos

	s.wp.x = int(math.Round(x))
	s.wp.y = int(math.Round(y))
}

func (s *waypointShip) Right(d int) {
	rad := float64(d) * (math.Pi / 180.0)
	cos := math.Cos(rad)
	sin := math.Sin(rad)

	x := float64(s.wp.x)*cos + float64(s.wp.y)*sin
	y := float64(-s.wp.x)*sin + float64(s.wp.y)*cos

	s.wp.x = int(math.Round(x))
	s.wp.y = int(math.Round(y))
}

func (s *waypointShip) Manhattan() int {
	m := math.Abs(float64(s.x)) + math.Abs(float64(s.y))
	return int(m)
}
