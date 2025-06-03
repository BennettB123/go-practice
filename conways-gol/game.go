package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const ScreenWidth int = 1000
const ScreenHeight int = 1000
const GridWidth int = 50
const GridHeight int = 50

type Game struct {
	grid   Grid
	images Images
}

func NewGame() *Game {
	return &Game{
		grid:   NewGrid(GridWidth, GridHeight),
		images: NewImages(ScreenWidth/GridWidth, ScreenHeight/GridHeight),
	}
}

func (g *Game) Update() error {
	g.grid.Tick()

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.images.aliveCell, nil)

	g.drawGridLines(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) drawGridLines(screen *ebiten.Image) {
	// draw grid lines
	lightGray := color.RGBA{50, 50, 50, 255}
	for i := 1; i < GridHeight; i++ {
		row := (float32(screen.Bounds().Size().Y) / float32(GridHeight)) * float32(i)
		vector.StrokeLine(screen, 0, row, float32(screen.Bounds().Size().X), row, 1, lightGray, false)
	}
	for i := 1; i < GridWidth; i++ {
		col := (float32(screen.Bounds().Size().X) / float32(GridWidth)) * float32(i)
		vector.StrokeLine(screen, col, 0, col, float32(screen.Bounds().Size().Y), 1, lightGray, false)
	}
}
