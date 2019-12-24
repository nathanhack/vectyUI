package color

import "github.com/gopherjs/vecty"

type Type string

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("color", string(v)).Apply(h)
}
