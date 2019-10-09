package margin

import (
	"github.com/gopherjs/vecty"
	"strconv"
)

type Weights struct {
	Left, Top, Right, Bottom, All int
	left, top, right, bottom, all string
}

func New(top, right, bottom, left int) Weights {
	return Weights{
		top:    strconv.FormatInt(int64(top), 10) + "px ",
		right:  strconv.FormatInt(int64(right), 10) + "px ",
		bottom: strconv.FormatInt(int64(bottom), 10) + "px ",
		left:   strconv.FormatInt(int64(left), 10) + "px",
		Top:    top,
		Right:  right,
		Bottom: bottom,
		Left:   left,
	}
}

func NewAll(padding int) Weights {
	return Weights{
		All: padding,
		all: strconv.FormatInt(int64(padding), 10) + "px",
	}
}

func (w Weights) StyleName() string {
	return "margin"
}

//ToString is a js toString() function that
// can be used to directly pass the value to js for rendering
func (w Weights) ToString() string {
	// for the order is standardized as the following:
	//margin: 10px 5px 15px 20px;
	//	top margin is 10px
	//	right margin is 5px
	//	bottom margin is 15px
	//	left margin is 20px
	if w.all != "" {
		return w.all
	}
	return w.top + " " + w.right + " " + w.bottom + " " + w.left
}

func (w Weights) Apply(h *vecty.HTML) {
	vecty.Style("margin", w.ToString()).Apply(h)
}
