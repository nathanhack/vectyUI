package background

import "github.com/gopherjs/vecty"

type Value string

func (bt Value) Apply(h *vecty.HTML) {
	vecty.Style("background", string(bt)).Apply(h)
}
