package learn

import (
	"errors"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Point struct {
	X int
	Y int
}

func (p *Point) Move(di Direction, distance int) error {
	switch di {
	case UP:
		p.Y -= distance

	case DOWN:
		p.Y += distance

	case LEFT:
		p.X -= distance

	case RIGHT:
		p.X += distance

	default:
		return errors.New("invalid direction value")

	}
	return nil

}
