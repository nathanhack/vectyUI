package fontawesome

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/nathanhack/vectyUI/style"
	"github.com/nathanhack/vectyUI/style/color"
	"github.com/nathanhack/vectyUI/style/display"
	"github.com/nathanhack/vectyUI/style/fontSize"
	"github.com/nathanhack/vectyUI/style/lineHeight"
	"github.com/nathanhack/vectyUI/style/position"
	"github.com/nathanhack/vectyUI/style/textAlign"
	"github.com/nathanhack/vectyUI/style/verticalAlign"
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

//FontAwesomeStacker will render
// a stack of fontawesome text
// note Extras for each font is ignored
type FontAwesomeStack struct {
	vecty.Core
	Fonts []FontAwesome   `vecty:"prop"`
	Size  fontSize.Value  `vecty:"prop"`
	Extra []vecty.Applyer `vecty:"prop"`
}

var Stack1X = []vecty.Applyer{
	lineHeight.Inherit,
	position.Absolute,
	style.Width("100%"),
	textAlign.Center,
}
var Stack2X = []vecty.Applyer{
	style.FontSize("2em"),
	position.Absolute,
	style.Width("100%"),
	textAlign.Center,
}

func (fas *FontAwesomeStack) Render() vecty.ComponentOrHTML {
	stackMarkups := []vecty.Applyer{
		fas.Size,
		position.Relative,
		display.InlineBlock,
		verticalAlign.Middle,
		style.Width("2em"),
		style.Height("2em"),
		style.LineHeight("2em"),
	}

	stackMarkups = append(stackMarkups, fas.Extra...)

	stackChildren := make([]vecty.MarkupOrChild, 0)

	stackChildren = append(stackChildren, vecty.Markup(stackMarkups...))

	for _, f := range fas.Fonts {
		markups := []vecty.Applyer{
			f.Style,
			f.Color,
		}

		markups = append(markups, f.Extra...)

		stackChildren = append(stackChildren,
			elem.Italic(
				vecty.Markup(
					markups...,
				),
				vecty.Text(f.Text),
			),
		)
	}

	return elem.Span(stackChildren...)
}
