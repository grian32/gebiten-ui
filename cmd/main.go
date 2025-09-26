package main

import (
	gebitenui "gebiten-ui"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var btn *gebitenui.GButton

func init() {
	btnTex, _, err := ebitenutil.NewImageFromFile("../testdata/btn.png")
	if err != nil {
		log.Fatalln(err)
	}

	fnt, err := gebitenui.NewGFont("../testdata/arial.ttf", 12)
	if err != nil {
		log.Fatalln(err)
	}

	gbtn, err := gebitenui.NewButton("Press Me!", 0, 20, btnTex, fnt, func() {
		log.Println("Hello!")
	})

	if err != nil {
		log.Fatalln(err)
	}

	btn = gbtn
}

type Test struct {
}

func (t *Test) Update() error {
	btn.Update()
	return nil
}

func (t *Test) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	btn.Draw(screen)
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
