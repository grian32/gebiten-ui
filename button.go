package gebiten_ui

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GButtonClick func()

type GButton struct {
	text       string
	textPos    Vec2
	minX, minY int
	maxX, maxY int
	x, y       float64 // TODO: maybe do this vec2? not sure..
	font       *GFont
	tex        *ebiten.Image
	onClick    GButtonClick
}

// TODO: semi dupe code between these, but would end up measuring twice if i did bounds check then called NewButtonNoCheck

// NewButton checks that the string fits within the texture then return a new GButton, if you don't want bounds checks
// you can use NewButtonNoCheck
func NewButton(text string, x, y float64, tex *ebiten.Image, font *GFont, onClick GButtonClick) (*GButton, error) {
	strWidth, strHeight := font.MeasureString(text)
	texBounds := tex.Bounds()
	if strWidth >= float64(texBounds.Dx()) || strHeight >= float64(texBounds.Dy()) {
		return nil, fmt.Errorf("string %s does not fit within texture")
	}

	metrics := font.face.Metrics()
	height := math.Max(metrics.XHeight, metrics.CapHeight)

	textX := (float64(texBounds.Dx()) - strWidth) / 2.0
	textY := (float64(texBounds.Dy()) + strHeight + height) / 2.0

	// float to int can be expensive, probably micro, but might aswell
	intX := int(x)
	intY := int(y)

	return &GButton{
		text:    text,
		textPos: Vec2{X: textX, Y: textY},
		minX:    intX,
		minY:    intY,
		maxX:    intX + texBounds.Dx(),
		maxY:    intY + texBounds.Dy(),
		x:       x,
		y:       y,
		tex:     tex,
		font:    font,
		onClick: onClick,
	}, nil
}

// NewButtonNoCheck builds a GButton without checking bounds of string against texture
func NewButtonNoCheck(text string, x, y float64, tex *ebiten.Image, font *GFont, onClick GButtonClick) *GButton {
	strWidth, strHeight := font.MeasureString(text)
	texBounds := tex.Bounds()

	metrics := font.face.Metrics()
	height := math.Max(metrics.XHeight, metrics.CapHeight)

	textX := (float64(texBounds.Dx()) - strWidth) / 2.0
	textY := (float64(texBounds.Dy()) + strHeight + height) / 2.0

	// float to int can be expensive, probably micro, but might aswell
	intX := int(x)
	intY := int(y)

	return &GButton{
		text:    text,
		textPos: Vec2{X: textX, Y: textY},
		minX:    intX,
		minY:    intY,
		maxX:    intX + texBounds.Dx(),
		maxY:    intY + texBounds.Dy(),
		x:       x,
		y:       y,
		tex:     tex,
		font:    font,
		onClick: onClick,
	}
}

func (gb *GButton) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		cursorX, cursorY := ebiten.CursorPosition()
		if cursorX >= gb.minX && cursorX <= gb.maxX && cursorY >= gb.minY && cursorY <= gb.maxY {
			gb.onClick()
		}
	}
}

// Draw draws the button and attempts to center label within texture
func (gb *GButton) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(gb.x, gb.y)
	op.DisableMipmaps = true

	tOp := &text.DrawOptions{}
	tOp.GeoM.Translate(gb.textPos.X, gb.textPos.Y)

	screen.DrawImage(
		gb.tex,
		op,
	)

	text.Draw(screen, gb.text, gb.font.face, tOp)
}
