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

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("position", string(t)).Apply(h)
}
