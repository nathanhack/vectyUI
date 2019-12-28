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
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("justify-content", string(t)).Apply(h)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("justify-content", string(v)).Apply(h)
}
