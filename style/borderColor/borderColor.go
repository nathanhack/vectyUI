package borderColor

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/color"
	"strings"
)

//Value is slice of string colors, CSS supports 1 to 4
// case 1( color ): top bottom right left = color
// case 2( c1 c2 ): top bottom = c1, right left = c2
// case 3( c1 c2 c3 ): top = c1, right left = c2, bottom = c3
// case 4( c1 c2 c3 c4 ) : top = c1, right = c2, bottom = c3, left = c4
type Value []color.Type

func (v Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for _, c := range v {
		sb.WriteString(string(c) + " ")
	}
	vecty.Style("border-color", sb.String()).Apply(h)
}
