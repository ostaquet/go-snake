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

func (m *Menu) UpdateKeys(keys []ebiten.Key) error {
	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	gray := color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xff}

	myText := "Bibi Bear!"

	b := text.BoundString(m.mplusNormalFont, myText)
	posX := (320 / 2) - (b.Dx() / 2)
	posY := (240 / 2) + (b.Dy() / 2)

	text.Draw(screen, myText, m.mplusNormalFont, posX, posY, gray)
}
