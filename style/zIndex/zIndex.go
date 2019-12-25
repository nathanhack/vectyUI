package zIndex

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
	vecty.Style("z-index", string(t)).Apply(h)
}

func Number(number interface{}) Type {
	return Type(internal.Stringify(number, ""))
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("z-index", string(v)).Apply(h)
}
