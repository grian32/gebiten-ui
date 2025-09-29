# gebiten-ui

`gebiten-ui` is a small UI library for Ebitengine. It currently provides a `GButton` and `GTextbox`, plus a thing `GFont` wrapper around Ebitengine's text rendering.

## Features

- Common `GWidget` interface.
- `GButton` with click handling and automatically centered label.
- `GTextureButton`, a simple button that just registers clicks with no additional text drawing.
- `GTextbox`, a single line textbox with focus, caret, text editing and y-axis centering.
- `GFont`, a utility wrapper for drawing text.

## Usage

A more complete example is available in `cmd/main.go`

```go
import (
    gebitenui "gebiten-ui"
)

texture := ...

fnt, _ := gebitenui.NewGFont("font.ttf", 24)
btn, _ := gebitenui.NewButton("Press Me!", 0, 0, texture, fnt, func () {
    // do something
})
```

Inside your Update() function you must call the widget's Update() function

```go
btn.Update()
```

Inside your Draw() function you must call the widget's Draw() function

```go
btn.Draw(screen)
```

