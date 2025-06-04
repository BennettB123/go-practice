package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Conways Game of Life")
	ebiten.SetTPS(60)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
