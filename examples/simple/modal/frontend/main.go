package main

import (
	"fmt"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/nathanhack/vectyUI/button"
	"github.com/nathanhack/vectyUI/modal"
	"github.com/nathanhack/vectyUI/style"
	"github.com/nathanhack/vectyUI/style/borderStyle"
	"github.com/nathanhack/vectyUI/style/display"
	"github.com/nathanhack/vectyUI/style/height"
	"github.com/nathanhack/vectyUI/style/margin"
	"github.com/nathanhack/vectyUI/style/width"
)

//go:generate env GOARCH=wasm GOOS=js go build -o ../static/main.wasm

type Body struct {
	vecty.Core
	Modal *modal.Div
}

var exampleModal = modal.Div{
	Background: "#0000004d",
	Display: func(m *modal.Div) vecty.ComponentOrHTML {
		return elem.Div(
			vecty.Markup(
				height.Percent(30),
				width.Percent(30),
				style.Background("white"),
				margin.Auto,
				event.Click(func(v *vecty.Event) {
					m.Hide()
				}),
			),
			vecty.Text("There be a message right here, but there could be buttons and other things. Click the White area to dismiss."),
		)
	},
	TransitionDuration: style.TransitionDuration("1s"),
	Extras:             []vecty.Applyer{display.Flex},
}

var exampleModal2 = modal.Div{
	Background: "#0000004d",
	Display: func(m *modal.Div) vecty.ComponentOrHTML {
		return elem.Div(
			vecty.Markup(
				height.Percent(30),
				width.Percent(30),
				style.Background("white"),
				margin.Auto,
				event.Click(func(v *vecty.Event) {

				}),
			),
			vecty.Text("There be a message right here, but there could be buttons and other things. Click anywhere to dismiss."),
		)
	},
	TransitionDuration: style.TransitionDuration("1s"),
	Dismissible:        true,
	Extras:             []vecty.Applyer{display.Flex},
}

var exampleModal3 = modal.Div{
	Background: "#0000004d",
	Display: func(m *modal.Div) vecty.ComponentOrHTML {
		return elem.Div(
			vecty.Markup(
				height.Percent(30),
				width.Percent(30),
				style.Background("white"),
				margin.Auto,
				event.Click(func(v *vecty.Event) {
					fmt.Println("click")
				}).StopPropagation(),
				event.KeyUp(
					func(v *vecty.Event) {
						fmt.Println(v.Value.Get("code"))
						if !v.Value.IsNull() && v.Value.Get("code").String() == "Escape" {
							m.Hide()
						}
					},
				),
			),
			vecty.Text("There be a message right here, but there could be buttons and other things. Click anything but White area to dismiss or press the Esc key.\n\n Note the nonfunctional button needed so this div would get focus there by allow us to look for the ESC key."),
			&button.Simple{
				Text:       "button",
				Color:      "Black",
				Background: "grey",
				Padding:    style.Padding(10),
				Border:     style.Border(1, borderStyle.Solid, "black"),
				Extra:      []vecty.Applyer{margin.Auto},
			},
		)
	},
	TransitionDuration: style.TransitionDuration("1s"),
	Dismissible:        true,
	Extras:             []vecty.Applyer{display.Flex},
}

func (b *Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(
			style.BackgroundColor("black"),
		),
		&button.Simple{
			Text:       "Show Modal (not Dismissible)",
			Color:      "Black",
			Background: "white",
			Padding:    style.Padding(10),
			Border:     style.Border(1, borderStyle.Solid, "black"),
			Click: func(i *vecty.Event) {
				exampleModal.ShowDiv()
			},
		},
		&button.Simple{
			Text:       "Show Modal (Dismissible with click Anywhere)",
			Color:      "Black",
			Background: "white",
			Padding:    style.Padding(10),
			Border:     style.Border(1, borderStyle.Solid, "black"),
			Click: func(i *vecty.Event) {
				exampleModal2.ShowDiv()
			},
		},
		&button.Simple{
			Text:       "Show Modal (Dismissible with click in black or Esc)",
			Color:      "Black",
			Background: "white",
			Padding:    style.Padding(10),
			Border:     style.Border(1, borderStyle.Solid, "black"),
			Click: func(i *vecty.Event) {
				exampleModal3.ShowDiv()

			},
		},

		&exampleModal,
		&exampleModal2,
		&exampleModal3,
	)
}

func main() {
	vecty.SetTitle("Simple Example")
	vecty.RenderBody(&Body{})
}
