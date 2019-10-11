package color

import "github.com/gopherjs/vecty"

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("color", string(v)).Apply(h)
}
