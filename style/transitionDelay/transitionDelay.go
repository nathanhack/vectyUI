package transitionDelay

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
)

type Type string

const (
	Initial Type = "initial"
	Inherit Type = "inherit"
	Unset   Type = "unset"
)

var styleNames = []string{"-webkit-transition-delay", "transition-delay"}

func (t Type) Apply(h *vecty.HTML) {
	for _, s := range styleNames {
		vecty.Style(s, string(t)).Apply(h)
	}
}

func (t Type) AddTo(m map[string]string) {
	for _, s := range styleNames {
		m[s] = string(t)
	}
}

func Seconds(timeInSeconds interface{}) Type {
	return Type(internal.Stringify(timeInSeconds, "s"))
}

func Milliseconds(timeInSeconds interface{}) Type {
	return Type(internal.Stringify(timeInSeconds, "ms"))
}

type Value string

func (v Value) Apply(h *vecty.HTML) {
	for _, s := range styleNames {
		vecty.Style(s, string(v)).Apply(h)
	}
}

func (v Value) AddTo(m map[string]string) {
	for _, s := range styleNames {
		m[s] = string(v)
	}
}
