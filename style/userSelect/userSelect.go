package userSelect

import "github.com/gopherjs/vecty"

type UserSelectType string

const (
	None UserSelectType = "none"
	All  UserSelectType = "all"
	Auto UserSelectType = "auto"
	Text UserSelectType = "text"
)

func (us UserSelectType) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-touch-callout", string(us)).Apply(h)
	vecty.Style("-webkit-user-select", string(string(us))).Apply(h)
	vecty.Style("-khtml-user-select", string(string(us))).Apply(h)
	vecty.Style("-moz-user-select", string(string(us))).Apply(h)
	vecty.Style("-ms-user-select", string(string(us))).Apply(h)
	vecty.Style("user-select", string(string(us))).Apply(h)
}
