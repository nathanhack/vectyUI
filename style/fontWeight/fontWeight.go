package fontWeight

import "github.com/gopherjs/vecty"

type Type string

const (
	normal    Type = "normal"
	bold      Type = "bold"
	bolder    Type = "bolder"
	lighter   Type = "lighter"
	Weight100 Type = "100"
	Weight200 Type = "200"
	Weight300 Type = "300"
	Weight400 Type = "400"
	Weight500 Type = "500"
	Weight600 Type = "600"
	Weight700 Type = "700"
	Weight800 Type = "800"
	Weight900 Type = "900"
	Initial   Type = "initial"
	Inherit   Type = "inherit"

	styleName = "font-weight"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}

func (v Value) AddTo(m map[string]string) {
	m[styleName] = string(v)
}
