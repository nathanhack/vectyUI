package justifyContent

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	FlexStart    Type = "flex-start"
	FlexEnd      Type = "flex-end"
	Center       Type = "center"
	SpaceBetween Type = "space-between"
	SpaceAround  Type = "space-around"
	Initial      Type = "initial"
	Inherit      Type = "inherit"

	styleName = "justify-content"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

func (v Value) AddTo(m map[string]string) {
	m[styleName] = string(v)
}
