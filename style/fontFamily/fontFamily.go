package fontFamily

import (
	"github.com/gopherjs/vecty"
	"strings"
)

type Value []string

func (v Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for i, c := range v {
		sb.WriteString(c)
		if i < len(v)-1 {
			sb.WriteString(",")
		}
	}
	vecty.Style("font-family", sb.String()).Apply(h)
}
