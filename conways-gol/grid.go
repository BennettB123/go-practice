package main

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
