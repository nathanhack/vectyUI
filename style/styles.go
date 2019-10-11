package style

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/style/fontFamily"
	"github.com/nathanhack/vectyUI/style/fontName"
	"strings"
)

type BackgroundType string

func Background(color string) BackgroundType {
	return BackgroundType(color)
}

func (bt BackgroundType) Apply(h *vecty.HTML) {
	vecty.Style("background", string(bt)).Apply(h)
}

type ColorType string

func Color(color string) ColorType {
	return ColorType(color)
}

func (ct ColorType) Apply(h *vecty.HTML) {
	vecty.Style("background", string(ct)).Apply(h)
}

func FontFamily(fonts ...fontName.FontNameType) fontFamily.FontFamilyType {
	sb := strings.Builder{}
	for i, f := range fonts {
		sb.WriteString(string(f))
		if i < len(fonts)-1 {
			sb.WriteString(",")
		}
	}
	return fontFamily.FontFamilyType(sb.String())
}
