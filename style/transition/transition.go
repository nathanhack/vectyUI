package transition

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	Initial Type = "initial"
	Inherit Type = "inherit"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("transition", string(t)).Apply(h)
	vecty.Style("-webkit-transition", string(t)).Apply(h)
}

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("transition", string(v)).Apply(h)
	vecty.Style("-webkit-transition", string(v)).Apply(h)
}
