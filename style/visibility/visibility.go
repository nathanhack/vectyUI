package visibility

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	Visible  Type = "visible"
	Hidden   Type = "hidden"
	Collapse Type = "collapse"
	Initial  Type = "initial"
	Inherit  Type = "inherit"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("visibility", string(t)).Apply(h)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("visibility", string(v)).Apply(h)
}
