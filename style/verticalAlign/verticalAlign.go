package verticalAlign

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
)

type Type string

const (
	Baseline   Type = "baseline"
	Sub        Type = "sub"
	Super      Type = "super"
	Top        Type = "top"
	TextTop    Type = "text-top"
	Middle     Type = "middle"
	Bottom     Type = "bottom"
	TextBottom Type = "text-bottom"
	Initial    Type = "initial"
	Inherit    Type = "inherit"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("vertical-align", string(t)).Apply(h)
}

func Percent(percent interface{}) Type {
	return Type(internal.Stringify(percent, "%"))
}

func Length(length interface{}) Type {
	return Type(internal.Stringify(length, "px"))
}

type Value Type

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("vertical-align", string(v)).Apply(h)
}
