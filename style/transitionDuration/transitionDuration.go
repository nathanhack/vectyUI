package transitionDuration

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

var styleNames = []string{"-webkit-transition-duration", "transition-duration"}

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

func Milliseconds(timeInMilliseconds interface{}) Type {
	return Type(internal.Stringify(timeInMilliseconds, "ms"))
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
