package main

import (
	_ "image/png"
	"log"

	gebitenui "github.com/grian32/gebiten-ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var btn *gebitenui.GButton
var textbox *gebitenui.GTextbox

func init() {
	btnTex, _, err := ebitenutil.NewImageFromFile("../testdata/btn.png")
	if err != nil {
		log.Fatalln(err)
	}
	textboxTex, _, err := ebitenutil.NewImageFromFile("../testdata/textbox.png")
	if err != nil {
		log.Fatalln(err)
	}

	fnt, err := gebitenui.NewGFont("../testdata/arial.ttf", 12)
	if err != nil {
		log.Fatalln(err)
	}

	btn, err = gebitenui.NewButton("Press Me!", 275, 200, btnTex, fnt, func() {
		log.Println("Hello!")
	})
	if err != nil {
		log.Fatalln(err)
	}

	textbox = gebitenui.NewTextBox(20, 200, 12, textboxTex, fnt, 15.0, 0.0)
}

type Test struct {
}

func (t *Test) Update() error {
	btn.Update()
	textbox.Update()
	return nil
}

func (t *Test) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	btn.Draw(screen)
	textbox.Draw(screen)
}

func (t *Test) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Test{}); err != nil {
		log.Fatalln(err)
	}
}
