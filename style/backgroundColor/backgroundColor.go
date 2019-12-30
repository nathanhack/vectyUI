package backgroundColor

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	Transparent Type = "transparent"
	Initial     Type = "initial"
	Inherit     Type = "inherit"
)

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("background-color", string(v)).Apply(h)
}
