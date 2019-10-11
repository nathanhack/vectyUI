package borderStyle

import (
	"github.com/gopherjs/vecty"
	"strings"
)

type Type string

const (
	Dashed Type = "dashed"
	Dotted Type = "dotted"
	Double Type = "double"
	Groove Type = "groove"
	Hidden Type = "hidden"
	Inset  Type = "inset"
	None   Type = "none"
	Outset Type = "outset"
	Ridge  Type = "ridge"
	Solid  Type = "solid"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("border-style", string(t)).Apply(h)
}

type Value []Type

func (v Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for _, l := range v {
		sb.WriteString(string(l) + " ")
	}
	vecty.Style("border-style", sb.String()).Apply(h)
}
