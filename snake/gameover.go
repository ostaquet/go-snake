package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"strconv"
)

type GameOver struct {
	mplusNormalFont font.Face
	mplusBigFont    font.Face
}

func NewGameOver(layoutWidth, layoutHeight int) *GameOver {
	gameOver := &GameOver{}

	tt := LoadFont("mplus-1p-regular.ttf")

	var err error
	const dpi = 72
	gameOver.mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	gameOver.mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return gameOver
}

func (g *GameOver) UpdateKeys(keys []ebiten.Key) (retry bool) {
	// No key pressed
	if len(keys) == 0 {
		return false
	}

	return keys[0] == ebiten.KeySpace
}

func (g *GameOver) Draw(screen *ebiten.Image, score int) {
	darkGrey := color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xff}
	white := color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff}

	title := "Game over"
	scoring := "Your score is " + strconv.Itoa(score)

	b1 := text.BoundString(g.mplusBigFont, title)
	posX1 := (320 / 2) - (b1.Dx() / 2)
	posY1 := (240 / 4) + (b1.Dy() / 2)

	text.Draw(screen, title, g.mplusBigFont, posX1, posY1, white)

	b2 := text.BoundString(g.mplusBigFont, scoring)
	posX2 := (320 / 2) - (b2.Dx() / 2)
	posY2 := (240 / 2) + (b2.Dy() / 2)

	text.Draw(screen, scoring, g.mplusBigFont, posX2, posY2, white)

	instructions := "Press space bar to retry"

	b3 := text.BoundString(g.mplusNormalFont, instructions)
	posX3 := (320 / 2) - (b3.Dx() / 2)
	posY3 := (240 / 3 * 2) + (b3.Dy() / 2)

	text.Draw(screen, instructions, g.mplusNormalFont, posX3, posY3, darkGrey)
}
