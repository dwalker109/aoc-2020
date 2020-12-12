package ship

import "math"

type SimpleShip struct {
	x       int
	y       int
	heading int
}

func NewSimpleShip(x, y, heading int) *SimpleShip {
	return &SimpleShip{x, y, heading}
}

func (s *SimpleShip) Move(h int, d int) {
	switch h {
	case 0:
		s.y += d
	case 90:
		s.x += d
	case 180:
		s.y -= d
	case 270:
		s.x -= d
	default:
		panic("only NSEW cardinal dirs supported")
	}
}

func (s *SimpleShip) MoveDir(o string, d int) {
	switch o {
	case "N":
		s.Move(0, d)
	case "E":
		s.Move(90, d)
	case "S":
		s.Move(180, d)
	case "W":
		s.Move(270, d)
	default:
		panic("only NSEW cardinal dirs supported")
	}
}

func (s *SimpleShip) Forward(d int) {
	s.Move(s.heading, d)
}

func (s *SimpleShip) Left(d int) {
	x := s.heading - d
	if x < 0 {
		s.heading = 360 + x
	} else {
		s.heading = x
	}
}

func (s *SimpleShip) Right(d int) {
	s.heading = (s.heading + d) % 360
}

func (s *SimpleShip) Manhattan() int {
	m := math.Abs(float64(s.x)) + math.Abs(float64(s.y))
	return int(m)
}
