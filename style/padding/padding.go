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
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("padding", string(t)).Apply(h)
}

func Percent(percent interface{}) Type {
	return Type(internal.Stringify(percent, "%"))
}

func Length(length interface{}) Type {
	return Type(internal.Stringify(length, "px"))
}

type Value []string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("padding", strings.Join(v, " ")).Apply(h)
}
