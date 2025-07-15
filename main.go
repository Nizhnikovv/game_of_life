package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type Coords struct {
	X, Y int
}

func ParseJSONSeed(seed io.Reader) ([][]bool, error) {
	var coords []Coords
	if err := json.NewDecoder(seed).Decode(&coords); err != nil {
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

	grid := make([][]bool, xLen)
	for i := range grid {
		grid[i] = make([]bool, yLen)
	}

	for _, coord := range coords {
		relX := coord.X - baseCoords.X
		relY := coord.Y - baseCoords.Y

		grid[relX][relY] = true
	}

	return grid, nil
}

func main() {
	println("hello world")
}
