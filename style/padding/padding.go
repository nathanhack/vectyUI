package padding

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
	"strings"
)

type Type string

const (
	Initial Type = "initial"
	Inherit Type = "inherit"

	styleName = "padding"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

func Percent(percent interface{}) Type {
	return Type(internal.Stringify(percent, "%"))
}

func Pixels(length interface{}) Type {
	return Type(internal.Stringify(length, "px"))
}

type Value []string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, strings.Join(v, " ")).Apply(h)
}

func (v Value) AddTo(m map[string]string) {
	m[styleName] = strings.Join(v, " ")

}
