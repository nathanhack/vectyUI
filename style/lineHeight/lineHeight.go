package lineHeight

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
)

type Type string

const (
	Normal  Type = "normal"
	Initial Type = "initial"
	Inherit Type = "inherit"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("line-height", string(t)).Apply(h)
}

func Percent(percent interface{}) Type {
	return Type(internal.Stringify(percent, "%"))
}

func Pixels(length interface{}) Type {
	return Type(internal.Stringify(length, "px"))
}

//Number - A number that will be multiplied with the current font-size to set the line height
func Number(number interface{}) Type {
	return Type(internal.Stringify(number, ""))
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("line-height", string(v)).Apply(h)
}
