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

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-align-self", string(t)).Apply(h)
	vecty.Style("align-self", string(t)).Apply(h)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-align-self", string(v)).Apply(h)
	vecty.Style("align-self", string(v)).Apply(h)
}
