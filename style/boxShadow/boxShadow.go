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

	styleName = "box-shadow"
)

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

type Value []string

func (v Value) AddTo(m map[string]string) {
	m[styleName] = strings.Join(v, ",")
}

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, strings.Join(v, ",")).Apply(h)
}
