package whitespace

import "github.com/gopherjs/vecty"

type Type string

const (
	Normal      Type = "normal"
	Pre         Type = "pre"
	Nowrap      Type = "nowrap"
	PreWrap     Type = "pre-wrap"
	PreLine     Type = "pre-line"
	BreakSpaces Type = "break-spaces"

	Initial Type = "initial"
	Inherit Type = "inherit"
	Unset   Type = "unset"

	styleName = "white-space"
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
