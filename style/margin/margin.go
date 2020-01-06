package margin

import (
	"github.com/gopherjs/vecty"
	"strings"
)

type Type string

const (
	Auto    Type = "auto"
	Initial Type = "initial"
	Inherit Type = "inherit"

	styleName = "margin"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

type Value []string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, strings.Join(v, " ")).Apply(h)
}

func (v Value) AddTo(m map[string]string) {
	m[styleName] = strings.Join(v, " ")
}
