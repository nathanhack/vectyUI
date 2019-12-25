package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	materialDesignColor "github.com/nathanhack/vectyUI/materialdesign"
	"github.com/nathanhack/vectyUI/style"
	"github.com/nathanhack/vectyUI/style/backgroundImage"
	"github.com/nathanhack/vectyUI/style/backgroundImage/colorStop"
	"github.com/nathanhack/vectyUI/style/backgroundPosition"
	"github.com/nathanhack/vectyUI/style/backgroundRepeat"
	"github.com/nathanhack/vectyUI/style/backgroundSize"
)

//go:generate env GOARCH=wasm GOOS=js go build -o ../static/main.wasm

type Body struct {
	vecty.Core
}

func (b Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(
			style.BackgroundColor("black"),
		),
		elem.Div(
			vecty.Markup(
				style.Height(100),
				style.Width(100),
				backgroundImage.LinearGradient(backgroundImage.ToBottom,
					colorStop.Type(materialDesignColor.Amber900),
					colorStop.From(materialDesignColor.Blue200),
					"green",
				),
			),
		),

		elem.Div(
			vecty.Markup(
				style.Height(100),
				style.Width(100),
				style.BackgroundImage(
					backgroundImage.URL("image1.png"),
					backgroundImage.LinearGradient(backgroundImage.ToBottom, "grey", "20%", "red", colorStop.Percent("blue", 30), colorStop.HintPercent(60), "green"),
				),
				style.BackgroundRepeat(backgroundRepeat.NoRepeat, backgroundRepeat.NoRepeat),
				style.BackgroundPosition(backgroundPosition.Center, backgroundPosition.Center),
			),
		),
		elem.Div(
			vecty.Markup(
				style.Height(100),
				style.Width(100),
				backgroundImage.LinearGradient(backgroundImage.Degree(20),
					colorStop.From(materialDesignColor.Amber200),
					colorStop.From(materialDesignColor.Blue900),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				style.Height(100),
				style.Width(100),
				style.Backgrounds(
					style.Background(
						backgroundPosition.Center,  //note order matters
						"/",                        // //note order matters and this is required
						backgroundSize.Percent(75), //note order matters
						backgroundImage.URL("image1.png"),
						backgroundRepeat.NoRepeat,
					),
					style.Background(
						"left", //note order matters
						"/25%", //note order matters to center this image
						backgroundRepeat.NoRepeat,
						backgroundImage.URL("image1.png"),
					),
					style.Background(
						backgroundImage.LinearGradient(backgroundImage.ToBottom, "grey", "20%", "red", colorStop.Percent("blue", 30), colorStop.HintPercent(60), "green"),
					),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				style.Height(100),
				style.Width(100),
				backgroundImage.URL("image1.png"),
				backgroundSize.Percent(50),
				backgroundRepeat.NoRepeat,
				style.BackgroundPosition(backgroundPosition.Center),
				style.BackgroundColor("grey"),
			),
		),
	)
}

func main() {
	vecty.SetTitle("Simple Example")
	vecty.RenderBody(&Body{})
}
