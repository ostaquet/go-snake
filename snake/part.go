package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
)

var AssetHead *ebiten.Image
var AssetBody *ebiten.Image

type Part struct {
	X, Y   int // X & Y position in the grid
	visual *ebiten.Image
	board  *Board
}

type PartType int

const (
	Head PartType = iota
	Body
)

func NewPart(X, Y int, partType PartType, board *Board) *Part {
	if AssetHead == nil {
		AssetHead = LoadImage("Snake.png").SubImage(image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{10, 10},
		}).(*ebiten.Image)
		AssetBody = LoadImage("Snake.png").SubImage(image.Rectangle{
			Min: image.Point{10, 0},
			Max: image.Point{20, 10},
		}).(*ebiten.Image)
	}

	var visual *ebiten.Image

	switch partType {
	case Head:
		visual = AssetHead
	case Body:
		visual = AssetBody
	}

	part := Part{
		X:      X,
		Y:      Y,
		visual: visual,
		board:  board,
	}
	return &part
}

func (p *Part) Draw(screen *ebiten.Image, direction Direction) {
	drawOptSquare := ebiten.DrawImageOptions{}

	drawOptSquare.GeoM.Translate(-float64(10)/2, -float64(10)/2)
	switch direction {
	case Up:
		drawOptSquare.GeoM.Rotate(float64(270%360) * 2 * math.Pi / 360)
	case Right:
		drawOptSquare.GeoM.Rotate(float64(0%360) * 2 * math.Pi / 360)
	case Down:
		drawOptSquare.GeoM.Rotate(float64(90%360) * 2 * math.Pi / 360)
	case Left:
		drawOptSquare.GeoM.Rotate(float64(180%360) * 2 * math.Pi / 360)
	}
	drawOptSquare.GeoM.Translate(+float64(10)/2, +float64(10)/2)
	drawOptSquare.GeoM.Translate(float64(p.X)*float64(p.visual.Bounds().Dx())+p.board.X(), float64(p.Y)*float64(p.visual.Bounds().Dy())+p.board.Y())

	screen.DrawImage(p.visual, &drawOptSquare)
}
