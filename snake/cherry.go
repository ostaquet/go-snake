package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

var red = color.RGBA{
	R: 255,
	G: 0,
	B: 0,
	A: 255,
}

type Cherry struct {
	X, Y, size int
	visual     *ebiten.Image
}

func NewCherry(X, Y, size int) *Cherry {
	square := ebiten.NewImage(size, size)
	square.Fill(red)

	cherry := Cherry{
		X:      X,
		Y:      Y,
		visual: square,
	}
	return &cherry
}

func (c *Cherry) Draw(screen *ebiten.Image) {
	drawOptSquare := ebiten.DrawImageOptions{}
	drawOptSquare.GeoM.Translate(float64(c.X)*float64(c.visual.Bounds().Dx()), float64(c.Y)*float64(c.visual.Bounds().Dy()))
	screen.DrawImage(c.visual, &drawOptSquare)
}
