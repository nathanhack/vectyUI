package overflow

import "github.com/gopherjs/vecty"

type OverflowType string

const (
	Visible OverflowType = "visible"
	Hidden  OverflowType = "hidden"
	Scroll  OverflowType = "scroll"
	Auto    OverflowType = "auto"
)

func (ot OverflowType) Apply(h *vecty.HTML) {
	vecty.Style("overflow", string(ot)).Apply(h)
}
