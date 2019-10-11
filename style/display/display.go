package display

import "github.com/gopherjs/vecty"

type DisplayType string

const (
	Inline           DisplayType = "inline"
	Block            DisplayType = "block"
	Contents         DisplayType = "contents"
	Flex             DisplayType = "flex"
	Grid             DisplayType = "grid"
	InlineBlock      DisplayType = "inline-block"
	InlineFlex       DisplayType = "inline-flex"
	InlineGrid       DisplayType = "inline-grid"
	InlineTable      DisplayType = "inline-table"
	ListItem         DisplayType = "list-item"
	RunIn            DisplayType = "run-in"
	Table            DisplayType = "table"
	TableCaption     DisplayType = "table-caption"
	TableColumnGroup DisplayType = "table-column-group"
	TableHeaderGroup DisplayType = "table-header-group"
	TableFooterGroup DisplayType = "table-footer-group"
	TableRowGroup    DisplayType = "table-row-group"
	TableCell        DisplayType = "table-cell"
	TableColumn      DisplayType = "table-column"
	TableRow         DisplayType = "table-row"
	None             DisplayType = "none"
)

func (dt DisplayType) Apply(h *vecty.HTML) {
	vecty.Style("display", string(dt)).Apply(h)
}
