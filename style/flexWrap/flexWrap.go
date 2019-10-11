package flexWrap

import "github.com/gopherjs/vecty"

type Type string

const (
	Nowrap      Type = "nowrap"
	Wrap        Type = "wrap"
	WrapReverse Type = "wrap-reverse"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-flex-wrap", string(t)).Apply(h)
	vecty.Style("flex-wrap", string(t)).Apply(h)
}
