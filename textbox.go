package gebiten_ui

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GTextbox struct {
	text          string
	metricsHeight float64
	textPos       Vec2
	focused       bool
	maxTextLen    int
	minX, minY    int
	maxX, maxY    int
	x, y          float64
	font          *GFont
	tex           *ebiten.Image
}

func NewTextBox(x, y float64, maxTextLen int, tex *ebiten.Image, font *GFont) *GTextbox {
	metrics := font.face.Metrics()
	height := math.Max(metrics.XHeight, metrics.CapHeight)

	texBounds := tex.Bounds()

	intX := int(x)
	intY := int(y)

	return &GTextbox{
		text:          "hello",
		metricsHeight: height,
		maxTextLen:    maxTextLen,
		minX:          intX,
		minY:          intY,
		maxX:          intX + texBounds.Dx(),
		maxY:          intY + texBounds.Dy(),
		x:             x,
		y:             y,
		font:          font,
		tex:           tex,
	}
}

func (gt *GTextbox) Update() {
	if gt.text != "" {
		strWidth, strHeight := gt.font.MeasureString(gt.text)
		texBounds := gt.tex.Bounds()
		textX := (float64(texBounds.Dx()) - strWidth) / 2.0
		textY := (float64(texBounds.Dy()) + strHeight + gt.metricsHeight) / 2.0

		gt.textPos = Vec2{X: textX, Y: textY}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		cursorX, cursorY := ebiten.CursorPosition()
		if cursorX >= gt.minX && cursorX <= gt.maxX && cursorY >= gt.minY && cursorY <= gt.maxY {
			gt.focused = true
		} else if gt.focused {
			gt.focused = false
		}
	}
}

func (gt *GTextbox) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(gt.x, gt.y)
	op.DisableMipmaps = true

	screen.DrawImage(gt.tex, op)
}
