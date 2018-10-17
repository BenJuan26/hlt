package hlt

import "fmt"

// Position - Location on the game map
type Position struct {
	x int
	y int
}

// Equals compares two positions and returns true if they are equal
func (p *Position) Equals(other *Position) bool {
	return p.x == other.GetX() && p.y == other.GetY()
}

func (p *Position) String() string {
	return fmt.Sprintf("Pos{x=%d,y=%d}", p.x, p.y)
}

// DirectionalOffset - Returns the position of a move in the direction
func (p *Position) DirectionalOffset(d *Direction) (*Position, error) {
	switch d.charValue {
	case NORTH:
		return &Position{p.x, p.y - 1}, nil
	case SOUTH:
		return &Position{p.x, p.y + 1}, nil
	case EAST:
		return &Position{p.x + 1, p.y}, nil
	case WEST:
		return &Position{p.x - 1, p.y}, nil
	case STILL:
		return &Position{p.x, p.y}, nil
	}
	return nil, fmt.Errorf("Invalid direction %c", d.charValue)
}

// GetX returns the x value of the position
func (p *Position) GetX() int {
	return p.x
}

// GetY returns the y value of the position
func (p *Position) GetY() int {
	return p.y
}
