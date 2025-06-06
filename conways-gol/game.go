package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const ScreenWidth int = 1250
const ScreenHeight int = 1250
const GridWidth int = 100
const GridHeight int = 100

var CellWidth float64 = float64(ScreenWidth) / float64(GridWidth)
var CellHeight float64 = float64(ScreenHeight) / float64(GridHeight)

type Game struct {
	gui           *Gui
	grid          Grid
	images        Images
	gridTickAccum float64 // accumulator for grid tick timing
	paused        bool
}

func NewGame() *Game {
	grid := NewGrid(GridWidth, GridHeight)
	grid.Randomize()

	return &Game{
		gui:           NewGui(),
		grid:          grid,
		images:        NewImages(int(CellWidth), int(CellHeight)),
		gridTickAccum: 0,
		paused:        false,
	}
}

func (game *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	// determine if we need to update grid cells based on updatesPerSec
	if !game.paused {
		dt := 1.0 / ebiten.ActualTPS()
		game.gridTickAccum += dt * float64(*game.gui.updatesPerSec)
		for game.gridTickAccum >= 1.0 {
			game.grid.Tick()
			game.gridTickAccum -= 1.0
		}
	}

	game.handleKeys()
	game.handleMouse()

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

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		game.paused = !game.paused
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		game.grid.Clear()
	}
}

func (game *Game) handleMouse() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		screenX, screenY := ebiten.CursorPosition()
		gridX, gridY := convertToGridCoords(screenX, screenY)
		game.grid.Set(gridX, gridY, Alive)
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		screenX, screenY := ebiten.CursorPosition()
		gridX, gridY := convertToGridCoords(screenX, screenY)
		game.grid.Set(gridX, gridY, Dead)
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
			opt.GeoM.Translate(float64(float64(x)*CellWidth), float64(float64(y)*CellHeight))

			if game.grid.Get(x, y) == Alive {
				screen.DrawImage(game.images.aliveCell, &opt)
			} else {
				screen.DrawImage(game.images.deadCell, &opt)
			}
		}
	}
}

// Returns the coordinates of the grid that coorespond to [screenX, screenY].
// If [screenX, screenY] is off the screen, returns (-1, -1)
func convertToGridCoords(screenX, screenY int) (x, y int) {
	if (screenX < 0 || screenX > ScreenWidth) ||
		(screenY < 0 || screenY > ScreenHeight) {
		return -1, -1
	}

	x = int((float32(screenX) / float32(ScreenWidth)) * float32(GridWidth))
	y = int((float32(screenY) / float32(ScreenHeight)) * float32(GridHeight))

	return
}
