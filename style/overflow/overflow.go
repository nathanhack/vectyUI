package overflow

import "github.com/gopherjs/vecty"

type Type string

const (
	Visible Type = "visible"
	Hidden  Type = "hidden"
	Scroll  Type = "scroll"
	Auto    Type = "auto"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("overflow", string(t)).Apply(h)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("overflow", string(v)).Apply(h)
}
