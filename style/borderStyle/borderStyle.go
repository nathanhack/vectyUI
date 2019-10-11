package borderStyle

import "github.com/gopherjs/vecty"

type Type string

const (
	Dashed = Type("dashed")
	Dotted = Type("dotted")
	Double = Type("double")
	Groove = Type("groove")
	Hidden = Type("hidden")
	Inset  = Type("inset")
	None   = Type("none")
	Outset = Type("outset")
	Ridge  = Type("ridge")
	Solid  = Type("solid")
)

type Style struct {
	Top, Bottom, Right, Left, All Type
}

func (s Style) Apply(h *vecty.HTML) {
	sa := string(s.All)
	if sa != "" {
		sa = string(s.Top) + " " + string(s.Right) + " " + string(s.Bottom) + " " + string(s.Left)
	}
	vecty.Style("border-style", sa).Apply(h)
}
