package gebiten_ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GTextureButtonClick func()

type GTextureButton struct {
	minX, minY int
	maxX, maxY int
	x, y       float64 // TODO: maybe do this vec2? not sure..
	tex        *ebiten.Image
	onClick    GButtonClick
}

// NewTextureButton creates a new GTextureButton
func NewTextureButton(x, y float64, tex *ebiten.Image, onClick GButtonClick) *GTextureButton {
	texBounds := tex.Bounds()

	// float to int can be expensive, probably micro, but might aswell
	intX := int(x)
	intY := int(y)

	return &GTextureButton{
		minX:    intX,
		minY:    intY,
		maxX:    intX + texBounds.Dx(),
		maxY:    intY + texBounds.Dy(),
		x:       x,
		y:       y,
		tex:     tex,
		onClick: onClick,
	}
}

func (gtb *GTextureButton) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		cursorX, cursorY := ebiten.CursorPosition()
		if cursorX >= gtb.minX && cursorX <= gtb.maxX && cursorY >= gtb.minY && cursorY <= gtb.maxY {
			gtb.onClick()
		}
	}
}

func (gtb *GTextureButton) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.Filter = ebiten.FilterNearest
	op.GeoM.Translate(gtb.x, gtb.y)
	op.DisableMipmaps = true

	screen.DrawImage(
		gtb.tex,
		op,
	)
}
