package button

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/nathanhack/vectyUI/border"
	"github.com/nathanhack/vectyUI/margin"
	"github.com/nathanhack/vectyUI/padding"
)

type Button struct {
	vecty.Core
	Text            string               `vecty:"prop"`
	Background      string               `vecty:"prop"`
	Padding         padding.Weights      `vecty:"prop"`
	Margin          margin.Weights       `vecty:"prop"`
	Border          border.Border        `vecty:"prop"`
	HoverBackground string               `vecty:"prop"`
	Click           func(i *vecty.Event) `vecty:"prop"`
	MouseEnter      func(i *vecty.Event) `vecty:"prop"`
	MouseLeave      func(i *vecty.Event) `vecty:"prop"`
	Extra           []vecty.Applyer      `vecty:"prop"`
	hovering        bool
}

func (b *Button) Render() vecty.ComponentOrHTML {
	markups := []vecty.Applyer{
		b.Padding,
		b.Margin,
		b.Border,
		vecty.MarkupIf(b.Click != nil, event.Click(b.Click)),
		vecty.MarkupIf(b.MouseEnter != nil, event.MouseEnter(b.MouseEnter)),
		vecty.MarkupIf(b.MouseLeave != nil, event.MouseLeave(b.MouseLeave)),
	}
	markups = append(markups, b.Extra...)

	return elem.Button(
		vecty.Markup(
			markups...,
		),
		vecty.Text(b.Text),
	)
}

type ButtonDiv struct {
	vecty.Core
	Text            string               `vecty:"prop"`
	Background      string               `vecty:"prop"`
	Padding         padding.Weights      `vecty:"prop"`
	Margin          margin.Weights       `vecty:"prop"`
	BorderColor     border.Color         `vecty:"prop"`
	BorderStyle     border.Style         `vecty:"prop"`
	BorderWidth     border.Width         `vecty:"prop"`
	HoverBackground string               `vecty:"prop"`
	Click           func(i *vecty.Event) `vecty:"prop"`
	MouseEnter      func(i *vecty.Event) `vecty:"prop"`
	MouseLeave      func(i *vecty.Event) `vecty:"prop"`
	Extra           []vecty.Applyer      `vecty:"prop"`
	hovering        bool
}

func (b *ButtonDiv) Render() vecty.ComponentOrHTML {
	markups := []vecty.Applyer{
		b.Padding,
		b.Margin,
		b.BorderColor,
		b.BorderStyle,
		b.BorderWidth,
		vecty.MarkupIf(b.Click != nil, event.Click(b.Click)),
		vecty.MarkupIf(b.MouseEnter != nil, event.MouseEnter(b.MouseEnter)),
		vecty.MarkupIf(b.MouseLeave != nil, event.MouseLeave(b.MouseLeave)),
	}
	markups = append(markups, b.Extra...)

	return elem.Div(
		vecty.Markup(
			markups...,
		),
		vecty.Text(b.Text),
	)
}
