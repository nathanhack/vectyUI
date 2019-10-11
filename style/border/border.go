package border

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/style/borderStyle"
	"strconv"
)

type Value struct {
	Width int
	Style borderStyle.Type
	Color string
}

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("border", strconv.FormatInt(int64(v.Width), 10)+"px "+string(v.Style)+" "+v.Color).Apply(h)
}
