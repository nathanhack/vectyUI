package alignItems

import "github.com/gopherjs/vecty"

type AlignItems string

const (
	Stretch   AlignItems = "stretch"
	Center    AlignItems = "center"
	FlexStart AlignItems = "flex-start"
	FlexEnd   AlignItems = "flex-end"
	Baseline  AlignItems = "baseline"
)

func (ai AlignItems) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-align-items", string(ai)).Apply(h)
	vecty.Style("align-items", string(ai)).Apply(h)
}
