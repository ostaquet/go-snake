package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Part struct {
	X, Y   int
	visual *ebiten.Image
}

var blue = color.RGBA{
	R: 0,
	G: 0,
	B: 255,
	A: 255,
}

func NewPart(X, Y, size int) *Part {
	square := ebiten.NewImage(size, size)
	square.Fill(blue)

	part := Part{
		X:      X,
		Y:      Y,
		visual: square,
	}
	return &part
}

func (p *Part) Draw(screen *ebiten.Image) {
	drawOptSquare := ebiten.DrawImageOptions{}
	drawOptSquare.GeoM.Translate(float64(p.X)*float64(p.visual.Bounds().Dx()), float64(p.Y)*float64(p.visual.Bounds().Dy()))
	screen.DrawImage(p.visual, &drawOptSquare)
}
