package overflow

import "github.com/gopherjs/vecty"

type Type string

const (
	Visible Type = "visible"
	Hidden  Type = "hidden"
	Scroll  Type = "scroll"
	Auto    Type = "auto"
)

func (ot Type) Apply(h *vecty.HTML) {
	vecty.Style("overflow", string(ot)).Apply(h)
}