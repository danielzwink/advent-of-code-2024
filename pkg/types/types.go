package types

import (
	"advent-of-code-2024/pkg/util"
	"fmt"
	"strings"
)

type Direction string

const (
	Up        Direction = "Up"
	Right     Direction = "Right"
	Down      Direction = "Down"
	Left      Direction = "Left"
	UpLeft    Direction = "UpLeft"
	UpRight   Direction = "UpRight"
	DownLeft  Direction = "DownLeft"
	DownRight Direction = "DownRight"
)

var Directions = map[Direction]*Coordinate{
	Up:        {0, -1},
	Right:     {1, 0},
	Down:      {0, 1},
	Left:      {-1, 0},
	UpLeft:    {-1, -1},
	UpRight:   {1, -1},
	DownLeft:  {-1, 1},
	DownRight: {1, 1},
}

var DirectionSigns = map[rune]*Coordinate{
	'^': {0, -1},
	'>': {1, 0},
	'v': {0, 1},
	'<': {-1, 0},
}

type Coordinate struct {
	X, Y int
}

func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{X: x, Y: y}
}

func NewCoordinateFromCsv(value string) *Coordinate {
	split := strings.Split(value, ",")
	return NewCoordinate(util.MustParseInt(split[0]), util.MustParseInt(split[1]))
}

func (c *Coordinate) Add(o *Coordinate) *Coordinate {
	return NewCoordinate(c.X+o.X, c.Y+o.Y)
}

func (c *Coordinate) Diff(o *Coordinate) *Coordinate {
	return NewCoordinate(c.X-o.X, c.Y-o.Y)
}

func (c *Coordinate) Within(b *Coordinate) bool {
	return c.X >= 0 && c.X < b.X && c.Y >= 0 && c.Y < b.Y
}

func (c *Coordinate) WithinBounds(lower *Coordinate, upper *Coordinate) bool {
	return c.X >= lower.X && c.X < upper.X && c.Y >= lower.Y && c.Y < upper.Y
}

func (c *Coordinate) OutOf(b *Coordinate) bool {
	return c.X < 0 || c.X >= b.X || c.Y < 0 || c.Y >= b.Y
}

func (c *Coordinate) Key() string {
	return fmt.Sprintf("%02d %02d", c.X, c.Y)
}

func (c *Coordinate) ManhattanDistance(o *Coordinate) int {
	return util.Abs(c.X-o.X) + util.Abs(c.Y-o.Y)
}
