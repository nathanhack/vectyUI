package width

import "github.com/gopherjs/vecty"

type Type string

const (
	Auto Type = "auto"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("width", string(t)).Apply(h)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("width", string(v)).Apply(h)
}
