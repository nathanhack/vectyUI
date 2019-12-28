package boxShadow

import (
	"github.com/gopherjs/vecty"
	"strings"
)

type Type string

const (
	None    Type = "none"
	Inset   Type = "inset"
	Unset   Type = "unset"
	Initial Type = "initial"
	Inherit Type = "inherit"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("box-shadow", string(t)).Apply(h)
}

type Value []string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("box-shadow", strings.Join(v, ",")).Apply(h)
}
