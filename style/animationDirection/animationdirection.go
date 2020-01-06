package animationDirection

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	Normal           Type = "normal"
	Reverse          Type = "revers"
	Alternate        Type = "alternate"
	AlternateReverse Type = "alternate-reverse"
	Initial          Type = "initial"
	Inherit          Type = "inherit"

	styleName = "animation-direction"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}
