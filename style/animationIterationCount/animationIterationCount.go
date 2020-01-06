package animationIterationCount

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
)

type Type string

const (
	Infinite Type = "infinite"
	Initial  Type = "initial"
	Inherit  Type = "inherit"

	styleName = "animation-iteration-count"
)

func Number(timeInSeconds interface{}) Type {
	return Type(internal.Stringify(timeInSeconds, "s"))
}

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}
