package textAlign

import "github.com/gopherjs/vecty"

type Type string

const (
	Left    Type = "left"
	Right   Type = "right"
	Center  Type = "center"
	Justify Type = "justify"
	Initial Type = "initial"
	Inherit Type = "inherit"

	styleName = "text-align"
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
