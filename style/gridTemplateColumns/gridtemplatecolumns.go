package gridTemplateColumns

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	None       Type = "none"
	MaxContent Type = "max-content"
	MinContent Type = "min-content"
	Initial    Type = "initial"
	Inherit    Type = "inherit"
	Unset      Type = "unset"

	styleName = "grid-template-columns"
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
