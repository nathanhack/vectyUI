package background

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	Initial   Type = "initial"
	Inherit   Type = "inherit"
	styleName      = "background"
)

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

type Value string

func (v Value) AddTo(m map[string]string) {
	m[styleName] = string(v)
}

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}
