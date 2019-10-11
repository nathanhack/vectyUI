package background

import "github.com/gopherjs/vecty"

type Type string

func (bt Type) Apply(h *vecty.HTML) {
	vecty.Style("background", string(bt)).Apply(h)
}
