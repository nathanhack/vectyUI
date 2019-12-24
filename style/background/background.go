package background

import (
	"github.com/gopherjs/vecty"
)

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("background", string(v)).Apply(h)
}
