package marginRight

import "github.com/gopherjs/vecty"

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("margin-right", string(v)).Apply(h)
}
