package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

type Menu struct {
	mplusNormalFont font.Face
	mplusBigFont    font.Face
}

func NewMenu(layoutWidth, layoutHeight int) *Menu {
	menu := &Menu{}

	tt := LoadFont("mplus-1p-regular.ttf")

	var err error
	const dpi = 72
	menu.mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	menu.mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return menu
}

func (m *Menu) UpdateKeys(keys []ebiten.Key) (startGame bool) {
	// No key pressed
	if len(keys) == 0 {
		return false
	}

	return keys[0] == ebiten.KeySpace
}

func (m *Menu) Draw(screen *ebiten.Image) {
	darkGrey := color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xff}
	white := color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff}

	title := "Go Snake!"

	b1 := text.BoundString(m.mplusBigFont, title)
	posX1 := (320 / 2) - (b1.Dx() / 2)
	posY1 := (240 / 4) + (b1.Dy() / 2)

	text.Draw(screen, title, m.mplusBigFont, posX1, posY1, white)

	instructions := "Press space bar to start"

	b2 := text.BoundString(m.mplusNormalFont, instructions)
	posX2 := (320 / 2) - (b2.Dx() / 2)
	posY2 := (240 / 3 * 2) + (b2.Dy() / 2)

	text.Draw(screen, instructions, m.mplusNormalFont, posX2, posY2, darkGrey)
}
