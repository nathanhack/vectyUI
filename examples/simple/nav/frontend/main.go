package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/nathanhack/vectyUI/button"
	"github.com/nathanhack/vectyUI/materialdesign"
	"github.com/nathanhack/vectyUI/nav"
	"github.com/nathanhack/vectyUI/style"
	"github.com/nathanhack/vectyUI/style/backgroundImage"
	"github.com/nathanhack/vectyUI/style/backgroundImage/colorStop"
	"github.com/nathanhack/vectyUI/style/display"
	"github.com/nathanhack/vectyUI/style/float"
	"github.com/nathanhack/vectyUI/style/fontFamily"
	"github.com/nathanhack/vectyUI/style/marginLeft"
	"syscall/js"
)

//go:generate env GOARCH=wasm GOOS=js go build -o ../static/main.wasm

type Body struct {
	vecty.Core
}

func (b Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(
			backgroundImage.LinearGradient(backgroundImage.ToBottom,
				colorStop.From(materialdesign.Amber50),
				colorStop.From(materialdesign.Blue50),
			),
		),
		&nav.HorzBar{
			Background: style.Background(
				backgroundImage.LinearGradient(backgroundImage.ToRight, "blue", "Red"),
			),
			Divs: []vecty.MarkupOrChild{
				&button.ButtonDiv{
					Text:            "click me",
					Padding:         style.Padding(14, 16),
					HoverBackground: "green",
					Click: func(i *vecty.Event) {
						js.Global().Get("alert").Invoke("CLICK")
					},
					Extra: []vecty.Applyer{
						display.InlineBlock,
					},
				},

				elem.Div(
					vecty.Markup(
						marginLeft.Auto,
						float.Right,
						style.Padding(14, 16),
						fontFamily.Arial,
					),
					vecty.Text("BRAND"),
				),
			},
			Extra: []vecty.Applyer{
				display.Block,
			},
		},
	)
}

func main() {
	vecty.SetTitle("Simple Example")
	vecty.RenderBody(&Body{})
}
