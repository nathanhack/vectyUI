package color

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
)

type Type string

const (
	Initial Type = "initial"
	Inherit Type = "inherit"

	styleName = "color"
)

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

//HSL creates a Type using the Hue-Saturation-Lightness-Alpha model.
//	h		Defines a degree on the color circle (from 0 to 360) - 0 (or 360) is red, 120 is green, 240 is blue
//	s		Defines the saturation; 0% is a shade of gray and 100% is the full color (full saturation)
//	l		Defines the lightness; 0% is black, 50% is normal, and 100% is white
func HSL(h, s, l interface{}) Type {
	return Type("hsl(" + internal.Stringify(h, "") + "," + internal.Stringify(s, "") + "," + internal.Stringify(l, "") + ")")
}

//HSLA creates a Type using the Hue-Saturation-Lightness-Alpha model.
//	h		Defines a degree on the color circle (from 0 to 360) - 0 (or 360) is red, 120 is green, 240 is blue
//	s		Defines the saturation; 0% is a shade of gray and 100% is the full color (full saturation)
//	l		Defines the lightness; 0% is black, 50% is normal, and 100% is white
//	a		Defines the opacity as a number between 0.0 (fully transparent) and 1.0 (fully opaque)
func HSLA(h, s, l, a interface{}) Type {
	return Type("hsla(" + internal.Stringify(h, "") + "," + internal.Stringify(s, "") + "," + internal.Stringify(l, "") + "," + internal.Stringify(a, "") + ")")
}

//RGB creates a Type using the Red-Green-Blue model.
//	r		values must be in range [0,255]
//	g		values must be in range [0,255]
//	b		values must be in range [0,255]
func RGB(r, g, b interface{}) Type {
	return Type("rgba(" + internal.Stringify(r, "") + "," + internal.Stringify(g, "") + "," + internal.Stringify(b, "") + ")")
}

//RGBA creates a Type using the Red-Green-Blue-Alpha model.
//	r		values must be in range [0,255]
//	g		values must be in range [0,255]
//	b		values must be in range [0,255]
//	a		values must be in range [0.0, 1.0]
func RGBA(r, g, b, a interface{}) Type {
	return Type("rgba(" + internal.Stringify(r, "") + "," + internal.Stringify(g, "") + "," + internal.Stringify(b, "") + "," + internal.Stringify(a, "") + ")")
}

type Value Type

func (v Value) AddTo(m map[string]string) {
	m[styleName] = string(v)
}

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(v)).Apply(h)
}
