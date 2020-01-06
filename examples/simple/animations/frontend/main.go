package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/nathanhack/vectyUI/style"
	"github.com/nathanhack/vectyUI/style/animationDirection"
	"github.com/nathanhack/vectyUI/style/animationIterationCount"
	"github.com/nathanhack/vectyUI/style/keyframes"
)

//go:generate env GOARCH=wasm GOOS=js go build -o ../static/main.wasm

type Body struct {
	vecty.Core
}

//to ensure keyframes are created only once we make them a global
var bodyKeyframe = keyframes.Keyframes{
	Name: "body",
	Frames: []keyframes.Keyframe{
		{Sel: keyframes.Percent(0), Styles: []keyframes.KeyframeStyler{style.Background("red")}},
		{Sel: keyframes.Percent(100), Styles: []keyframes.KeyframeStyler{style.Background("blue")}},
	},
}

var divKeyframe1 = keyframes.Keyframes{
	Frames: []keyframes.Keyframe{
		{Sel: keyframes.Percent(0), Styles: []keyframes.KeyframeStyler{style.Background("black")}},
		{Sel: keyframes.Percent(100), Styles: []keyframes.KeyframeStyler{style.Background("white")}},
	},
}

var divKeyframe2 = keyframes.Keyframes{
	Frames: []keyframes.Keyframe{
		{Sel: keyframes.Percent(0), Styles: []keyframes.KeyframeStyler{style.FontSize(18)}},
		{Sel: keyframes.Percent(100), Styles: []keyframes.KeyframeStyler{style.FontSize(25)}},
	},
}

func (b *Body) Mount() {
	vecty.Rerender(b)

	bodyKeyframe.Mount() //keyframes must be mounted
	divKeyframe1.Mount()
	divKeyframe2.Mount()
}

func (b Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(
			style.BackgroundColor("black"),
			style.AnimationName(bodyKeyframe.Name),
			style.AnimationDuration(3),
			animationIterationCount.Infinite,
			animationDirection.Alternate,
		),

		elem.Div(
			vecty.Markup(
				style.Width(100),
				style.Height(100),
				style.AnimationNames(divKeyframe1.Name, divKeyframe2.Name),
				style.AnimationDurations(3, 3),
				style.AnimationIterationCounts(animationIterationCount.Infinite, animationIterationCount.Infinite),
				style.AnimationDirections(animationDirection.Alternate, animationDirection.Alternate),
			),
			vecty.Text("Hello There"),
		),
	)
}

func main() {
	vecty.SetTitle("Simple Example")
	vecty.RenderBody(&Body{})
}
