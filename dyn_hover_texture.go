package gebiten_ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type GDynHoverTexture struct {
	maxXOffset           int
	maxYOffset           int
	maxPosY              float64
	hoverTex             *ebiten.Image
	hoverMsg             *string
	hoverFont            *GFont
	hoverTextColor       color.Color
	shouldRenderHoverMsg bool
	tex                  *ebiten.Image

	currHoverTextX float64
	currHoverTextY float64
	currHoverY     float64
}

func NewDynHoverTexture(maxPosY float64, tex *ebiten.Image, hoverMsg *string, hoverTex *ebiten.Image, font *GFont, hoverTextColor color.Color) *GDynHoverTexture {
	return &GDynHoverTexture{
		maxPosY:        maxPosY,
		maxXOffset:     tex.Bounds().Dx(),
		maxYOffset:     tex.Bounds().Dy(),
		tex:            tex,
		hoverTex:       hoverTex,
		hoverMsg:       hoverMsg,
		hoverFont:      font,
		hoverTextColor: hoverTextColor,
	}
}

func (ght *GDynHoverTexture) Update(currX, currY float64) {
	x, y := ebiten.CursorPosition()
	intX := int(currX)
	intY := int(currY)

	belowY := currY + float64(ght.tex.Bounds().Dy())
	if belowY < ght.maxPosY {
		ght.currHoverY = belowY
	} else {
		ght.currHoverY = currY - float64(ght.hoverTex.Bounds().Dy())
	}

	if x >= intX && x <= intX+ght.maxXOffset && y >= intY && y <= intY+ght.maxYOffset {
		textWidth, textHeight := ght.hoverFont.MeasureString(*ght.hoverMsg)

		hoverTexBounds := ght.hoverTex.Bounds()
		hoverTexWidth := float64(hoverTexBounds.Dx())
		hoverTexHeight := float64(hoverTexBounds.Dy())

		ght.currHoverTextX = currX + (hoverTexWidth-textWidth)/2.0
		ght.currHoverTextY = ght.currHoverY + (hoverTexHeight-textHeight)/2.0
		ght.shouldRenderHoverMsg = true
	} else if ght.shouldRenderHoverMsg {
		ght.shouldRenderHoverMsg = false
	}
}

func (ght *GDynHoverTexture) Draw(screen *ebiten.Image, currX, currY float64) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest
	op.GeoM.Translate(currX, currY)

	screen.DrawImage(
		ght.tex,
		op,
	)
	if ght.shouldRenderHoverMsg {
		op := &ebiten.DrawImageOptions{}
		op.Filter = ebiten.FilterNearest
		op.GeoM.Translate(currX, ght.currHoverY)

		screen.DrawImage(ght.hoverTex, op)
		ght.hoverFont.Draw(screen, *ght.hoverMsg, ght.currHoverTextX, ght.currHoverTextY, ght.hoverTextColor)
	}
}
