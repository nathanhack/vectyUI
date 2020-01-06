package userSelect

import "github.com/gopherjs/vecty"

type Type string

const (
	None Type = "none"
	All  Type = "all"
	Auto Type = "auto"
	Text Type = "text"
)

var styleNames = []string{
	"-webkit-user-select",
	"-khtml-user-select",
	"-moz-user-select",
	"-ms-user-select",
	"user-select",
}

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
