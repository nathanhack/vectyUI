package fontstyle

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	Normal  Type = "normal"
	Italic  Type = "italic"
	Oblique Type = "oblique"
	Initial Type = "initial"
	Inherit Type = "inherit"
	Unset   Type = "unset"

	styleName = "font-size"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}

func (v Value) AddTo(m map[string]string) {
	m[styleName] = string(v)
}
