package border

import "github.com/gopherjs/vecty"

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

func (s *Style) StyleName() string {
	return "border-style"
}

//ToString is the javascript version
// of tostring to be used in css
func (s *Style) ToString() string {
	if s.All != "" {
		return string(s.All)
	}
	return string(s.Top) + " " + string(s.Right) + " " + string(s.Bottom) + " " + string(s.Left)
}

func (s Style) Apply(h *vecty.HTML) {
	vecty.Style("border-style", s.ToString()).Apply(h)
}

type Color string

func (c *Color) StyleName() string {
	return "border-color"
}

func (c *Color) ToString() string {
	return string(*c)
}

func (c Color) Apply(h *vecty.HTML) {
	vecty.Style("border-color", c.ToString()).Apply(h)
}

type Width struct {
	Top, Bottom, Right, Left, All string
}

func (b *Width) StyleName() string {
	return "border-width"
}

//ToString is the javascript version
// of tostring to be used in css
func (b *Width) ToString() string {
	if b.All != "" {
		return string(b.All)
	}
	return b.Top + " " + b.Right + " " + b.Bottom + " " + b.Left
}

func (w Width) Apply(h *vecty.HTML) {
	vecty.Style("border-width", w.ToString()).Apply(h)
}
