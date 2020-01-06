package flexWrap

import "github.com/gopherjs/vecty"

type Type string

const (
	Nowrap      Type = "nowrap"
	Wrap        Type = "wrap"
	WrapReverse Type = "wrap-reverse"
	Initial     Type = "initial"
	Inherit     Type = "inherit"
)

var styleNames = []string{"-webkit-flex-wrap", "flex-wrap"}

func (t Type) Apply(h *vecty.HTML) {
	for _, s := range styleNames {
		vecty.Style(s, string(t)).Apply(h)
	}
}

func (t Type) AddTo(m map[string]string) {
	for _, s := range styleNames {
		m[s] = string(t)
	}
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	for _, s := range styleNames {
		vecty.Style(s, string(v)).Apply(h)
	}
}

func (v Value) AddTo(m map[string]string) {
	for _, s := range styleNames {
		m[s] = string(v)
	}
}
