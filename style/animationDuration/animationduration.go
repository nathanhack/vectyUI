package animationDuration

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
)

type Type string

const (
	Initial Type = "initial"
	Inherit Type = "inherit"

	styleName = "animation-duration"
)

func Seconds(timeInSeconds interface{}) Type {
	return Type(internal.Stringify(timeInSeconds, "s"))
}

func Milliseconds(timeInMilliseconds interface{}) Type {
	return Type(internal.Stringify(timeInMilliseconds, "ms"))
}

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}
