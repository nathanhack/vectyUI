package backgroundOrigin

import "github.com/gopherjs/vecty"

type Type string

const (
	BorderBox  Type = "border-box"
	PaddingBox Type = "padding-box"
	ContentBox Type = "content-box"

	styleName = "background-origin"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}
