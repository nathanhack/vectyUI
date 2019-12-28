package button

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/nathanhack/vectyUI/style/background"
	"github.com/nathanhack/vectyUI/style/border"
	"github.com/nathanhack/vectyUI/style/margin"
	"github.com/nathanhack/vectyUI/style/padding"
	"github.com/nathanhack/vectyUI/style/userSelect"
)

type Simple struct {
	vecty.Core
	Text            string               `vecty:"prop"`
	Background      background.Value     `vecty:"prop"`
	Padding         padding.Value        `vecty:"prop"`
	Margin          margin.Value         `vecty:"prop"`
	Border          border.Value         `vecty:"prop"`
	HoverBackground background.Value     `vecty:"prop"`
	Click           func(i *vecty.Event) `vecty:"prop"`
	MouseEnter      func(i *vecty.Event) `vecty:"prop"`
	MouseLeave      func(i *vecty.Event) `vecty:"prop"`
	Extra           []vecty.Applyer      `vecty:"prop"`
	Disabled        bool                 `vecty:"prop"`
	hovering        bool
}

func (b *Simple) mouseEnter(i *vecty.Event) {
	b.hovering = true
	if b.MouseEnter != nil {
		b.MouseEnter(i)
	}
	vecty.Rerender(b)
}

func (b *Simple) mouseLeave(i *vecty.Event) {
	b.hovering = false
	if b.MouseLeave != nil {
		b.MouseLeave(i)
	}
	vecty.Rerender(b)
}

func (b *Simple) Render() vecty.ComponentOrHTML {
	markups := []vecty.Applyer{
		b.Padding,
		b.Margin,
		b.Border,
		userSelect.None,
		event.MouseEnter(b.mouseEnter),
		event.MouseLeave(b.mouseLeave),
		vecty.MarkupIf(b.Click != nil, event.Click(b.Click)),
		vecty.MarkupIf(!b.hovering, b.Background),
		vecty.MarkupIf(b.hovering, b.HoverBackground),
	}
	markups = append(markups, b.Extra...)

	return elem.Button(
		vecty.Markup(
			markups...,
		),
		vecty.Text(b.Text),
	)
}

type Div struct {
	vecty.Core
	Text            string               `vecty:"prop"`
	Background      background.Value     `vecty:"prop"`
	Padding         padding.Value        `vecty:"prop"`
	Margin          margin.Value         `vecty:"prop"`
	Border          border.Value         `vecty:"prop"`
	HoverBackground background.Value     `vecty:"prop"`
	Click           func(i *vecty.Event) `vecty:"prop"`
	MouseEnter      func(i *vecty.Event) `vecty:"prop"`
	MouseLeave      func(i *vecty.Event) `vecty:"prop"`
	Extra           []vecty.Applyer      `vecty:"prop"`
	hovering        bool
}

func (b *Div) mouseEnter(i *vecty.Event) {
	b.hovering = true
	if b.MouseEnter != nil {
		b.MouseEnter(i)
	}
	vecty.Rerender(b)
}

func (b *Div) mouseLeave(i *vecty.Event) {
	b.hovering = false
	if b.MouseLeave != nil {
		b.MouseLeave(i)
	}
	vecty.Rerender(b)
}

func (b *Div) Render() vecty.ComponentOrHTML {
	markups := []vecty.Applyer{
		b.Padding,
		b.Margin,
		b.Border,
		userSelect.None,
		event.MouseEnter(b.mouseEnter),
		event.MouseLeave(b.mouseLeave),
		vecty.MarkupIf(b.Click != nil, event.Click(b.Click)),
		vecty.MarkupIf(!b.hovering, b.Background),
		vecty.MarkupIf(b.hovering, b.HoverBackground),
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
	Div        func() vecty.ComponentOrHTML          `vecty:"prop"`
	HoverDiv   func() vecty.ComponentOrHTML          `vecty:"prop"`
	Click      func(i *vecty.Event, button *Generic) `vecty:"prop"`
	MouseEnter func(i *vecty.Event, button *Generic) `vecty:"prop"`
	MouseLeave func(i *vecty.Event, button *Generic) `vecty:"prop"`
	Extra      []vecty.Applyer                       `vecty:"prop"`
	Disabled   bool                                  `vecty:"prop"`
	hovering   bool
}

func (g *Generic) click(i *vecty.Event) {
	if g.Click != nil {
		g.Click(i, g)
	}
}

func (g *Generic) mouseEnter(i *vecty.Event) {
	if g.MouseEnter != nil {
		g.MouseEnter(i, g)
	} else {
		g.hovering = true
		vecty.Rerender(g)
	}
}

func (g *Generic) mouseLeave(i *vecty.Event) {
	if g.MouseLeave != nil {
		g.MouseLeave(i, g)
	} else {
		g.hovering = false
		vecty.Rerender(g)
	}
}

func (g *Generic) Hovering() bool {
	return g.hovering
}

func (g *Generic) SetHover(hovering bool) {
	g.hovering = hovering
}

func (g *Generic) Render() vecty.ComponentOrHTML {
	markups := []vecty.Applyer{
		event.MouseEnter(g.mouseEnter),
		event.MouseLeave(g.mouseLeave),
		event.Click(g.click),
	}
	markups = append(markups, g.Extra...)

	if g.hovering && g.HoverDiv != nil {
		return elem.Div(
			vecty.Markup(
				markups...,
			),
			g.HoverDiv(),
		)
	}
	if g.Div != nil {
		return elem.Div(
			vecty.Markup(
				markups...,
			),
			g.Div(),
		)
	}
	return elem.Div(
		vecty.Markup(
			markups...,
		),
	)
}
