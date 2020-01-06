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

	styleName = "border"
)

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

type Value struct {
	Width string
	Style borderStyle.Type
	Color color.Type
}

func (v Value) AddTo(m map[string]string) {
	m[styleName] = v.Width + " " + string(v.Style) + " " + string(v.Color)
}

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, v.Width+" "+string(v.Style)+" "+string(v.Color)).Apply(h)
}
