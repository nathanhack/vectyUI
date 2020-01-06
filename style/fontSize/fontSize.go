package fontSize

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
)

type Type string

const (
	Medium  Type = "medium"
	XXSmall Type = "xx-small"
	XSmall  Type = "x-small"
	Small   Type = "small"
	Large   Type = "large"
	XLarge  Type = "x-large"
	XXLarge Type = "xx-large"
	Smaller Type = "smaller"
	Larger  Type = "larger"
	Initial Type = "initial"
	Inherit Type = "inherit"

	styleName = "font-size"
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

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}

func (v Value) AddTo(m map[string]string) {
	m[styleName] = string(v)
}
