package backgroundColor

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/color"
)

type Value color.Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("background-color", string(v)).Apply(h)
}
