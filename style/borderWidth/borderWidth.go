package borderWidth

import (
	"github.com/gopherjs/vecty"
	"strings"
)

//Value is slice of string widths, CSS supports 1 to 4
// case 1( width ): top bottom right left = width
// case 2( w1 w2 ): top bottom = w1, right left = w2
// case 3( w1 w2 w3 ): top = w1, right left = w2, bottom = w3
// case 4( w1 w2 w3 w4 ) : top = w1, right = w2, bottom = w3, left = w4
type Value []string

func (wdith Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for _, w := range wdith {
		sb.WriteString(w + " ")
	}
	vecty.Style("border-width", sb.String()).Apply(h)
}
