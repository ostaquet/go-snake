package snake

import "github.com/hajimehoshi/ebiten/v2"

type Menu struct {
}

func NewMenu(layoutWidth, layoutHeight int) *Menu {
	return &Menu{}
}

func (m *Menu) UpdateKeys(keys []ebiten.Key) error {
	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {

}
