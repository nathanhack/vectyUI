package flexDirection

import "github.com/gopherjs/vecty"

type Type string

const (
	Row           Type = "row"
	RowReverse    Type = "row-reverse"
	Column        Type = "column"
	ColumnReverse Type = "column-reverse"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-flex-direction", string(t)).Apply(h)
	vecty.Style("flex-direction", string(t)).Apply(h)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-flex-direction", string(v)).Apply(h)
	vecty.Style("flex-direction", string(v)).Apply(h)
}
