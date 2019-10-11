package marginBottom

import "github.com/gopherjs/vecty"

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("margin-bottom", string(v)).Apply(h)
}
