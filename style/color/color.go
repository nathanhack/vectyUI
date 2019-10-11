package color

import "github.com/gopherjs/vecty"

type Type string

func (ct Type) Apply(h *vecty.HTML) {
	vecty.Style("color", string(ct)).Apply(h)
}
