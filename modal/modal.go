package modal

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"github.com/nathanhack/vectyUI/internal"
	"github.com/nathanhack/vectyUI/style"
	"github.com/nathanhack/vectyUI/style/background"
	"github.com/nathanhack/vectyUI/style/display"
	"github.com/nathanhack/vectyUI/style/height"
	"github.com/nathanhack/vectyUI/style/left"
	"github.com/nathanhack/vectyUI/style/opacity"
	"github.com/nathanhack/vectyUI/style/overflow"
	"github.com/nathanhack/vectyUI/style/position"
	"github.com/nathanhack/vectyUI/style/top"
	"github.com/nathanhack/vectyUI/style/transitionDuration"
	"github.com/nathanhack/vectyUI/style/transitionTimingFunction"
	"github.com/nathanhack/vectyUI/style/visibility"
	"github.com/nathanhack/vectyUI/style/width"
	"github.com/nathanhack/vectyUI/style/zIndex"
	"strconv"
)

const (
	defaultModalID = "divModalID"
)

var divModalCount = 0

//modal.Div
//Notes
// 1) GrabFocus is a component's string ID, if specified upon ShowDiv() it will try and have it grab focus.
// 2) Dismissible if true will set click and keydown listeners on the background to hide them if ESC is pressed or the background is clicked
// 3) ID must be globally unique, if not provided one will be created.
type Div struct {
	vecty.Core
	ID                 string                                 `vecty:"prop"`
	Background         background.Value                       `vecty:"prop"`
	Display            func(modal *Div) vecty.ComponentOrHTML `vecty:"prop"`
	TransitionDuration transitionDuration.Value               `vecty:"prop"`
	Extras             []vecty.Applyer                        `vecty:"prop"`
	Dismissible        bool                                   `vecty:"prop"`
	GrabFocus          string                                 `vecty:"prop"`
	Show               bool                                   `vecty:"prop"`
	hiddenID           string
}

func (div *Div) ShowDiv() {
	div.Show = true
	vecty.Rerender(div)
	if div.GrabFocus != "" {
		go func() { internal.RequestFocusEvent(div.GrabFocus) }()
	}
}

func (div *Div) Hide() {
	div.Show = false
	vecty.Rerender(div)
}

func (div *Div) Render() vecty.ComponentOrHTML {
	if div.ID == "" {
		div.ID = defaultModalID + strconv.Itoa(divModalCount)
		divModalCount++
	}
	if div.hiddenID == "" {
		div.hiddenID = div.ID + "hidden"
	}

	vis := visibility.Hidden
	op := opacity.Number(0)
	if div.Show {
		vis = visibility.Visible
		op = opacity.Number(1)
	}

	markups := []vecty.Applyer{
		prop.ID(div.ID),
		display.Block,
		position.Fixed,
		zIndex.Number(1),
		left.Pixels(0),
		top.Pixels(0),
		width.Percent(100),
		height.Percent(100),
		overflow.Auto,
		style.Background(div.Background),
		vecty.Attribute("tabindex", -1),
		vis,
		op,
		style.Transitions(
			style.Transition("visibility", transitionDuration.Seconds(0)),
			style.Transition("opacity", div.TransitionDuration, transitionTimingFunction.Linear),
		),
		vecty.MarkupIf(div.Dismissible,
			event.Click(func(v *vecty.Event) { div.Hide() }),
			event.KeyUp(func(v *vecty.Event) {
				if !v.Value.IsNull() &&
					v.Value.Get("code").String() == "Escape" {
					div.Hide()
				}
			}),
		),
	}

	markups = append(markups, div.Extras...)

	return elem.Div(
		vecty.Markup(markups...),
		div.Display(div),
	)
}
