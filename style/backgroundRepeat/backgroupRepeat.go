package backgroundRepeat

import (
	"github.com/gopherjs/vecty"
	"strings"
)

type Type string

const (
	Repeat   Type = "repeat"
	RepeatX  Type = "repeat"
	RepeatY  Type = "repeat-y"
	NoRepeat Type = "no-repeat"
	Space    Type = "space"
	Round    Type = "round"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("background-repeat", string(t)).Apply(h)
}

type Value []Type

func (v Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for i, t := range v {
		sb.WriteString(string(t))
		if i < len(v)-1 {
			sb.WriteString(",")
		}
	}
	vecty.Style("background-repeat", sb.String()).Apply(h)
}
