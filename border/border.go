package border

import (
	"github.com/gopherjs/vecty"
	"strconv"
)

type StyleType string

const (
	Dashed = StyleType("dashed")
	Dotted = StyleType("dotted")
	Double = StyleType("double")
	Groove = StyleType("groove")
	Hidden = StyleType("hidden")
	Inset  = StyleType("inset")
	None   = StyleType("none")
	Outset = StyleType("outset")
	Ridge  = StyleType("ridge")
	Solid  = StyleType("solid")
)

type Style struct {
	Top, Bottom, Right, Left, All StyleType
}

func (s Style) Apply(h *vecty.HTML) {
	sa := string(s.All)
	if sa != "" {
		sa = string(s.Top) + " " + string(s.Right) + " " + string(s.Bottom) + " " + string(s.Left)
	}
	vecty.Style("border-style", sa).Apply(h)
}

type Color struct {
	Top, Bottom, Right, Left, All string
}

func (c Color) Apply(h *vecty.HTML) {
	ca := c.All
	if ca != "" {
		ca = c.Top + " " + c.Right + " " + c.Bottom + " " + c.Left
	}

	vecty.Style("border-color", ca).Apply(h)
}

type Width struct {
	Top, Bottom, Right, Left, All int
}

func (w Width) Apply(h *vecty.HTML) {
	if w.Top > 0 || w.Right > 0 || w.Bottom > 0 || w.Left > 0 {
		top := strconv.FormatInt(int64(w.Top), 10)
		right := strconv.FormatInt(int64(w.Right), 10)
		bottom := strconv.FormatInt(int64(w.Bottom), 10)
		left := strconv.FormatInt(int64(w.Left), 10)

		vecty.Style("border-width", top+"px "+right+"px "+bottom+"px "+left+"px").Apply(h)
	} else {
		vecty.Style("border-width", strconv.FormatInt(int64(w.All), 10)+"px ").Apply(h)
	}
}

type Border struct {
	Width int
	Style StyleType
	Color string
}

func (b Border) Apply(h *vecty.HTML) {
	vecty.Style("border", strconv.FormatInt(int64(b.Width), 10)+"px "+string(b.Style)+" "+b.Color).Apply(h)
}
