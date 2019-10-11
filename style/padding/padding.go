package padding

import (
	"github.com/gopherjs/vecty"
	"strings"
)

type Value []string

func (v Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for _, l := range v {
		sb.WriteString(l + " ")
	}
	println("Padding Value:" + sb.String())
	vecty.Style("padding", sb.String()).Apply(h)
}
