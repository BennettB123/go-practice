package main

import "math/rand/v2"

const (
	Dead  = iota
	Alive = iota
)

type Grid struct {
	width  int
	height int
	cells  [][]int
}

func NewGrid(width, height int) Grid {
	cells := make([][]int, height)
	for i := range height {
		cells[i] = make([]int, width)
	}

	return Grid{
		width,
		height,
		cells,
	}
}

func (grid *Grid) Tick() {
	// Rules (from https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)
	// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	// Any live cell with two or three live neighbours lives on to the next generation.
	// Any live cell with more than three live neighbours dies, as if by overpopulation.
	// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

	// TODO: re-use an array to avoid allocating this every Tick
	nextGen := make([][]int, grid.height)
	for i := range grid.height {
		nextGen[i] = make([]int, grid.width)
	}

	for y := range grid.height {
		for x := range grid.width {
			aliveNeighbors := 0

			// count neighbors
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}

					if grid.Get(x+j, y+i) == Alive {
						aliveNeighbors++
					}
				}
			}

			// handle live cell
			if grid.Get(x, y) == Alive {
				if aliveNeighbors < 2 { // death by underpopulation
					nextGen[y][x] = Dead
				} else if aliveNeighbors < 4 { // alive!
					nextGen[y][x] = Alive
				} else { // death by overpopulation
					nextGen[y][x] = Dead
				}

			} else { // handle dead cell
				if aliveNeighbors == 3 {
					nextGen[y][x] = Alive
				} else {
					nextGen[y][x] = Dead
				}
			}
		}
	}

	grid.cells = nextGen
}

// Get returns the value of the cell at [x, y].
// If [x, y] is outside the bounds of the grid, returns Dead
func (g *Grid) Get(x, y int) int {
	if g.outOfBounds(x, y) {
		return Dead
	}

	return g.cells[y][x]
}

func (g *Grid) Set(x, y, val int) {
	if g.outOfBounds(x, y) {
		return
	}

	g.cells[y][x] = val
}

func (g Grid) outOfBounds(x, y int) bool {
	return x < 0 || x >= g.width || y < 0 || y >= g.height
}

func (g *Grid) Randomize() {
	for y := range g.height {
		for x := range g.width {
			if rand.Float32() < .25 {
				g.Set(x, y, Alive)
			} else {
				g.Set(x, y, Dead)
			}
		}
	}
}

func (g *Grid) Clear() {
	for y := range g.height {
		for x := range g.width {
			g.Set(x, y, Dead)
		}
	}
}
