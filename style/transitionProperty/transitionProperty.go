package transitionProperty

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
)

type Type string

const (
	All     Type = "all"
	None    Type = "none"
	Initial Type = "initial"
	Inherit Type = "inherit"
	Unset   Type = "unset"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-transition-property", string(t)).Apply(h)
	vecty.Style("transition-property", string(t)).Apply(h)
}

func Seconds(timeInSeconds interface{}) Type {
	return Type(internal.Stringify(timeInSeconds, "s"))
}

func Milliseconds(timeInSeconds interface{}) Type {
	return Type(internal.Stringify(timeInSeconds, "ms"))
}

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-transition-property", string(v)).Apply(h)
	vecty.Style("transition-property", string(v)).Apply(h)
}
