package float

import "github.com/gopherjs/vecty"

type Type string

const (
	None  Type = "none"
	Left  Type = "left"
	Right Type = "right"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("float", string(t)).Apply(h)
}
