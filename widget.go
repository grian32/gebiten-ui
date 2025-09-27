package gebiten_ui

import "github.com/hajimehoshi/ebiten/v2"

type GWidget interface {
	Update()
	Draw(screen *ebiten.Image)
}
