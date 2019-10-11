package color

import "github.com/gopherjs/vecty"

type Value string

func (ct Value) Apply(h *vecty.HTML) {
	vecty.Style("color", string(ct)).Apply(h)
}
