package marginLeft

import "github.com/gopherjs/vecty"

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("margin-left", string(v)).Apply(h)
}
