package display

import "github.com/gopherjs/vecty"

type Type string

const (
	Inline           Type = "inline"
	Block            Type = "block"
	Contents         Type = "contents"
	Flex             Type = "flex"
	Grid             Type = "grid"
	InlineBlock      Type = "inline-block"
	InlineFlex       Type = "inline-flex"
	InlineGrid       Type = "inline-grid"
	InlineTable      Type = "inline-table"
	ListItem         Type = "list-item"
	RunIn            Type = "run-in"
	Table            Type = "table"
	TableCaption     Type = "table-caption"
	TableColumnGroup Type = "table-column-group"
	TableHeaderGroup Type = "table-header-group"
	TableFooterGroup Type = "table-footer-group"
	TableRowGroup    Type = "table-row-group"
	TableCell        Type = "table-cell"
	TableColumn      Type = "table-column"
	TableRow         Type = "table-row"
	None             Type = "none"
)

func (dt Type) Apply(h *vecty.HTML) {
	vecty.Style("display", string(dt)).Apply(h)
}
