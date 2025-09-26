package gebiten_ui

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// GFont thin wrapper around ebiten text drawing mainly meant to be used to pass to other gebitenui components, but can be used independently if desired
type GFont struct {
	face *text.GoXFace
}

func NewGFont(ttfPath string, size float64) (*GFont, error) {
	ttfBytes, err := os.ReadFile(ttfPath)
	if err != nil {
		return nil, err
	}

	fnt, err := opentype.Parse(ttfBytes)
	if err != nil {
		return nil, err
	}
	oFace, err := opentype.NewFace(fnt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	face := text.NewGoXFace(oFace)

	return &GFont{
		face: face,
	}, nil
}

func (gf *GFont) MeasureString(msg string) (float64, float64) {
	return text.Measure(msg, gf.face, 1)
}

func (gf *GFont) Draw(screen *ebiten.Image, msg string, x, y float64) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	text.Draw(screen, msg, gf.face, op)
}
