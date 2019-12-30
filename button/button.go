package button

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"github.com/nathanhack/vectyUI/style/background"
	"github.com/nathanhack/vectyUI/style/border"
	"github.com/nathanhack/vectyUI/style/color"
	"github.com/nathanhack/vectyUI/style/margin"
	"github.com/nathanhack/vectyUI/style/padding"
	"github.com/nathanhack/vectyUI/style/userSelect"
	"strconv"
)

const (
	defaultSimpleButtonID  = "simpleButtonID"
	defaultDivButtonID     = "divButtonID"
	defaultGenericButtonID = "genericButtonID"
)

var (
	simpleButtonCount  = 0
	divButtonCount     = 0
	genericButtonCount = 0
)

type Simple struct {
	vecty.Core
	ID                 string               `vecty:"prop"`
	Text               string               `vecty:"prop"`
	Color              color.Value          `vecty:"prop"`
	Background         background.Value     `vecty:"prop"`
	Padding            padding.Value        `vecty:"prop"`
	Margin             margin.Value         `vecty:"prop"`
	Border             border.Value         `vecty:"prop"`
	HoverColor         color.Value          `vecty:"prop"`
	HoverBackground    background.Value     `vecty:"prop"`
	DisabledColor      color.Value          `vecty:"prop"`
	DisabledBackground background.Value     `vecty:"prop"`
	Click              func(i *vecty.Event) `vecty:"prop"`
	MouseEnter         func(i *vecty.Event) `vecty:"prop"`
	MouseLeave         func(i *vecty.Event) `vecty:"prop"`
	Extra              []vecty.Applyer      `vecty:"prop"`
	Disabled           bool                 `vecty:"prop"`
	hovering           bool
}

func (s *Simple) mouseEnter(i *vecty.Event) {
	s.hovering = true
	if s.MouseEnter != nil {
		s.MouseEnter(i)
	}
	vecty.Rerender(s)
}

func (s *Simple) mouseLeave(i *vecty.Event) {
	s.hovering = false
	if s.MouseLeave != nil {
		s.MouseLeave(i)
	}
	vecty.Rerender(s)
}

func (s *Simple) Render() vecty.ComponentOrHTML {
	if s.ID == "" {
		s.ID = defaultSimpleButtonID + strconv.Itoa(simpleButtonCount)
		simpleButtonCount++
	}

	bg := s.Background
	switch {
	case s.Disabled && s.DisabledBackground != "":
		bg = s.DisabledBackground
	case !s.Disabled && s.hovering && s.HoverBackground != "":
		bg = s.HoverBackground
	}

	col := s.Color
	switch {
	case s.Disabled && s.DisabledColor != "":
		col = s.DisabledColor
	case !s.Disabled && s.hovering && s.HoverBackground != "":
		col = s.HoverColor
	}

	markups := []vecty.Applyer{
		prop.ID(s.ID),
		s.Padding,
		s.Margin,
		s.Border,
		col,
		bg,
		prop.Disabled(s.Disabled),
		userSelect.None,
		event.MouseEnter(s.mouseEnter),
		event.MouseLeave(s.mouseLeave),
		vecty.MarkupIf(s.Click != nil, event.Click(s.Click)),
	}
	markups = append(markups, s.Extra...)

	return elem.Button(
		vecty.Markup(
			markups...,
		),
		vecty.Text(s.Text),
	)
}

type Div struct {
	vecty.Core
	ID                 string               `vecty:"prop"`
	Text               string               `vecty:"prop"`
	Color              color.Value          `vecty:"prop"`
	Background         background.Value     `vecty:"prop"`
	Padding            padding.Value        `vecty:"prop"`
	Margin             margin.Value         `vecty:"prop"`
	Border             border.Value         `vecty:"prop"`
	HoverColor         color.Value          `vecty:"prop"`
	HoverBackground    background.Value     `vecty:"prop"`
	DisabledColor      color.Value          `vecty:"prop"`
	DisabledBackground background.Value     `vecty:"prop"`
	Click              func(i *vecty.Event) `vecty:"prop"`
	MouseEnter         func(i *vecty.Event) `vecty:"prop"`
	MouseLeave         func(i *vecty.Event) `vecty:"prop"`
	Extra              []vecty.Applyer      `vecty:"prop"`
	Disabled           bool                 `vecty:"prop"`
	hovering           bool
}

func (d *Div) mouseEnter(i *vecty.Event) {
	d.hovering = true
	if d.MouseEnter != nil {
		d.MouseEnter(i)
	}
	vecty.Rerender(d)
}

func (d *Div) mouseLeave(i *vecty.Event) {
	d.hovering = false
	if d.MouseLeave != nil {
		d.MouseLeave(i)
	}
	vecty.Rerender(d)
}

func (d *Div) Render() vecty.ComponentOrHTML {
	if d.ID == "" {
		d.ID = defaultDivButtonID + strconv.Itoa(divButtonCount)
		divButtonCount++
	}

	bg := d.Background
	switch {
	case d.Disabled && d.DisabledBackground != "":
		bg = d.DisabledBackground
	case !d.Disabled && d.hovering && d.HoverBackground != "":
		bg = d.HoverBackground
	}

	col := d.Color
	switch {
	case d.Disabled && d.DisabledColor != "":
		col = d.DisabledColor
	case !d.Disabled && d.hovering && d.HoverBackground != "":
		col = d.HoverColor
	}
	markups := []vecty.Applyer{
		d.Padding,
		d.Margin,
		d.Border,
		col,
		bg,
		userSelect.None,
		event.MouseEnter(d.mouseEnter),
		event.MouseLeave(d.mouseLeave),
		vecty.MarkupIf(!d.Disabled && d.Click != nil, event.Click(d.Click)),
	}
	markups = append(markups, d.Extra...)

	return elem.Div(
		vecty.Markup(
			markups...,
		),
		vecty.Text(d.Text),
	)
}

type Generic struct {
	vecty.Core
	ID          string                                `vecty:"prop"`
	Div         vecty.ComponentOrHTML                 `vecty:"prop"`
	HoverDiv    vecty.ComponentOrHTML                 `vecty:"prop"`
	DisabledDiv vecty.ComponentOrHTML                 `vecty:"prop"`
	Click       func(i *vecty.Event, button *Generic) `vecty:"prop"`
	MouseEnter  func(i *vecty.Event, button *Generic) `vecty:"prop"`
	MouseLeave  func(i *vecty.Event, button *Generic) `vecty:"prop"`
	Extra       []vecty.Applyer                       `vecty:"prop"`
	Disabled    bool                                  `vecty:"prop"`
	hovering    bool
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
	if g.ID == "" {
		g.ID = defaultGenericButtonID + strconv.Itoa(genericButtonCount)
		genericButtonCount++
	}

	markups := []vecty.Applyer{
		prop.ID(g.ID),
		vecty.MarkupIf(!g.Disabled, event.MouseEnter(g.mouseEnter)),
		vecty.MarkupIf(!g.Disabled, event.MouseLeave(g.mouseLeave)),
		vecty.MarkupIf(!g.Disabled, event.Click(g.click)),
	}
	markups = append(markups, g.Extra...)

	switch {
	case g.Disabled && g.DisabledDiv != nil:
		return elem.Div(
			vecty.Markup(
				markups...,
			),
			g.DisabledDiv,
		)
	case g.hovering && g.HoverDiv != nil:
		return elem.Div(
			vecty.Markup(
				markups...,
			),
			g.HoverDiv,
		)
	case g.Div != nil:
		return elem.Div(
			vecty.Markup(
				markups...,
			),
			g.Div,
		)
	}

	return elem.Div(
		vecty.Markup(
			markups...,
		),
	)
}
