package fontSize

import "github.com/gopherjs/vecty"

type Type string

const (
	Medium  Type = "medium"
	XXSmall Type = "xx-small"
	XSmall  Type = "x-small"
	Small   Type = "small"
	Large   Type = "large"
	XLarge  Type = "x-large"
	XXLarge Type = "xx-large"
	Smaller Type = "smaller"
	Larger  Type = "larger"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("font-size", string(t)).Apply(h)
}

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("font-size", string(v)).Apply(h)
}
