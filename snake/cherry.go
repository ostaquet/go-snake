package snake

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
	"os"
)

var AssetCherry *ebiten.Image

type Cherry struct {
	X, Y, size int // X & Y position in the grid
	visual     *ebiten.Image
	board      *Board
}

func NewCherry(X, Y int, board *Board) *Cherry {
	if AssetCherry == nil {
		// Load the assets
		dat, err := os.ReadFile("snake/assets/images/Cherry.png")
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(bytes.NewReader(dat))
		if err != nil {
			log.Fatal(err)
		}
		AssetCherry = ebiten.NewImageFromImage(img)
	}

	cherry := Cherry{
		X:      X,
		Y:      Y,
		visual: AssetCherry,
		board:  board,
	}
	return &cherry
}

func (c *Cherry) Draw(screen *ebiten.Image) {
	drawOptSquare := ebiten.DrawImageOptions{}
	drawOptSquare.GeoM.Translate(float64(c.X)*float64(c.visual.Bounds().Dx())+c.board.X(), float64(c.Y)*float64(c.visual.Bounds().Dy())+c.board.Y())
	screen.DrawImage(c.visual, &drawOptSquare)
}
