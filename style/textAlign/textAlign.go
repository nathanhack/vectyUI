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
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("text-align", string(t)).Apply(h)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("text-align", string(v)).Apply(h)
}
