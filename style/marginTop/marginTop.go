package marginTop

import "github.com/gopherjs/vecty"

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("margin-top", string(v)).Apply(h)
}
