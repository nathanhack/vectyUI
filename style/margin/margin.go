package margin

import (
	"github.com/gopherjs/vecty"
	"strings"
)

type Type string

const (
	Auto    Type = "auto"
	Initial Type = "initial"
	Inherit Type = "inherit"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("margin", string(t)).Apply(h)
}

type Value []string

func (v Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for _, l := range v {
		sb.WriteString(l + " ")
	}
	vecty.Style("margin", sb.String()).Apply(h)
}
