package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// Board The board is a simple rectangle in which the snake will evolve
type Board struct {
	x1, y1 float64
	x2, y2 float64
	lines  []*ebiten.Image
}

var white = color.RGBA{
	R: 255,
	G: 255,
	B: 255,
	A: 255,
}

func NewBoard(X1, Y1, X2, Y2 float64) *Board {
	b := &Board{
		x1: X1,
		y1: Y1,
		x2: X2,
		y2: Y2,
	}

	b.lines = make([]*ebiten.Image, 4)

	b.lines[0] = ebiten.NewImage(int(b.x2-b.x1)+2, 1)
	b.lines[0].Fill(white)
	b.lines[1] = ebiten.NewImage(int(b.x2-b.x1)+2, 1)
	b.lines[1].Fill(white)
	b.lines[2] = ebiten.NewImage(1, int(b.y2-b.y1)+2+1)
	b.lines[2].Fill(white)
	b.lines[3] = ebiten.NewImage(1, int(b.y2-b.y1)+2+1)
	b.lines[3].Fill(white)

	return b
}

func (b *Board) Draw(screen *ebiten.Image) {
	op1 := ebiten.DrawImageOptions{}
	op1.GeoM.Translate(b.x1-1, b.y1-1)
	screen.DrawImage(b.lines[0], &op1)
	screen.DrawImage(b.lines[2], &op1)

	op2 := ebiten.DrawImageOptions{}
	op2.GeoM.Translate(b.x1-1, b.y2+1)
	screen.DrawImage(b.lines[1], &op2)

	op3 := ebiten.DrawImageOptions{}
	op3.GeoM.Translate(b.x2+1, b.y1-1)
	screen.DrawImage(b.lines[3], &op3)
}

func (b *Board) X() float64 {
	return b.x1
}

func (b *Board) Y() float64 {
	return b.y1
}

func (b *Board) Width() int {
	return int(b.x2 - b.x1)
}

func (b *Board) Height() int {
	return int(b.y2 - b.y1)
}
