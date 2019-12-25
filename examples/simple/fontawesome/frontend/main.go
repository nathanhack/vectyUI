package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/nathanhack/vectyUI/fontawesome"
	"github.com/nathanhack/vectyUI/style"
	"github.com/nathanhack/vectyUI/style/color"
	"github.com/nathanhack/vectyUI/style/fontSize"
)

//go:generate env GOARCH=wasm GOOS=js go build -o ../static/main.wasm

type Body struct {
	vecty.Core
}

func (b Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		&fontawesome.FontAwesome{
			Style: fontawesome.Far,
			Text:  "\uf152",
			Color: style.Color("red"),
		},
		&fontawesome.FontAwesome{
			Style: fontawesome.Far,
			Text:  "\uf058",
			Size:  style.FontSize(fontSize.XXLarge),
			Color: style.Color(color.RGBA(30, 30, 30, 255)),
		},
	)
}

func main() {
	vecty.SetTitle("Simple Example")
	vecty.RenderBody(&Body{})
}
