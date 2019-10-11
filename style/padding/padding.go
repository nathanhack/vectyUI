package padding

import (
	"github.com/gopherjs/vecty"
	"strconv"
)

type Weights struct {
	Left, Top, Right, Bottom      int
	left, top, right, bottom, all string
}

func New(top, right, bottom, left int) Weights {
	return Weights{
		top:    strconv.FormatInt(int64(top), 10) + "px ",
		right:  strconv.FormatInt(int64(right), 10) + "px ",
		bottom: strconv.FormatInt(int64(bottom), 10) + "px ",
		left:   strconv.FormatInt(int64(left), 10) + "px ",
		Top:    top,
		Right:  right,
		Bottom: bottom,
		Left:   left,
	}
}

func NewAll(padding int) Weights {
	return Weights{
		all:    strconv.FormatInt(int64(padding), 10) + "px",
		Top:    padding,
		Right:  padding,
		Bottom: padding,
		Left:   padding,
	}
}

func (w Weights) StyleName() string {
	return "padding"
}

//ToString is a js toString() function that
// can be used to directly pass the value to js for rendering
func (w Weights) ToString() string {
	// for the order is standardized as the following:
	//padding:10px 5px 15px 20px
	//	top padding is 10px
	//	right padding is 5px
	//	bottom padding is 15px
	//	left padding is 20px
	if w.all != "" {
		return w.all
	}
	return w.top + " " + w.right + " " + w.bottom + " " + w.left
}

func (w Weights) Apply(h *vecty.HTML) {
	vecty.Style("padding", w.ToString()).Apply(h)
}
