package gebiten_ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type GHoverTexture struct {
	minX, minY           int
	maxX, maxY           int
	hoverY               float64
	hoverTextY           float64
	hoverTextX           float64
	x, y                 float64
	hoverTex             *ebiten.Image
	hoverMsg             *string
	hoverFont            *GFont
	hoverTextColor       color.Color
	shouldRenderHoverMsg bool
	tex                  *ebiten.Image
}

func NewHoverTexture(x, y, maxPosY float64, tex *ebiten.Image, hoverMsg *string, hoverTex *ebiten.Image, font *GFont, hoverTextColor color.Color) *GHoverTexture {
	var hoverY float64

	texHeight := float64(tex.Bounds().Dy())
	hoverTexHeight := float64(hoverTex.Bounds().Dy())

	belowY := y + texHeight
	if belowY < maxPosY {
		hoverY = belowY
	} else {
		hoverY = y - hoverTexHeight
	}

	intX := int(x)
	intY := int(y)

	return &GHoverTexture{
		minX:           intX,
		minY:           intY,
		maxX:           intX + tex.Bounds().Dx(),
		maxY:           intY + tex.Bounds().Dy(),
		x:              x,
		y:              y,
		tex:            tex,
		hoverY:         hoverY,
		hoverTex:       hoverTex,
		hoverMsg:       hoverMsg,
		hoverFont:      font,
		hoverTextColor: hoverTextColor,
	}
}

func (ght *GHoverTexture) Update() {
	x, y := ebiten.CursorPosition()

	if x >= ght.minX && x <= ght.maxX && y >= ght.minY && y <= ght.maxY {
		textWidth, textHeight := ght.hoverFont.MeasureString(*ght.hoverMsg)

		hoverTexBounds := ght.hoverTex.Bounds()
		hoverTexWidth := float64(hoverTexBounds.Dx())
		hoverTexHeight := float64(hoverTexBounds.Dy())

		ght.hoverTextX = ght.x + (hoverTexWidth-textWidth)/2.0
		ght.hoverTextY = ght.hoverY + (hoverTexHeight-textHeight)/2.0
		ght.shouldRenderHoverMsg = true
	} else if ght.shouldRenderHoverMsg {
		ght.shouldRenderHoverMsg = false
	}
}

func (ght *GHoverTexture) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ght.x, ght.y)

	screen.DrawImage(
		ght.tex,
		op,
	)
	if ght.shouldRenderHoverMsg {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(ght.x, ght.hoverY)

		screen.DrawImage(ght.hoverTex, op)
		ght.hoverFont.Draw(screen, *ght.hoverMsg, ght.hoverTextX, ght.hoverTextY, ght.hoverTextColor)
	}
}
