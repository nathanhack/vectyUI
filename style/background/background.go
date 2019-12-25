package background

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	Initial Type = "initial"
	Inherit Type = "inherit"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("background", string(t)).Apply(h)
}

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("background", string(v)).Apply(h)
}
