package height

import "github.com/gopherjs/vecty"

type Type string

const (
	Auto Type = "auto"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("height", string(t)).Apply(h)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("height", string(v)).Apply(h)
}
