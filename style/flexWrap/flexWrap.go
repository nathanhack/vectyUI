package flexWrap

import "github.com/gopherjs/vecty"

type FlexWrapType string

const (
	Nowrap      FlexWrapType = "nowrap"
	Wrap        FlexWrapType = "wrap"
	WrapReverse FlexWrapType = "wrap-reverse"
)

func (fwt FlexWrapType) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-flex-wrap", string(fwt)).Apply(h)
	vecty.Style("flex-wrap", string(fwt)).Apply(h)
}
