package ship

import "math"

type WaypointShip struct {
	x  int
	y  int
	wp struct {
		x int
		y int
	}
}

func NewWaypointShip(x, y, wpX, wpY int) *WaypointShip {
	return &WaypointShip{x, y, struct{ x, y int }{wpX, wpY}}
}

func (s *WaypointShip) MoveWaypoint(h int, d int) {
	switch h {
	case 0:
		s.wp.y += d
	case 90:
		s.wp.x += d
	case 180:
		s.wp.y -= d
	case 270:
		s.wp.x -= d
	default:
		panic("only NSEW cardinal dirs supported")
	}
}

func (s *WaypointShip) MoveWaypointDir(o string, d int) {
	switch o {
	case "N":
		s.MoveWaypoint(0, d)
	case "E":
		s.MoveWaypoint(90, d)
	case "S":
		s.MoveWaypoint(180, d)
	case "W":
		s.MoveWaypoint(270, d)
	default:
		panic("only NSEW cardinal dirs supported")
	}
}

func (s *WaypointShip) Forward(q int) {
	s.x += s.wp.x * q
	s.y += s.wp.y * q
}

func (s *WaypointShip) Left(d int) {
	rad := float64(d) * (math.Pi / 180.0)
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	x := float64(s.wp.x)*cos - float64(s.wp.y)*sin
	y := float64(s.wp.x)*sin + float64(s.wp.y)*cos
	s.wp.x = int(math.Round(x))
	s.wp.y = int(math.Round(y))
}

func (s *WaypointShip) Right(d int) {
	rad := float64(d) * (math.Pi / 180.0)
	x := float64(s.wp.x)*math.Cos(rad) + float64(s.wp.y)*math.Sin(rad)
	y := float64(-s.wp.x)*math.Sin(rad) + float64(s.wp.y)*math.Cos(rad)
	s.wp.x = int(math.Round(x))
	s.wp.y = int(math.Round(y))
}

func (s *WaypointShip) Manhattan() int {
	m := math.Abs(float64(s.x)) + math.Abs(float64(s.y))
	return int(m)
}
