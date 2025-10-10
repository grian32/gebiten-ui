package gebiten_ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type GHoverTexture struct {
	minX, minY           int
	maxX, maxY           int
	x, y                 float64
	hoverMsg             string
	hoverBg              *ebiten.Image
	shouldRenderHoverMsg bool
	tex                  *ebiten.Image
}

func NewGHoverTexture(x, y float64, tex *ebiten.Image, hoverMsg string, font *GFont) *GHoverTexture {
	textX, textY := font.MeasureString(hoverMsg)

	bgTex := ebiten.NewImage(int(textX)+4, int(textY)+4)
	bgTex.Fill(color.Black)

	return &GHoverTexture{
		minX:     int(x),
		minY:     int(y),
		maxX:     int(x) + tex.Bounds().Dx(),
		maxY:     int(y) + tex.Bounds().Dy(),
		x:        x,
		y:        y,
		tex:      tex,
		hoverBg:  bgTex,
		hoverMsg: hoverMsg,
	}
}

func (ght *GHoverTexture) Update() {
	x, y := ebiten.CursorPosition()

	if x >= ght.minX && x <= ght.maxX && y >= ght.minY && y <= ght.maxY {
		ght.shouldRenderHoverMsg = true
	}
}

func (ght *GHoverTexture) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ght.x, ght.y)

	screen.DrawImage(
		ght.tex,
		op,
	)
}
