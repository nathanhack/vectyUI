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

	styleName = "background-repeat"
)

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

type Value []Type

func (v Value) AddTo(m map[string]string) {
	sb := strings.Builder{}
	for i, t := range v {
		sb.WriteString(string(t))
		if i < len(v)-1 {
			sb.WriteString(",")
		}
	}
	m[styleName] = sb.String()
}

func (v Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for i, t := range v {
		sb.WriteString(string(t))
		if i < len(v)-1 {
			sb.WriteString(",")
		}
	}
	vecty.Style(styleName, sb.String()).Apply(h)
}
