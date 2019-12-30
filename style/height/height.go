package height

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
)

type Type string

const (
	Auto    Type = "auto"
	Initial Type = "initial"
	Inherit Type = "inherit"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("height", string(t)).Apply(h)
}

func Percent(percent interface{}) Type {
	return Type(internal.Stringify(percent, "%"))
}

func Pixels(length interface{}) Type {
	return Type(internal.Stringify(length, "px"))
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("height", string(v)).Apply(h)
}
