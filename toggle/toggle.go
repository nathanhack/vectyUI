package toggle

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

type Generic struct {
	vecty.Core
	On          func() vecty.ComponentOrHTML          `vecty:"prop"` //these are the bare min to implement
	Off         func() vecty.ComponentOrHTML          `vecty:"prop"` //these are the bare min to implement
	HoverOn     func() vecty.ComponentOrHTML          `vecty:"prop"`
	HoverOff    func() vecty.ComponentOrHTML          `vecty:"prop"`
	DisabledOn  func() vecty.ComponentOrHTML          `vecty:"prop"`
	DisabledOff func() vecty.ComponentOrHTML          `vecty:"prop"`
	OnChange    func(i *vecty.Event, toggle *Generic) `vecty:"prop"`
	MouseEnter  func(i *vecty.Event, toggle *Generic) `vecty:"prop"`
	MouseLeave  func(i *vecty.Event, toggle *Generic) `vecty:"prop"`
	State       bool                                  `vecty:"prop"`
	Extra       []vecty.Applyer                       `vecty:"prop"`
	Disabled    bool                                  `vecty:"prop"`
	Focusable   bool                                  `vecty:"prop"`
	hovering    bool
}

func (g *Generic) click(i *vecty.Event) {
	g.State = !g.State
	vecty.Rerender(g)

	if g.OnChange != nil {
		g.OnChange(i, g)
	}
}

func (g *Generic) mouseEnter(i *vecty.Event) {
	g.hovering = true
	vecty.Rerender(g)

	if g.MouseEnter != nil {
		g.MouseEnter(i, g)
	}
}

func (g *Generic) mouseLeave(i *vecty.Event) {
	g.hovering = false
	vecty.Rerender(g)

	if g.MouseLeave != nil {
		g.MouseLeave(i, g)
	}
}

func (g *Generic) Render() vecty.ComponentOrHTML {
	items := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.MarkupIf(!g.Disabled,
				event.MouseEnter(g.mouseEnter),
				event.MouseLeave(g.mouseLeave),
				event.Click(g.click),
			),
			vecty.MarkupIf(g.Focusable, vecty.Attribute("tabindex", 0)),
			vecty.Markup(g.Extra...),
		),
	}

	switch {
	case g.State && g.Disabled && g.DisabledOn != nil:
		items = append(items, g.DisabledOn())
	case g.State && g.Disabled && g.On != nil:
		items = append(items, g.On())
	case g.State && g.hovering && g.HoverOn != nil:
		items = append(items, g.HoverOn())
	case g.State && g.hovering && g.On != nil:
		items = append(items, g.On())
	case g.State && g.On != nil:
		items = append(items, g.On())

	case !g.State && g.Disabled && g.DisabledOff != nil:
		items = append(items, g.DisabledOff())
	case !g.State && g.Disabled && g.Off != nil:
		items = append(items, g.Off())
	case !g.State && g.hovering && g.HoverOff != nil:
		items = append(items, g.HoverOff())
	case !g.State && g.hovering && g.Off != nil:
		items = append(items, g.Off())
	case !g.State && g.Off != nil:
		items = append(items, g.Off())
	}

	return elem.Div(items...)
}
