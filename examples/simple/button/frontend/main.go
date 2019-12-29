package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/nathanhack/vectyUI/button"
	"github.com/nathanhack/vectyUI/style"
	"github.com/nathanhack/vectyUI/style/alignItems"
	"github.com/nathanhack/vectyUI/style/borderStyle"
	"github.com/nathanhack/vectyUI/style/color"
	"github.com/nathanhack/vectyUI/style/display"
	"github.com/nathanhack/vectyUI/style/flexDirection"
	"github.com/nathanhack/vectyUI/style/justifyContent"
	"github.com/nathanhack/vectyUI/style/padding"
	"github.com/nathanhack/vectyUI/style/position"
	"github.com/nathanhack/vectyUI/style/userSelect"
	"syscall/js"
)

//go:generate env GOARCH=wasm GOOS=js go build -o ../static/main.wasm

type Body struct {
	vecty.Core
}

func (b Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(
			display.Flex,
			flexDirection.Column,
			justifyContent.SpaceAround,
			position.Absolute,
			alignItems.Center,
			style.Height("100%"),
			style.Width("100%"),
		),

		elem.Div(
			&button.Simple{
				Text:            "Button.Simple",
				Color:           style.Color(color.RGB(0, 0, 0)),
				Background:      style.Background(color.RGB(123, 44, 180)),
				Padding:         style.Padding(10),
				Margin:          style.Margin(20),
				Border:          style.Border(1, borderStyle.Solid, "red"),
				HoverBackground: style.Background(color.RGB(143, 64, 200)),
				Click: func(i *vecty.Event) {
					js.Global().Get("alert").Invoke("Button.Simple Click")
				},
			},
			&button.Simple{
				Text:            "Button.Simple",
				Color:           style.Color(color.RGB(0, 0, 0)),
				DisabledColor:   "lightgrey",
				Background:      style.Background(color.RGB(123, 44, 180)),
				HoverBackground: style.Background(color.RGB(143, 64, 200)),
				Padding:         style.Padding(10),
				Margin:          style.Margin(20),
				Border:          style.Border(1, borderStyle.Solid, "red"),
				Disabled:        true,
				Click: func(i *vecty.Event) {
					js.Global().Get("alert").Invoke("Button.Simple Click")
				},
			},
		),
		elem.Div(
			&button.Div{
				Text:            "Button.Div",
				Color:           style.Color("black"),
				Background:      "#7b2cb4",
				Padding:         style.Padding(padding.Length(10)),
				Margin:          style.Margin("20px"),
				Border:          style.Border("1px", borderStyle.Solid, "red"),
				HoverBackground: style.Background(color.RGB("143", "64", 200)),
				Click: func(i *vecty.Event) {
					js.Global().Get("alert").Invoke("Button.Div Click")
				},
			},
			&button.Div{
				Text:            "Button.Div",
				Color:           style.Color("black"),
				DisabledColor:   "lightgrey",
				Background:      "#7b2cb4",
				Padding:         style.Padding(padding.Length(10)),
				Margin:          style.Margin("20px"),
				Border:          style.Border("1px", borderStyle.Solid, "red"),
				HoverBackground: style.Background(color.RGB("143", "64", 200)),
				Click: func(i *vecty.Event) {
					js.Global().Get("alert").Invoke("Button.Div Click")
				},
				Disabled: true,
			},
		),

		elem.Div(
			&button.Generic{
				Div: elem.Div(
					vecty.Markup(
						style.Color(color.RGB(0, 0, 0)),
						style.Background(color.RGB(123, 44, 180)),
						style.Padding(10),
						style.Margin(20),
						style.Border(1, borderStyle.Solid, "red"),
					),
					vecty.Text("Button.Generic"),
				),
				HoverDiv: elem.Div(
					vecty.Markup(
						style.Color(color.RGB(0, 0, 0)),
						style.Background(color.RGB(143, 64, 200)),
						style.Padding(10),
						style.Margin(20),
						style.Border(1, borderStyle.Solid, "red"),
					),
					vecty.Text("Button.Generic"),
				),
				Click: func(i *vecty.Event, g *button.Generic) {
					js.Global().Get("alert").Invoke("Button.Generic Click")
				},
			},
			&button.Generic{
				Div: elem.Div(
					vecty.Markup(
						style.Color(color.RGB(0, 0, 0)),
						style.Background(color.RGB(123, 44, 180)),
						style.Padding(10),
						style.Margin(20),
						style.Border(1, borderStyle.Solid, "red"),
						userSelect.None,
					),
					vecty.Text("Button.Generic"),
				),
				HoverDiv: elem.Div(
					vecty.Markup(
						style.Color(color.RGB(0, 0, 0)),
						style.Background(color.RGB(143, 64, 200)),
						style.Padding(10),
						style.Margin(20),
						style.Border(1, borderStyle.Solid, "red"),
						userSelect.None,
					),
					vecty.Text("Button.Generic"),
				),
				Click: func(i *vecty.Event, g *button.Generic) {
					js.Global().Get("alert").Invoke("Button.Generic Click")
				},
				Disabled: true,
			},
		),
	)
}

func main() {
	vecty.SetTitle("Simple Example")
	vecty.RenderBody(&Body{})
}
