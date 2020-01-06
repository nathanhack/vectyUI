package borderWidth

import (
	"github.com/gopherjs/vecty"
	"strings"
)

type Type string

const (
	Medium  Type = "medium"
	Thin    Type = "thin"
	Thick   Type = "thick"
	Initial Type = "initial"
	Inherit Type = "inherit"

	styleName = "border-width"
)

func (t Type) AddTo(m map[string]string) {
	m[styleName] = string(t)
}

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style(styleName, string(t)).Apply(h)
}

//Value is slice of string widths, CSS supports 1 to 4
// case 1( width ): top bottom right left = width
// case 2( w1 w2 ): top bottom = w1, right left = w2
// case 3( w1 w2 w3 ): top = w1, right left = w2, bottom = w3
// case 4( w1 w2 w3 w4 ) : top = w1, right = w2, bottom = w3, left = w4
type Value []string

func (v Value) AddTo(m map[string]string) {
	sb := strings.Builder{}
	for _, w := range v {
		sb.WriteString(w + " ")
	}
	m[styleName] = sb.String()
}

func (v Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for _, w := range v {
		sb.WriteString(w + " ")
	}
	vecty.Style(styleName, sb.String()).Apply(h)
}
