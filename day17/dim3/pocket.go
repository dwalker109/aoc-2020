package dim3

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/dwalker109/aoc-2020/util"
)

const (
	active   = "#"
	inactive = "."
)

func NewDimension(state []string) *Dimension {
	dim := &Dimension{
		Cycle:    0,
		GridSize: len(state),
		Planes:   make(map[int]*Plane),
	}
	dim.AddPlane(0)

	var lookupStates = map[string]bool{
		active:   true,
		inactive: false,
	}

	for i, s := range state {
		for j, c := range s {
			dim.Planes[0].Cubes[i][j].State.EndActive = lookupStates[string(c)]
			dim.Planes[0].Cubes[i][j].State.IniActive = lookupStates[string(c)]
		}
	}

	return dim
}

type Dimension struct {
	Cycle    int
	PlaneQty int
	GridSize int
	Planes   map[int]*Plane
}

func (d *Dimension) AddPlane(z int) {
	cubes := make([][]*Cube, d.GridSize)
	for i := 0; i < d.GridSize; i++ {
		row := make([]*Cube, d.GridSize)
		cubes[i] = row
		for j := 0; j < d.GridSize; j++ {
			cubes[i][j] = NewCube(j, i, z)
		}
	}

	d.Planes[z] = &Plane{z, cubes}
	d.PlaneQty++
}

func (d *Dimension) GrowPlanes() {
	zBase := (len(d.Planes) - 1) / 2
	zBelow, zAbove := -(zBase + 1), zBase+1
	d.AddPlane(zBelow)
	d.AddPlane(zAbove)
}

func (d *Dimension) Expand() {
	d.GridSize += 2

	for z, plane := range d.Planes {
		newStartRow := make([]*Cube, d.GridSize)
		for x := 0; x < d.GridSize; x++ {
			newStartRow[x] = NewCube(x, 0, z)
		}
		newFinalRow := make([]*Cube, d.GridSize)
		for x := 0; x < d.GridSize; x++ {
			newFinalRow[x] = NewCube(x, d.GridSize-1, z)
		}

		for y, row := range plane.Cubes {
			newStartCol := NewCube(0, y+1, z)
			newFinalCol := NewCube(d.GridSize-1, y+1, z)
			for _, c := range row {
				c.X++
				c.Y++
			}

			tmp := append([]*Cube{newStartCol}, row...)
			plane.Cubes[y] = append(tmp, newFinalCol)
		}

		tmp := append([][]*Cube{newStartRow}, plane.Cubes...)
		plane.Cubes = append(tmp, newFinalRow)
	}
}

func (d *Dimension) Prepare() {
	for _, plane := range d.Planes {
		for _, row := range plane.Cubes {
			for _, c := range row {
				c.State.IniActive = c.State.EndActive
			}
		}
	}
}

func (d *Dimension) Simulate() {
	d.GrowPlanes()
	d.Expand()
	d.Prepare()
	for _, plane := range d.Planes {
		for _, cubes := range plane.Cubes {
			for _, cube := range cubes {
				cube.UpdateState(d)
			}
		}
	}
	d.Cycle++
}

func (d *Dimension) CountEndActive() (n int) {
	for _, plane := range d.Planes {
		for _, row := range plane.Cubes {
			for _, c := range row {
				if c.State.EndActive {
					n++
				}
			}
		}
	}
	return
}

func (d *Dimension) DebugPrint() {
	fmt.Println("++++++++", "CYCLE", d.Cycle, "++++++++")
	for _, plane := range d.Planes {
		fmt.Println("~~~~~~~~", plane.Z, "~~~~~~~~")
		for _, row := range plane.Cubes {
			for _, c := range row {
				if c.State.EndActive {
					fmt.Print(active)
				} else {
					fmt.Print(inactive)
				}
			}
			fmt.Print("\n")
		}
	}
}

type Plane struct {
	Z     int
	Cubes [][]*Cube
}

type Cube struct {
	// Active bool
	State struct {
		IniActive bool
		EndActive bool
	}
	X int
	Y int
	Z int
}

func NewCube(x, y, z int) *Cube {
	c := &Cube{
		State: struct {
			IniActive bool
			EndActive bool
		}{false, false},
		X: x,
		Y: y,
		Z: z,
	}

	return c
}

func (c *Cube) IniActive() string {
	if c.State.IniActive {
		return active
	}
	return inactive
}

func (c *Cube) UpdateState(d *Dimension) {
	sx := util.MaxInt(c.X-1, 0)
	ex := util.MinInt(c.X+1, d.GridSize-1)
	sy := util.MaxInt(c.Y-1, 0)
	ey := util.MinInt(c.Y+1, d.GridSize-1)
	sz := util.MaxInt(c.Z-1, -((d.PlaneQty - 1) / 2))
	ez := util.MinInt(c.Z+1, (d.PlaneQty-1)/2)

	var sb strings.Builder
	for z := sz; z <= ez; z++ {
		for y := sy; y <= ey; y++ {
			for x := sx; x <= ex; x++ {
				if c.X == x && c.Y == y && c.Z == z {
					sb.WriteString("+")
				} else {
					sb.WriteString(d.Planes[z].Cubes[y][x].IniActive())
				}
			}
		}
	}

	check := sb.String()
	matches := regexp.MustCompile(active).FindAllStringIndex(check, -1)

	if c.State.IniActive {
		if len(matches) == 2 || len(matches) == 3 {
			c.State.EndActive = true
		} else {
			c.State.EndActive = false

		}
	} else {
		if len(matches) == 3 {
			c.State.EndActive = true
		} else {
			c.State.EndActive = false
		}
	}
}
