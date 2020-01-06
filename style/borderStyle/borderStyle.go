package borderStyle

import (
	"github.com/gopherjs/vecty"
	"strings"
)

type Type string

const (
	Dashed  Type = "dashed"
	Dotted  Type = "dotted"
	Double  Type = "double"
	Groove  Type = "groove"
	Hidden  Type = "hidden"
	Inset   Type = "inset"
	None    Type = "none"
	Outset  Type = "outset"
	Ridge   Type = "ridge"
	Solid   Type = "solid"
	Initial Type = "initial"
	Inherit Type = "inherit"

	styleName = "border-style"
)

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

type Value []Type

func (v Value) AddTo(m map[string]string) {
	sb := strings.Builder{}
	for _, l := range v {
		sb.WriteString(string(l) + " ")
	}
	m[styleName] = sb.String()
}

func (v Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for _, l := range v {
		sb.WriteString(string(l) + " ")
	}
	vecty.Style(styleName, sb.String()).Apply(h)
}
