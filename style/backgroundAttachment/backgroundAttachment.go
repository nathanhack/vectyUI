package backgroundAttachment

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	Scroll Type = "scroll"
	Fixed  Type = "fixed"
	Local  Type = "local"

	styleName = "background-attachment"
)

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

type Value Type

func (v Value) AddTo(m map[string]string) {
	m[styleName] = string(v)
}

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}
