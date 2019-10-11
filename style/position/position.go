package position

import "github.com/gopherjs/vecty"

type PositionType string

const (
	Static   PositionType = "static"
	Absolute PositionType = "absolute"
	Fixed    PositionType = "fixed"
	Relative PositionType = "relative"
	Sticky   PositionType = "sticky"
	Initial  PositionType = "initial"
)

func (pt PositionType) Apply(h *vecty.HTML) {
	vecty.Style("position", string(pt)).Apply(h)
}
