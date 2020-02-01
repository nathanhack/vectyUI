package justifyContent

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	Center       Type = "center"
	Start        Type = "start"
	End          Type = "end"
	FlexStart    Type = "flex-start"
	FlexEnd      Type = "flex-end"
	Left         Type = "left"
	Right        Type = "right"
	Normal       Type = "normal"
	SpaceBetween Type = "space-between"
	SpaceAround  Type = "space-around"
	SpaceEvenly  Type = "space-evenly"
	Stretch      Type = "stretch"
	SafeCenter   Type = "safe center"
	UnsafeCenter Type = "unsafe center"
	Initial      Type = "initial"
	Inherit      Type = "inherit"
	Unset        Type = "unset"

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
