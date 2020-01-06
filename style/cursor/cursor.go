package cursor

import (
	"github.com/gopherjs/vecty"
)

type Type string

const (
	Alias        Type = "alias"
	AllScroll    Type = "all-scroll"
	Auto         Type = "auto"
	Cell         Type = "cell"
	ContextMenu  Type = "context-menu"
	ColResize    Type = "col-resize"
	Copy         Type = "copy"
	Crosshair    Type = "crosshair"
	Default      Type = "default"
	EResize      Type = "e-resize"
	EwResize     Type = "ew-resize"
	Grab         Type = "grab"
	Grabbing     Type = "grabbing"
	Help         Type = "help"
	Move         Type = "move"
	NResize      Type = "n-resize"
	NeResize     Type = "ne-resize"
	NeswResize   Type = "nesw-resize"
	NsResize     Type = "ns-resize"
	NwResize     Type = "nw-resize"
	NwseResize   Type = "nwse-resize"
	NoDrop       Type = "no-drop"
	None         Type = "none"
	NotAllowed   Type = "not-allowed"
	Pointer      Type = "pointer"
	Progress     Type = "progress"
	RowResize    Type = "row-resize"
	SResize      Type = "s-resize"
	SeResize     Type = "se-resize"
	SwResize     Type = "sw-resize"
	Text         Type = "text"
	VerticalText Type = "vertical-text"
	WResize      Type = "w-resize"
	Wait         Type = "wait"
	ZoomIn       Type = "zoom-in"
	ZoomOut      Type = "zoom-out"

	Initial Type = "initial"
	Inherit Type = "inherit"

	styleName = "cursor"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}

func (v Value) AddTo(m map[string]string) {
	m[styleName] = string(v)
}
