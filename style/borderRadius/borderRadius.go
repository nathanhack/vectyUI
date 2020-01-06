package borderRadius

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
	"strings"
)

type Type string

const (
	Initial Type = "initial"
	Inherit Type = "inherit"
	Unset   Type = "unset"

	styleName = "border-radius"
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

func Pixels(pixels interface{}) Type {
	return Type(internal.Stringify(pixels, "px"))
}

func Ems(ems interface{}) Type {
	return Type(internal.Stringify(ems, "em"))
}

type Value []Type

func (v Value) AddTo(m map[string]string) {
	sb := strings.Builder{}
	for _, l := range v {
		sb.WriteString(string(l) + " ")
	}
	m[styleName] = sb.String()
}

func (v Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for _, l := range v {
		sb.WriteString(string(l) + " ")
	}
	vecty.Style(styleName, sb.String()).Apply(h)
}
