package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const ScreenWidth int = 1000
const ScreenHeight int = 1000
const GridWidth int = 50
const GridHeight int = 50
const CellWidth int = ScreenWidth / GridWidth
const CellHeight int = ScreenHeight / GridHeight

type Game struct {
	gui           *Gui
	grid          Grid
	images        Images
	gridTickAccum float64 // accumulator for grid tick timing
}

func NewGame() *Game {
	grid := NewGrid(GridWidth, GridHeight)
	grid.Randomize()

	return &Game{
		gui:    NewGui(),
		grid:   grid,
		images: NewImages(CellWidth, CellHeight),
	}
}

func (game *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	// determine if we need to update grid cells based on updatesPerSec
	dt := 1.0 / ebiten.ActualTPS()
	game.gridTickAccum += dt * float64(*game.gui.updatesPerSec)
	for game.gridTickAccum >= 1.0 {
		game.grid.Tick()
		game.gridTickAccum -= 1.0
	}

	game.handleKeys()

	game.gui.Update()

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.drawGrid(screen)
	game.drawGridLines(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(ebiten.ActualFPS())))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %d", int(ebiten.ActualTPS())), 0, 20)

	game.gui.Draw(screen)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (game *Game) handleKeys() {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		game.grid.Randomize()
	}
}

func (game *Game) drawGridLines(screen *ebiten.Image) {
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

func (game *Game) drawGrid(screen *ebiten.Image) {
	opt := ebiten.DrawImageOptions{}
	for y := range GridHeight {
		for x := range GridWidth {
			opt.GeoM.Reset()
			opt.GeoM.Translate(float64(x*CellWidth), float64(y*CellHeight))

			if game.grid.Get(x, y) == Alive {
				screen.DrawImage(game.images.aliveCell, &opt)
			} else {
				screen.DrawImage(game.images.deadCell, &opt)
			}
		}
	}
}
