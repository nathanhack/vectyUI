package flexDirection

import "github.com/gopherjs/vecty"

type FlexDirectionType string

const (
	Row           FlexDirectionType = "row"
	RowReverse    FlexDirectionType = "row-reverse"
	Column        FlexDirectionType = "column"
	ColumnReverse FlexDirectionType = "column-reverse"
)

func (fdt FlexDirectionType) Apply(h *vecty.HTML) {
	vecty.Style("-webkit-flex-direction", string(fdt)).Apply(h)
	vecty.Style("flex-direction", string(fdt)).Apply(h)
}
