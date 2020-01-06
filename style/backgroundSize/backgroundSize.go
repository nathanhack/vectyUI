package backgroundSize

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
	"strings"
)

type Type string

const (
	Cover   Type = "cover"
	Contain Type = "contain"
	Auto    Type = "auto"
	Initial Type = "initial"
	Inherit Type = "inherit"

	styleName = "background-size"
)

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

func Percent(percent interface{}) Type {
	return Type(internal.Stringify(percent, "%"))
}

func Pixels(length interface{}) Type {
	return Type(internal.Stringify(length, "px"))
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
