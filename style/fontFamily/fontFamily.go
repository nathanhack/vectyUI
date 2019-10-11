package fontFamily

import (
	"github.com/gopherjs/vecty"
	"strings"
)

type Value []string

func (fft Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for i, c := range fft {
		sb.WriteString(c)
		if i < len(fft)-1 {
			sb.WriteString(",")
		}
	}
	vecty.Style("font-family", sb.String()).Apply(h)
}
