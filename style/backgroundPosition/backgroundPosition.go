package backgroundPosition

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
	"strings"
)

type Type string

const (
	Left   Type = "left"
	Center Type = "center"
	Right  Type = "right"
	Top    Type = "top"
	Bottom Type = "bottom"

	styleName = "background-position"
)

func Len(length interface{}) Type {
	return Type(internal.Stringify(length, "px"))
}

func Percent(percent interface{}) Type {
	return Type(internal.Stringify(percent, "%"))
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
	vecty.Style(styleName, sb.String()).Apply(h)
}

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

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
