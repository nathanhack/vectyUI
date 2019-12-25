package float

import "github.com/gopherjs/vecty"

type Type string

const (
	None  Type = "none"
	Left  Type = "left"
	Right Type = "right"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("float", string(t)).Apply(h)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("float", string(v)).Apply(h)
}
