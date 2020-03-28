package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/nathanhack/vectyUI/style"
	"github.com/nathanhack/vectyUI/style/color"
	"github.com/nathanhack/vectyUI/style/transitionDuration"
	"github.com/nathanhack/vectyUI/style/transitionTimingFunction"
	"github.com/nathanhack/vectyUI/style/transitionTimingFunction/jumpTerm"
)

//go:generate env GOARCH=wasm GOOS=js go build -o ../static/main.wasm

type Body struct {
	vecty.Core
}

type Item struct {
	vecty.Core
	Width int
	Color color.Type
}

func (i *Item) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			style.Height(100),
			style.Width(i.Width),
			style.Background(i.Color),
			style.Transitions(
				style.Transition("width 2s", transitionTimingFunction.Steps(5, jumpTerm.Both)),
				style.Transition("background", transitionDuration.Milliseconds(1500)),
			),
			event.MouseEnter(func(v *vecty.Event) {
				i.Width = 300
				i.Color = "blue"
				vecty.Rerender(i)
			}),
			event.MouseLeave(func(v *vecty.Event) {
				i.Width = 100
				i.Color = "red"
				vecty.Rerender(i)
			}),
		),
	)
}

func (b Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(
			style.BackgroundColor("black"),
		),
		&Item{
			Width: 100,
			Color: "purple",
		},
	)
}

func main() {
	vecty.SetTitle("Simple Example")
	vecty.RenderBody(&Body{})
}
