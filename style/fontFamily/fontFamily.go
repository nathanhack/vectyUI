package fontFamily

import (
	"github.com/gopherjs/vecty"
	"strings"
)

type Type string

const (
	//generic fonts
	Serif     Type = "serif"
	SansSerif Type = "sans-serif"
	Monospace Type = "monospace"
	Cursive   Type = "cursive"
	Fantasy   Type = "fantasy"

	//specific sans-serif fonts from https://www.w3.org/Style/Examples/007/fonts.en.html
	Arial       Type = "Arial"
	Helvetica   Type = "Helvetica"
	Verdana     Type = "Verdana"
	TrebuchetMS Type = "Trebuchet MS"
	GillSans    Type = "Gill Sans"
	NotoSans    Type = "Noto Sans"
	Avantgarde  Type = "Avantgarde"
	Optima      Type = "Optima"
	ArialNarrow Type = "Arial Narrow"

	//specific serif fonts from https://www.w3.org/Style/Examples/007/fonts.en.html
	Times                Type = "Times"
	TimesNewRoman        Type = "Times New Roman"
	Didot                Type = "Didot"
	Georgia              Type = "Georgia"
	Palatino             Type = "Palatino"
	Bookman              Type = "Bookman"
	NewCenturySchoolbook Type = "New Century Schoolbook"
	AmericanTypewriter   Type = "American Typewriter"

	//specific monospace fonts from https://www.w3.org/Style/Examples/007/fonts.en.html
	AndaleMono     Type = "Andale Mono"
	CourierNew     Type = "Courier New"
	Courier        Type = "Courier"
	FreeMono       Type = "FreeMono"
	OCRAStd        Type = "OCR A Std"
	DejaVuSansMono Type = "DejaVu Sans Mono"

	//specific cursive fonts from https://www.w3.org/Style/Examples/007/fonts.en.html
	ComicSansMS    Type = "Comic Sans MS"
	ComicSans      Type = "Comic Sans"
	AppleChancery  Type = "Apple Chancery"
	BradleyHand    Type = "Bradley Hand"
	BrushScriptMT  Type = "Brush Script MT"
	BrushScriptStd Type = "Brush Script Std"
	SnellRoundhand Type = "Snell Roundhand"

	//specific fantasy fonts from https://www.w3.org/Style/Examples/007/fonts.en.html
	Impact      Type = "Impact"
	Luminari    Type = "Luminari"
	Chalkduster Type = "Chalkduster"
	JazzLET     Type = "Jazz LET"
	Blippo      Type = "Blippo"
	StencilStd  Type = "Stencil Std"
	MarkerFelt  Type = "Marker Felt"
	Trattatello Type = "Trattatello"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("font-family", string(t)).Apply(h)
}

type Value []string

func (v Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for i, c := range v {
		sb.WriteString(c)
		if i < len(v)-1 {
			sb.WriteString(",")
		}
	}
	vecty.Style("font-family", sb.String()).Apply(h)
}
