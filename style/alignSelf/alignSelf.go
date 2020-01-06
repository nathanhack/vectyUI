package alignSelf

import "github.com/gopherjs/vecty"

type Type string

const (
	Auto      Type = "auto"
	Stretch   Type = "stretch"
	Center    Type = "center"
	FlexStart Type = "flex-start"
	FlexEnd   Type = "flex-end"
	Baseline  Type = "baseline"
	Initial   Type = "initial"
	Inherit   Type = "inherit"
)

var styleNames = []string{"-webkit-align-self", "align-self"}

func (t Type) AddTo(m map[string]string) {
	for _, s := range styleNames {
		m[s] = string(t)
	}
}

func (t Type) Apply(h *vecty.HTML) {
	for _, s := range styleNames {
		vecty.Style(s, string(t)).Apply(h)
	}
}

type Value Type

func (v Value) AddTo(m map[string]string) {
	for _, s := range styleNames {
		m[s] = string(v)
	}
}

func (v Value) Apply(h *vecty.HTML) {
	for _, s := range styleNames {
		vecty.Style(s, string(v)).Apply(h)
	}
}
