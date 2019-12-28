package opacity

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
)

type Type string

const (
	Initial Type = "initial"
	Inherit Type = "inherit"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("opacity", string(t)).Apply(h)
}

func Number(value interface{}) Type {
	return Type(internal.Stringify(value, ""))
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("opacity", string(v)).Apply(h)
}
