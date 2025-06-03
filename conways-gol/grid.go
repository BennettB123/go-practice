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

func (g *Grid) Tick() {

}

func (g *Grid) Get(x, y int) int {
	return g.cells[y][x]
}

func (g *Grid) Set(x, y, val int) {
	g.cells[y][x] = val
}

func (g *Grid) Randomize(chance float32) {
	for y := range g.height {
		for x := range g.width {
			if rand.Float32() < chance {
				g.Set(x, y, Alive)
			} else {
				g.Set(x, y, Dead)
			}
		}
	}
}
