package marginTop

import "github.com/gopherjs/vecty"

type Type string

const (
	Auto Type = "auto"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("margin-top", string(t)).Apply(h)
}

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("margin-top", string(v)).Apply(h)
}
