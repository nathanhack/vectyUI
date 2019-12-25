package border

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/style/borderStyle"
	"github.com/nathanhack/vectyUI/style/color"
)

type Value struct {
	Width string
	Style borderStyle.Type
	Color color.Type
}

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("border", v.Width+" "+string(v.Style)+" "+string(v.Color)).Apply(h)
}
