package alignItems

import "github.com/gopherjs/vecty"

type Type string

const (
	Stretch   Type = "stretch"
	Center    Type = "center"
	FlexStart Type = "flex-start"
	FlexEnd   Type = "flex-end"
	Baseline  Type = "baseline"
	Initial   Type = "initial"
	Inherit   Type = "inherit"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-align-items", string(t)).Apply(h)
	vecty.Style("align-items", string(t)).Apply(h)
}
