package nav

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/nathanhack/vectyUI/style/background"
	"github.com/nathanhack/vectyUI/style/height"
)

type HorzBar struct {
	vecty.Core
	Background background.Value      `vecty:"prop"`
	Height     height.Type           `vecty:"prop"`
	Divs       []vecty.MarkupOrChild `vecty:"prop"`
	Extra      []vecty.Applyer       `vecty:"prop"`
	ID         string                `vecty:"prop"`
}

func (hb *HorzBar) Render() vecty.ComponentOrHTML {
	markups := []vecty.Applyer{
		vecty.MarkupIf(hb.ID != "", prop.ID(hb.ID), nil),
		hb.Height,
		hb.Background,
	}
	markups = append(markups, hb.Extra...)
	children := make([]vecty.MarkupOrChild, 0)
	children = append(children, vecty.Markup(markups...))
	children = append(children, hb.Divs...)
	return elem.Div(children...)
}
