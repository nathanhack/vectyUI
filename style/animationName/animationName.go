package animationName

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	None    Type = "none"
	Initial Type = "initial"
	Inherit Type = "inherit"

	styleName = "animation-name"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}
