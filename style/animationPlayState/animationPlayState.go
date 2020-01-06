package animationPlayState

import "github.com/gopherjs/vecty"

type Type string

const (
	Paused  Type = "paused"
	Running Type = "running"
	Initial Type = "initial"
	Inherit Type = "inherit"

	styleName = "animation-play-state"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

type Value string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}
