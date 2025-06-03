package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Images struct {
	aliveCell *ebiten.Image
}

func NewImages(width, height int) Images {
	aliveCell := ebiten.NewImage(width, height)
	aliveCell.Fill(color.RGBA{0, 255, 0, 255})

	return Images{
		aliveCell,
	}
}
