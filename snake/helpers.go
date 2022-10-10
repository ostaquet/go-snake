package snake

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/opentype"
	"image"
	"log"
	"os"
)

func LoadImage(filename string) *ebiten.Image {
	// Load the assets
	dat, err := os.ReadFile("snake/assets/images/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(dat))
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

func LoadFont(filename string) *opentype.Font {
	// Load the asset
	dat, err := os.ReadFile("snake/assets/fonts/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	ft, err := opentype.Parse(dat)
	if err != nil {
		log.Fatal(err)
	}
	return ft
}
