package backgroundClip

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	BorderBox  Type = "border-box"
	PaddingBox Type = "padding-box"
	ContentBox Type = "content-box"
	Text       Type = "text"

	styleName = "background-clip"
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
