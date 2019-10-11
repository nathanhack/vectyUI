package fontName

type FontNameType string

const (
	//generic fonts
	Serif     FontNameType = "serif"
	SansSerif FontNameType = "sans-serif"
	Monospace FontNameType = "monospace"
	Cursive   FontNameType = "cursive"
	Fantasy   FontNameType = "fantasy"

	//specific sans-serif fonts from https://www.w3.org/Style/Examples/007/fonts.en.html
	Arial       FontNameType = "Arial"
	Helvetica   FontNameType = "Helvetica"
	Verdana     FontNameType = "Verdana"
	TrebuchetMS FontNameType = "Trebuchet MS"
	GillSans    FontNameType = "Gill Sans"
	NotoSans    FontNameType = "Noto Sans"
	Avantgarde  FontNameType = "Avantgarde"
	Optima      FontNameType = "Optima"
	ArialNarrow FontNameType = "Arial Narrow"

	//specific serif fonts from https://www.w3.org/Style/Examples/007/fonts.en.html
	Times                FontNameType = "Times"
	TimesNewRoman        FontNameType = "Times New Roman"
	Didot                FontNameType = "Didot"
	Georgia              FontNameType = "Georgia"
	Palatino             FontNameType = "Palatino"
	Bookman              FontNameType = "Bookman"
	NewCenturySchoolbook FontNameType = "New Century Schoolbook"
	AmericanTypewriter   FontNameType = "American Typewriter"

	//specific monospace fonts from https://www.w3.org/Style/Examples/007/fonts.en.html
	AndaleMono     FontNameType = "Andale Mono"
	CourierNew     FontNameType = "Courier New"
	Courier        FontNameType = "Courier"
	FreeMono       FontNameType = "FreeMono"
	OCRAStd        FontNameType = "OCR A Std"
	DejaVuSansMono FontNameType = "DejaVu Sans Mono"

	//specific cursive fonts from https://www.w3.org/Style/Examples/007/fonts.en.html
	ComicSansMS    FontNameType = "Comic Sans MS"
	ComicSans      FontNameType = "Comic Sans"
	AppleChancery  FontNameType = "Apple Chancery"
	BradleyHand    FontNameType = "Bradley Hand"
	BrushScriptMT  FontNameType = "Brush Script MT"
	BrushScriptStd FontNameType = "Brush Script Std"
	SnellRoundhand FontNameType = "Snell Roundhand"

	//specific fantasy fonts from https://www.w3.org/Style/Examples/007/fonts.en.html
	Impact      FontNameType = "Impact"
	Luminari    FontNameType = "Luminari"
	Chalkduster FontNameType = "Chalkduster"
	JazzLET     FontNameType = "Jazz LET"
	Blippo      FontNameType = "Blippo"
	StencilStd  FontNameType = "Stencil Std"
	MarkerFelt  FontNameType = "Marker Felt"
	Trattatello FontNameType = "Trattatello"
)
