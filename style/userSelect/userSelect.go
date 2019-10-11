package userSelect

import "github.com/gopherjs/vecty"

type Type string

const (
	None Type = "none"
	All  Type = "all"
	Auto Type = "auto"
	Text Type = "text"
)

func (us Type) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-touch-callout", string(us)).Apply(h)
	vecty.Style("-webkit-user-select", string(string(us))).Apply(h)
	vecty.Style("-khtml-user-select", string(string(us))).Apply(h)
	vecty.Style("-moz-user-select", string(string(us))).Apply(h)
	vecty.Style("-ms-user-select", string(string(us))).Apply(h)
	vecty.Style("user-select", string(string(us))).Apply(h)
}
