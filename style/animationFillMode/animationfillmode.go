package animationFillMode

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	None      Type = "none"
	Forwards  Type = "forwards"
	Backwards Type = "backwards"
	Both      Type = "both"
	Initial   Type = "initial"
	Inherit   Type = "inherit"

	styleName = "animation-fill-mode"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}
