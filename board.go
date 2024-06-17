package main

import (
	"errors"
	"fmt"
)

// like set in the mathematical sense
type Set[T comparable] map[T]struct{}

var exists = struct{}{}

var OutOfBounds = errors.New("Out of bounds.")

var Bounds = struct{ Width, Height int }{8, 8}

type Point struct {
	X, Y int
}

type Piece interface {
	PossibleMoves(Point) []Point
	fmt.Stringer
}

type Board interface {
	At(Point) Piece
	Move(p1, p2 Point) error
	Set(Point, Piece) error
}

type BitMap uint64

func Points(b BitMap) Set[Point] {
	points := Set[Point]{}
	for y := 0; y < Bounds.Height; y++ {
		for x := 0; x < Bounds.Width; x++ {
			if 1 & (b << (y*(Bounds.Width-1) + x)) == 1 {
				points[Point{x, y}] = exists
			}
		}
	}
	return points
}

type BitMapBoard map[Piece]BitMap

func IsOutOfBounds(pt Point) bool {
	return pt.X >= Bounds.Width || 
		pt.Y >= Bounds.Height
}

func (b BitMapBoard) At(pt Point) ([]Piece, error) {
	if IsOutOfBounds(pt) {
		return nil, OutOfBounds
	}

	pieces := make([]Piece, 0, 2)
	for k, v := range b {
		m := Points(v)
		if _, ok := m[pt]; ok {
			pieces = append(pieces, k)
		}
	}
	return pieces, nil
}

func (b BitMapBoard) Move(p Piece, pt1, pt2 Point) error {
	if IsOutOfBounds(pt1) || IsOutOfBounds(pt2) {
		return OutOfBounds
	}

	return nil
}

func (b BitMapBoard) Set(p Piece, pt Point) error {
	if IsOutOfBounds(pt) {
		return OutOfBounds
	}

	return nil
}
