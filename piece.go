package piece

import (
	"errors"
)

var (
	InvalidMove = errors.New("invalid move")
)

type Color int

// Color at the same time act as
// direction to move for pawns.
const (
	White Color = 1
	Black Color = -1
)

type Point struct {
	X, Y int
}

type Piece interface {
	PossibleMoves() []Point
	// Moves Piece on Board.
	Move(x, y int) error
}

// nil represent empty cell...
type Cell Piece

type Board [][]Cell

// example piece 
type Pawn struct {
	Color
	Pos Point
	// is its first move?
	first bool
}

func (p *Pawn) PossibleMoves() (coords []Point) {
	coords = append(coords, Point{p.Pos.X, p.Pos.Y + p.Color})
	if p.first {
		c := coords[0]
		coords = append(coords, Point{c.X, c.Y + p.Color})
		p.first = false
	}
	return
}

func NewPawn(c Color, p Pos) Pawn {
	return Pawn{
		Color: c,
		Pos:   p,
		first: true,
	}
}

func (p Pawn) Move(x, y int) error {
	for _, v := range p.PossibleMoves() {
		if (Point{x, y}) != v {
			return InvalidMove
		}
	}
	
	return nil
}
