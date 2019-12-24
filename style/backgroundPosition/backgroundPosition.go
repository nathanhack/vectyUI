package backgroundPosition

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/position"
	"strings"
)

type Value []position.Type

func (v Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for i, t := range v {
		sb.WriteString(string(t))
		if i < len(v)-1 {
			sb.WriteString(",")
		}
	}
	vecty.Style("background-position", sb.String()).Apply(h)
}
