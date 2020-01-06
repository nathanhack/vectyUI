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

	styleName = "line-height"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
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
	vecty.Style(styleName, string(v)).Apply(h)
}

func (v Value) AddTo(m map[string]string) {
	m[styleName] = string(v)
}
