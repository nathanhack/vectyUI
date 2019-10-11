package border

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/style/borderStyle"
)

type Value struct {
	Width string
	Style borderStyle.Type
	Color string
}

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("border", v.Width+" "+string(v.Style)+" "+v.Color).Apply(h)
}
