package border

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/style/borderStyle"
	"strconv"
)

type Border struct {
	Width int
	Style borderStyle.Type
	Color string
}

func (b Border) Apply(h *vecty.HTML) {
	vecty.Style("border", strconv.FormatInt(int64(b.Width), 10)+"px "+string(b.Style)+" "+b.Color).Apply(h)
}
