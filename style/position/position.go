package position

import "github.com/gopherjs/vecty"

type Type string

const (
	Static   Type = "static"
	Absolute Type = "absolute"
	Fixed    Type = "fixed"
	Relative Type = "relative"
	Sticky   Type = "sticky"
	Initial  Type = "initial"
)

func (pt Type) Apply(h *vecty.HTML) {
	vecty.Style("position", string(pt)).Apply(h)
}
