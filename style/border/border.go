package border

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/style/borderStyle"
	"github.com/nathanhack/vectyUI/style/color"
)

type Type string

const (
	Initial Type = "initial"
	Inherit Type = "inherit"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("border", string(t)).Apply(h)
}

type Value struct {
	Width string
	Style borderStyle.Type
	Color color.Type
}

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("border", v.Width+" "+string(v.Style)+" "+string(v.Color)).Apply(h)
}
