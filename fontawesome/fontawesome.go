package fontawesome

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/nathanhack/vectyUI/style/color"
	"github.com/nathanhack/vectyUI/style/fontSize"
)

type Style string

const (
	Fas Style = "fas"
	Far Style = "far"
	Fal Style = "fal"
	Fad Style = "fad"
)

func (s Style) Apply(h *vecty.HTML) {
	vecty.Class(string(s)).Apply(h)
}

//FontAwesome is a simple helper for showing css derived Font Awesome fonts
// the css file must be included by the user, either in the html as a <link>...</link>
// or through vecty.AddStylesheet().
type FontAwesome struct {
	vecty.Core
	Style Style           `vecty:"prop"`
	Text  string          `vecty:"prop"`
	Color color.Value     `vecty:"prop"`
	Size  fontSize.Value  `vecty:"prop"`
	Extra []vecty.Applyer `vecty:"prop"`
}

func (fa *FontAwesome) Render() vecty.ComponentOrHTML {
	markups := []vecty.Applyer{
		fa.Style,
		fa.Color,
		fa.Size,
	}

	markups = append(markups, fa.Extra...)

	return elem.Italic(
		vecty.Markup(
			markups...,
		),
		vecty.Text(fa.Text),
	)
}
