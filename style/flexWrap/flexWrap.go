package flexWrap

import "github.com/gopherjs/vecty"

type Type string

const (
	Nowrap      Type = "nowrap"
	Wrap        Type = "wrap"
	WrapReverse Type = "wrap-reverse"
	Initial     Type = "initial"
	Inherit     Type = "inherit"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-flex-wrap", string(t)).Apply(h)
	vecty.Style("flex-wrap", string(t)).Apply(h)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-flex-wrap", string(v)).Apply(h)
	vecty.Style("flex-wrap", string(v)).Apply(h)
}
