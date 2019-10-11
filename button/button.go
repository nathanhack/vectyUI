package button

import (
	"fmt"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/nathanhack/vectyUI/style/border"
	"github.com/nathanhack/vectyUI/style/margin"
	"github.com/nathanhack/vectyUI/style/padding"
	"github.com/nathanhack/vectyUI/style/userSelect"
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

func (b *Button) mouseEnter(i *vecty.Event) {
	b.hovering = true
	if b.MouseEnter != nil {
		b.MouseEnter(i)
	}
	vecty.Rerender(b)
}

func (b *Button) mouseLeave(i *vecty.Event) {
	b.hovering = false
	if b.MouseLeave != nil {
		b.MouseLeave(i)
	}
	vecty.Rerender(b)
}

func (b *Button) Render() vecty.ComponentOrHTML {
	markups := []vecty.Applyer{
		b.Padding,
		b.Margin,
		b.Border,
		userSelect.None,
		event.MouseEnter(b.mouseEnter),
		event.MouseLeave(b.mouseLeave),
		vecty.MarkupIf(b.Click != nil, event.Click(b.Click)),
		vecty.MarkupIf(!b.hovering, vecty.Style("background", b.Background)),
		vecty.MarkupIf(b.hovering, vecty.Style("background", b.HoverBackground)),
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
	Border          border.Border        `vecty:"prop"`
	HoverBackground string               `vecty:"prop"`
	Click           func(i *vecty.Event) `vecty:"prop"`
	MouseEnter      func(i *vecty.Event) `vecty:"prop"`
	MouseLeave      func(i *vecty.Event) `vecty:"prop"`
	Extra           []vecty.Applyer      `vecty:"prop"`
	hovering        bool
}

func (b *ButtonDiv) mouseEnter(i *vecty.Event) {
	b.hovering = true
	if b.MouseEnter != nil {
		b.MouseEnter(i)
	}
	vecty.Rerender(b)
}

func (b *ButtonDiv) mouseLeave(i *vecty.Event) {
	b.hovering = false
	if b.MouseLeave != nil {
		b.MouseLeave(i)
	}
	vecty.Rerender(b)
}

func (b *ButtonDiv) Render() vecty.ComponentOrHTML {
	markups := []vecty.Applyer{
		b.Padding,
		b.Margin,
		b.Border,
		userSelect.None,
		event.MouseEnter(b.mouseEnter),
		event.MouseLeave(b.mouseLeave),
		vecty.MarkupIf(b.Click != nil, event.Click(b.Click)),
		vecty.MarkupIf(!b.hovering, vecty.Style("background", b.Background)),
		vecty.MarkupIf(b.hovering, vecty.Style("background", b.HoverBackground)),
	}
	markups = append(markups, b.Extra...)

	return elem.Div(
		vecty.Markup(
			markups...,
		),
		vecty.Text(b.Text),
	)
}

type Generic struct {
	vecty.Core
	Div        func() *vecty.HTML                    `vecty:"prop"`
	HoverDiv   func() *vecty.HTML                    `vecty:"prop"`
	Click      func(i *vecty.Event, button *Generic) `vecty:"prop"`
	MouseEnter func(i *vecty.Event, button *Generic) `vecty:"prop"`
	MouseLeave func(i *vecty.Event, button *Generic) `vecty:"prop"`
	Extra      []vecty.Applyer                       `vecty:"prop"`
	hovering   bool
}

func (b *Generic) click(i *vecty.Event) {
	if b.Click != nil {
		b.Click(i, b)
	}
}

func (b *Generic) mouseEnter(i *vecty.Event) {
	fmt.Println("hovering")
	b.hovering = true
	if b.MouseEnter != nil {
		b.MouseEnter(i, b)
	}
	vecty.Rerender(b)
}

func (b *Generic) mouseLeave(i *vecty.Event) {
	b.hovering = false
	if b.MouseLeave != nil {
		b.MouseLeave(i, b)
	}
	vecty.Rerender(b)
}

func (b *Generic) Render() vecty.ComponentOrHTML {
	markups := []vecty.Applyer{
		event.MouseEnter(b.mouseEnter),
		event.MouseLeave(b.mouseLeave),
		event.Click(b.click),
	}
	markups = append(markups, b.Extra...)

	if b.hovering && b.HoverDiv != nil {
		return elem.Div(
			vecty.Markup(
				markups...,
			),
			b.HoverDiv(),
		)
	}
	if b.Div != nil {
		return elem.Div(
			vecty.Markup(
				markups...,
			),
			b.Div(),
		)
	}
	return elem.Div(
		vecty.Markup(
			markups...,
		),
	)
}
