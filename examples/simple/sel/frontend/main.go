package main

import (
	"fmt"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/nathanhack/vectyUI/fontawesome"
	"github.com/nathanhack/vectyUI/materialdesign"
	"github.com/nathanhack/vectyUI/sel"
	"github.com/nathanhack/vectyUI/style"
	"github.com/nathanhack/vectyUI/style/alignItems"
	"github.com/nathanhack/vectyUI/style/borderStyle"
	"github.com/nathanhack/vectyUI/style/color"
	"github.com/nathanhack/vectyUI/style/cursor"
	"github.com/nathanhack/vectyUI/style/display"
	"github.com/nathanhack/vectyUI/style/flexDirection"
	"github.com/nathanhack/vectyUI/style/justifyContent"
	"github.com/nathanhack/vectyUI/style/position"
	"github.com/nathanhack/vectyUI/style/userSelect"
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
		elem.Select(
			elem.Option(
				vecty.Markup(vecty.Property("value", 0)),
				vecty.Text("op 0"),
			),
			elem.Option(
				vecty.Markup(vecty.Property("value", 1)),
				vecty.Text("op 1")),
			elem.Option(
				vecty.Markup(vecty.Property("value", 2)),
				vecty.Text("op 3"),
			),
		),

		elem.Div(
			vecty.Text("Enabled"),
			&sel.Generic{
				Options: []*sel.GenericOption{
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "teal", "lightseagreen", "Option 1")
						},
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "darkblue", "cornflowerblue", "Option 2")
						},
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "purple", "purple", "Option 3 (no highlight)")
						},
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "brown", "chocolate", "Option 4 (disabled)")
						},
						Disabled: true,
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "black", "grey", "red", "pink", "Option 5")
						},
					},
				},
				Placeholder: sel.GenericOption{
					Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
						return option(disabled, false, true, "white", "grey", "", "", "Choose an Option")
					},
				},
				SelectedOptionStyles: []vecty.Applyer{
					style.Width(250),
					cursor.Pointer,
					userSelect.None,
					style.BackgroundColor(materialdesign.BlueGray800),
					style.BorderRadius("4px"),
					style.Border(1, borderStyle.Solid, "black"),
				},

				SelectedEvent: func(index int, g *sel.Generic) {
					fmt.Println("selected ", index)
				},
				Button: button,
			},
		),

		elem.Div(
			vecty.Text("Enabled no button"),
			&sel.Generic{
				Options: []*sel.GenericOption{
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "teal", "lightseagreen", "Option 1")
						},
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "darkblue", "cornflowerblue", "Option 2")
						},
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "purple", "purple", "Option 3 (no highlight)")
						},
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "brown", "chocolate", "Option 4 (disabled)")
						},
						Disabled: true,
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "black", "grey", "red", "pink", "Option 5")
						},
					},
				},
				Placeholder: sel.GenericOption{
					Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
						return option(disabled, false, true, "white", "grey", "", "", "Choose an Option")
					},
				},
				SelectedOptionStyles: []vecty.Applyer{
					style.BackgroundColor(materialdesign.BlueGray800),
					style.Width(250),
					cursor.Pointer,
					userSelect.None,
				},

				SelectedEvent: func(index int, g *sel.Generic) {
					fmt.Println("selected ", index)
				},
			},
		),

		elem.Div(
			vecty.Text("Disabled"),
			&sel.Generic{
				Options: []*sel.GenericOption{
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "teal", "lightseagreen", "Option 1")
						},
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "darkblue", "cornflowerblue", "Option 2")
						},
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "purple", "purple", "Option 3 (no highlight)")
						},
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "brown", "chocolate", "Option 4 (disabled)")
						},
						Disabled: true,
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "black", "grey", "red", "pink", "Option 5")
						},
					},
				},
				Placeholder: sel.GenericOption{
					Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
						return option(disabled, false, true, "white", "grey", "", "", "Choose an Option")
					},
				},
				SelectedOptionStyles: []vecty.Applyer{
					style.BackgroundColor(materialdesign.BlueGray800),
					style.Width(250),
					cursor.Pointer,
					userSelect.None,
				},

				SelectedEvent: func(index int, g *sel.Generic) {
					fmt.Println("selected ", index)
				},
				Button: button,
				//SelectedPos: 1,
				Disabled: true,
			},
		),

		elem.Div(
			vecty.Text("Disabled: option selected"),
			&sel.Generic{
				Options: []*sel.GenericOption{
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "teal", "lightseagreen", "Option 1")
						},
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "darkblue", "cornflowerblue", "Option 2")
						},
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "purple", "purple", "Option 3 (no highlight)")
						},
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "white", "grey", "brown", "chocolate", "Option 4 (disabled)")
						},
						Disabled: true,
					},
					{
						Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
							return option(disabled, highlight, selected, "black", "grey", "red", "pink", "Option 5")
						},
					},
				},
				Placeholder: sel.GenericOption{
					Option: func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML {
						return option(disabled, false, true, "white", "grey", "", "", "Choose an Option")
					},
				},
				SelectedOptionStyles: []vecty.Applyer{
					style.BackgroundColor(materialdesign.BlueGray800),
					style.Width(250),
					cursor.Pointer,
					userSelect.None,
				},

				Button: button,
				SelectedEvent: func(index int, g *sel.Generic) {
					fmt.Println("selected ", index)
				},
				SelectedPos: 5,
				Disabled:    true,
			},
		),
	)
}

func option(disabled, highlighted, selected bool, text, textDisabled, bg, bgHighlighted color.Type, option string) vecty.ComponentOrHTML {
	textColor := text
	background := bg
	if highlighted {
		background = bgHighlighted
	}

	if selected {
		textColor = color.RGB(255, 255, 255)
		background = materialdesign.BlueGray800
	}
	if disabled {
		textColor = textDisabled
	}

	return elem.Div(
		vecty.Markup(
			style.Color(textColor),
			style.BackgroundColor(background),
			style.Padding(10),
		),
		vecty.Text(option),
	)
}

func button(disabled, over, focused, expended bool) vecty.ComponentOrHTML {
	if expended {
		return &fontawesome.FontAwesome{
			Style: fontawesome.Fas,
			Text:  "\uf077",
			Color: "white",
			Extra: []vecty.Applyer{
				style.Padding(0, 15),
			},
		}
	}
	return &fontawesome.FontAwesome{
		Style: fontawesome.Fas,
		Text:  "\uf078",
		Color: "white",
		Extra: []vecty.Applyer{
			style.Padding(0, 15),
		},
	}
}

func main() {
	vecty.SetTitle("Simple Example")
	vecty.RenderBody(&Body{})
}
