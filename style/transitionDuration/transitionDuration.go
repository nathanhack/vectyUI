package transitionDuration

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
)

type Type string

const (
	Initial Type = "initial"
	Inherit Type = "inherit"
	Unset   Type = "unset"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-transition-duration", string(t)).Apply(h)
	vecty.Style("transition-duration", string(t)).Apply(h)
}

func Seconds(timeInSeconds interface{}) Type {
	return Type(internal.Stringify(timeInSeconds, "s"))
}

func Milliseconds(timeInMilliseconds interface{}) Type {
	return Type(internal.Stringify(timeInMilliseconds, "ms"))
}

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-transition-duration", string(v)).Apply(h)
	vecty.Style("transition-duration", string(v)).Apply(h)
}
