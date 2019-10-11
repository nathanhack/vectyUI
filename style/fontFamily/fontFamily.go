package fontFamily

import "github.com/gopherjs/vecty"

type FontFamilyType string

func (fft FontFamilyType) Apply(h *vecty.HTML) {
	vecty.Style("font-family", string(fft)).Apply(h)
}
