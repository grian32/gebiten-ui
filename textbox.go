package gebiten_ui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GTextboxOnChange func(newText string)

// GTextbox centers text on the Y axis and accepts a text x/y offset
type GTextbox struct {
	text                     []rune
	metricsHeight            float64
	metricsAscent            float64
	textOffsetX, textOffsetY float64
	textX, textY             float64
	focused                  bool
	maxTextLen               int
	minX, minY               int
	maxX, maxY               int
	x, y                     float64
	font                     *GFont
	tex                      *ebiten.Image
	caretPos                 int
	caretX                   float64
	caretTex                 *ebiten.Image
	caretCurrTickCount       int
	onChange                 GTextboxOnChange
}

// NewTextBox creates a new GTextBox
func NewTextBox(x, y float64, maxTextLen int, tex *ebiten.Image, font *GFont, textOffsetX, textOffsetY float64, onChange GTextboxOnChange) *GTextbox {
	metrics := font.face.Metrics()
	height := math.Max(metrics.XHeight, metrics.CapHeight)
	ascent := metrics.HAscent

	texBounds := tex.Bounds()

	intX := int(x)
	intY := int(y)

	caretTex := ebiten.NewImage(2, int(height))
	caretTex.Fill(color.White)

	return &GTextbox{
		metricsHeight: height,
		metricsAscent: ascent,
		textOffsetX:   textOffsetX,
		textX:         textOffsetX,
		textOffsetY:   textOffsetY,
		maxTextLen:    maxTextLen,
		minX:          intX,
		minY:          intY,
		maxX:          intX + texBounds.Dx(),
		maxY:          intY + texBounds.Dy(),
		x:             x,
		y:             y,
		font:          font,
		tex:           tex,
		caretTex:      caretTex,
		caretX:        textOffsetX,
		onChange:      onChange,
	}
}

func (gt *GTextbox) Update() {
	if gt.focused {
		if len(gt.text) < gt.maxTextLen && gt.caretPos >= 0 {
			newText := ebiten.AppendInputChars(nil)

			if len(newText) > 0 {
				tmp := make([]rune, 0, len(gt.text)+len(newText))
				tmp = append(tmp, gt.text[:gt.caretPos]...)
				tmp = append(tmp, newText...)
				tmp = append(tmp, gt.text[gt.caretPos:]...)
				gt.text = tmp
				gt.caretPos += len(newText)
			}

			gt.onChange(string(gt.text))
		}

		if len(gt.text) != 0 {
			_, strHeight := gt.font.MeasureString(string(gt.text))
			texBounds := gt.tex.Bounds()
			gt.textY = gt.textOffsetY + (float64(texBounds.Dy())-strHeight)/2.0

			strWidthCaret, _ := gt.font.MeasureString(string(gt.text[:gt.caretPos]))
			gt.caretX = gt.textOffsetX + strWidthCaret

			//gt.textPos = Vec2{X: gt.x + textX, Y: gt.y + textY}
		}

		if len(gt.text) > 0 && gt.caretPos > 0 && inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			gt.text = append(gt.text[:gt.caretPos-1], gt.text[gt.caretPos:]...)
			gt.caretPos--
			gt.onChange(string(gt.text))
		}

		if gt.caretPos > 0 && inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			gt.caretPos--
		}
		if gt.caretPos < len(gt.text) && inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			gt.caretPos++
		}

		gt.caretCurrTickCount = gt.caretCurrTickCount % 30
		gt.caretCurrTickCount++
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

	tOp := &text.DrawOptions{}
	tOp.GeoM.Translate(gt.x+gt.textX, gt.y+gt.textY)
	text.Draw(screen, string(gt.text), gt.font.face, tOp)

	if gt.focused && gt.caretCurrTickCount > 15 {
		cOp := &ebiten.DrawImageOptions{}
		// TODO: maybe move to update
		var usedYCoord float64
		if gt.textY == 0 {
			usedYCoord = gt.y + gt.textOffsetY + ((float64(gt.tex.Bounds().Dy()) - gt.metricsHeight) / 2.0)
		} else {
			usedYCoord = gt.y + gt.textY - gt.metricsHeight + gt.metricsAscent
		}
		cOp.GeoM.Translate(gt.x+gt.caretX, usedYCoord)
		screen.DrawImage(gt.caretTex, cOp)
	}
}
