package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Coords struct {
	X, Y int
}

type Grid [][]bool

var (
	ErrInvalidGridDimensions = errors.New("invalid grid dimensions")
)

func NewEmptyGrid(x, y int) (Grid, error) {
	if x < 1 || y < 1 {
		return nil, ErrInvalidGridDimensions
	}

	new := make([][]bool, x)
	for i := range new {
		new[i] = make([]bool, y)
	}
	return new, nil
}

func NewGridFromJSON(r io.Reader) (Grid, error) {
	var coords []Coords
	if err := json.NewDecoder(r).Decode(&coords); err != nil {
		return nil, fmt.Errorf("error when decoding json: %w", err)
	}

	var baseCoords, maxCoords Coords
	for _, coord := range coords {
		baseCoords.X = min(baseCoords.X, coord.X)
		baseCoords.Y = min(baseCoords.Y, coord.Y)

		maxCoords.X = max(maxCoords.X, coord.X)
		maxCoords.Y = max(maxCoords.Y, coord.Y)
	}

	xLen := maxCoords.X - baseCoords.X
	yLen := maxCoords.Y - baseCoords.Y

	grid := make([][]bool, yLen+1)
	for i := range grid {
		grid[i] = make([]bool, xLen+1)
	}

	for _, coord := range coords {
		relX := coord.X - baseCoords.X
		relY := coord.Y - baseCoords.Y

		grid[relY][relX] = true
	}

	return grid, nil
}

// returns new Grid
func (g Grid) Pad(padding int) Grid {
	padded := make([][]bool, len(g)+padding*2)
	for i := range padded {
		padded[i] = make([]bool, len(g[0])+padding*2)
	}

	for x := range g {
		copy(padded[x+padding][padding:len(padded[x])-padding], g[x])
	}

	return padded
}

func (g Grid) EmptyCopy() Grid {
	new := make([][]bool, len(g))
	for i := range new {
		new[i] = make([]bool, len(g[i]))
	}
	return new
}

// works with underlying array but changes length
func (g Grid) Trim() Grid {
	var minX, maxX, minY, maxY int
	for x := range g {
		for y := range g[x] {
			if g[x][y] {
				minX = min(minX, x)
				maxX = max(maxX, x)
				minY = min(minY, y)
				maxY = max(maxY, y)
			}
		}
	}
	g = g[minX : maxX+1]
	for i := range g {
		g[i] = g[i][minY : maxY+1]
	}

	return g
}

const (
	liveCell = "██"
	deadCell = "  "
)

func (g Grid) PrettyPrint(w io.Writer) {
	for i := len(g) - 1; i >= 0; i-- {
		for _, cell := range g[i] {
			if cell {
				fmt.Fprint(w, liveCell)
			} else {
				fmt.Fprint(w, deadCell)
			}
		}
		fmt.Fprint(w, "\n")
	}
}
