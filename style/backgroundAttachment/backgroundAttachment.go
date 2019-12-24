package backgroundAttachment

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	Scroll Type = "scroll"
	Fixed  Type = "fixed"
	Local  Type = "local"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("background-attachment", string(t)).Apply(h)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("background-attachment", string(v)).Apply(h)
}
