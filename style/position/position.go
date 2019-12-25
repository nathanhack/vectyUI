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
	Inherit  Type = "inherit"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("position", string(t)).Apply(h)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("position", string(v)).Apply(h)
}
