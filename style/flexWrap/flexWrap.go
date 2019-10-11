package flexWrap

import "github.com/gopherjs/vecty"

type Type string

const (
	Nowrap      Type = "nowrap"
	Wrap        Type = "wrap"
	WrapReverse Type = "wrap-reverse"
)

func (fwt Type) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-flex-wrap", string(fwt)).Apply(h)
	vecty.Style("flex-wrap", string(fwt)).Apply(h)
}
