package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
)

var AssetHead *ebiten.Image
var AssetBody *ebiten.Image
var AssetTail *ebiten.Image

type Part struct {
	X, Y   int // X & Y position in the grid
	visual *ebiten.Image
	board  *Board
}

type PartType int

const (
	Head PartType = iota
	Body
	Tail
)

func NewPart(X, Y int, board *Board) *Part {
	if AssetHead == nil {
		AssetHead = LoadImage("Snake.png").SubImage(image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{10, 10},
		}).(*ebiten.Image)
		AssetBody = LoadImage("Snake.png").SubImage(image.Rectangle{
			Min: image.Point{10, 0},
			Max: image.Point{20, 10},
		}).(*ebiten.Image)
		AssetTail = LoadImage("Snake.png").SubImage(image.Rectangle{
			Min: image.Point{10, 10},
			Max: image.Point{20, 20},
		}).(*ebiten.Image)
	}

	part := Part{
		X:      X,
		Y:      Y,
		visual: AssetBody,
		board:  board,
	}
	return &part
}

func (p *Part) Draw(screen *ebiten.Image, direction Direction, partType PartType) {
	switch partType {
	case Head:
		p.visual = AssetHead
	case Tail:
		p.visual = AssetTail
	case Body:
		p.visual = AssetBody
	}

	op := ebiten.DrawImageOptions{}

	op.GeoM.Translate(-float64(10)/2, -float64(10)/2)
	switch direction {
	case Up:
		op.GeoM.Rotate(float64(270%360) * 2 * math.Pi / 360)
	case Right:
		op.GeoM.Rotate(float64(0%360) * 2 * math.Pi / 360)
	case Down:
		op.GeoM.Rotate(float64(90%360) * 2 * math.Pi / 360)
	case Left:
		op.GeoM.Rotate(float64(180%360) * 2 * math.Pi / 360)
	}
	op.GeoM.Translate(+float64(10)/2, +float64(10)/2)
	op.GeoM.Translate(float64(p.X)*float64(p.visual.Bounds().Dx())+p.board.X(), float64(p.Y)*float64(p.visual.Bounds().Dy())+p.board.Y())

	screen.DrawImage(p.visual, &op)
}
