package main

import (
	"flag"
	"os"
)

func neighborsCount(grid Grid, cell Coords) (count uint8) {
	for x := cell.X - 1; x <= cell.X+1; x++ {
		for y := cell.Y - 1; y <= cell.Y+1; y++ {
			if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[x]) &&
				grid[x][y] && !(x == cell.X && y == cell.Y) {

				count++
			}
		}
	}
	return
}

func doStep(grid Grid) Grid {
	newGrid := grid.EmptyCopy()

	for x := range grid {
		for y := range grid[x] {
			neighbors := neighborsCount(grid, Coords{x, y})
			if (!grid[x][y] && neighbors == 3) ||
				(grid[x][y] && (neighbors == 2 || neighbors == 3)) {

				newGrid[x][y] = true
			}
		}
	}

	return newGrid
}

func main() {
	seedFileName := flag.String("seed-file", "seed.json", "seed file in json format")
	flag.Parse()

	seedFile, err := os.Open(*seedFileName)
	if err != nil {
		panic(err)
	}

	grid, err := NewGridFromJSON(seedFile)
	if err != nil {
		panic(err)
	}
	grid = grid.Trim()
	grid = grid.Pad(1)
	grid.PrettyPrint(os.Stdout)

	for range 1 {
		grid = doStep(grid)
		grid = grid.Trim()
		grid = grid.Pad(1)
		grid.PrettyPrint(os.Stdout)
	}
	grid = grid.Trim()
	grid.PrettyPrint(os.Stdout)

}
