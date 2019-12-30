package sel

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"github.com/nathanhack/vectyUI/internal"
	"github.com/nathanhack/vectyUI/style"
	"github.com/nathanhack/vectyUI/style/alignItems"
	"github.com/nathanhack/vectyUI/style/display"
	"github.com/nathanhack/vectyUI/style/flexWrap"
	"github.com/nathanhack/vectyUI/style/justifyContent"
	"github.com/nathanhack/vectyUI/style/position"
	"github.com/nathanhack/vectyUI/style/visibility"
	"github.com/nathanhack/vectyUI/style/zIndex"
	"strconv"
	"syscall/js"
	"time"
)

//GenericOption is effectively a div container for the option
// Notes:
// 1) User should not directly create the following events: MouseMove and Click.
// 2) Each call to GenericOption.Option() should return a distinct vecty.ComponentOrHTML.
type GenericOption struct {
	vecty.Core
	Option     func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML `vecty:"prop"`
	MouseMove  func(v *vecty.Event)                                                     `vecty:"prop"`
	Disabled   bool                                                                     `vecty:"prop"`
	Hightlight bool                                                                     `vecty:"prop"`
	Position   int                                                                      `vecty:"prop"`
}

func (g *GenericOption) Render() vecty.ComponentOrHTML {
	return g.Option(g.Disabled, g.Hightlight, false)
}

const genericSelectID = "GenericSelectID"

var genericSelectCount = 0

var DefaultSelectedOptionFocusedStyles = []vecty.Applyer{
	style.BoxShadow(0, 0, 4, 2, "#a8d1ff"),
}

var DefaultOptionsStyles = []vecty.Applyer{
	zIndex.Number(1),
	style.BoxShadow(0, 3, 5, -1, "black"),
}

var DefaultSelectedOptionButtonAlignmentStyles = []vecty.Applyer{
	position.Relative,
	display.Flex,
	justifyContent.SpaceBetween,
	flexWrap.Nowrap,
	alignItems.Center,
}

//Generic is a generic select
// It is suggested that the user make the placeholder and all options the same dimensions.
// Use .SelectedPos (<=0 undefined, 1 zeroth option index) to set the starting selected option
// position and to retrieve selected (if not using SelectedEvent).
// Notes:
// 1) User should NOT create the following events (applies to everything): MouseEnter,MouseLeave and Click.
// 2) If a *Styles is specified it will replace default styles. (The defaults are public for user to build from)
// 3) The ID must be globally unique, if not assigned a unique ID will be generated.
// 4) The following should be avoided in styles or components: position, display and justifyContent.
// 5) To create without a button, set the Button to and empty Div.
// 5) Normally disabled options are unselectable, at startup this can be overridden by setting .SelectPos (use carefully).
type Generic struct {
	vecty.Core
	ID                                  string                                                                  `vecty:"prop"`
	Options                             []*GenericOption                                                        `vecty:"prop"`
	Placeholder                         GenericOption                                                           `vecty:"prop"`
	Styles                              []vecty.Applyer                                                         `vecty:"prop"`
	SelectedOptionStyles                []vecty.Applyer                                                         `vecty:"prop"`
	SelectedOptionButtonAlignmentStyles []vecty.Applyer                                                         `vecty:"prop"`
	SelectedOptionFocusedStyles         []vecty.Applyer                                                         `vecty:"prop"`
	OptionsStyles                       []vecty.Applyer                                                         `vecty:"prop"`
	Button                              func(disabled, mouseOver, focused, expanded bool) vecty.ComponentOrHTML `vecty:"prop"`
	SelectedEvent                       func(index int, g *Generic)                                             `vecty:"prop"`
	SelectedPos                         int                                                                     `vecty:"prop"`
	MouseEnter                          func(v *vecty.Event)                                                    `vecty:"prop"`
	MouseLeave                          func(v *vecty.Event)                                                    `vecty:"prop"`
	Disabled                            bool                                                                    `vecty:"prop"` //disables all interactions
	expanded                            bool                                                                    // will render option list visible
	blurRerender                        bool                                                                    // internal state used to help communicate events
	focused                             bool                                                                    // will render component as focused
	selectID                            string                                                                  // used to identify the hidden select
	mouseOver                           bool                                                                    // used to identify when mouse is over placeholder/selected or button
}

func (g *Generic) Render() vecty.ComponentOrHTML {
	//since this is effectively a custom component
	// it means we lose all the niceties line focus
	// and cursor key interaction
	// to remedy this we create a hidden
	// select and connect up all the focus events
	// between our custom one and the hidden one
	if g.ID == "" {
		g.ID = genericSelectID + strconv.Itoa(genericSelectCount)
		genericSelectCount++
	}
	if g.selectID == "" {
		g.selectID = g.ID + "SelectHidden"
	}

	markups := make([]vecty.Applyer, 0)

	markups = append(markups, prop.ID(g.ID))

	markups = append(markups, g.Styles...)

	return elem.Div(
		vecty.Markup(markups...),
		g.makeHiddenSelect(),
		g.makeSelectedOptionOrPlaceholder(),
	)
}

func (g *Generic) makeSelectedOptionOrPlaceholder() vecty.ComponentOrHTML {

	markups := make([]vecty.Applyer, 0)

	markups = append(markups, position.Relative)

	switch {
	case g.focused && len(g.SelectedOptionFocusedStyles) > 0:
		markups = append(markups, g.SelectedOptionFocusedStyles...)
	case g.focused && len(g.SelectedOptionStyles) > 0:
		markups = append(markups, g.SelectedOptionStyles...)
		markups = append(markups, DefaultSelectedOptionFocusedStyles...)
	case g.focused:
		markups = append(markups, DefaultSelectedOptionFocusedStyles...)
	case len(g.SelectedOptionStyles) > 0:
		markups = append(markups, g.SelectedOptionStyles...)
	}

	ph := &g.Placeholder
	if (ph == nil && g.SelectedPos > 0) ||
		g.SelectedPos > 0 && len(g.Options) >= g.SelectedPos {
		ph = g.Options[g.SelectedPos-1]
	}

	if !g.Disabled {

		markups = append(markups,
			//we always take care of clicks not the user
			event.Click(func(v *vecty.Event) {
				//as soon as we
				if g.SelectedPos == 0 && len(g.Options) >= 1 {
					g.SelectedPos++
					g.Options[0].Hightlight = true
				}

				// We use a hidden select to help with focus awareness
				// and user input. But this (click) action will take the focus away from the
				// hidden select and will cause a blur event on the hidden select.
				// Any time blur events occurs to the hidden select it sets blurRerender
				// true. Then after some small amount of time it will rerender
				// (effectively loosing focus for this component). To disable that
				//from occurring we just set blurRerender to false.
				g.blurRerender = false

				//when we click we always want to toggle showing the options
				g.expanded = !g.expanded
				g.focused = true
				go func() {
					internal.RequestFocusEvent(g.selectID)
				}()
				vecty.Rerender(g)
			}).StopPropagation(),

			event.MouseEnter(func(v *vecty.Event) {
				g.mouseOver = true
				vecty.Rerender(g)
				if g.MouseEnter != nil {
					g.MouseEnter(v)
				}
			}),
			event.MouseLeave(func(v *vecty.Event) {
				g.mouseOver = false
				vecty.Rerender(g)
				if g.MouseLeave != nil {
					g.MouseLeave(v)
				}
			}),
		)
	}

	optionButtonAlignment := DefaultSelectedOptionButtonAlignmentStyles
	if len(g.SelectedOptionButtonAlignmentStyles) > 0 {
		optionButtonAlignment = g.SelectedOptionButtonAlignmentStyles
	}

	var button vecty.ComponentOrHTML
	if g.Button != nil {
		button = g.Button(g.Disabled, g.mouseOver, g.focused, g.expanded)
	}

	return elem.Div(
		vecty.Markup(markups...),
		elem.Div(
			vecty.Markup(optionButtonAlignment...),
			elem.Div(
				vecty.Markup(position.Relative),
				ph.Option(g.Disabled, false, true),
			),

			elem.Div(
				vecty.Markup(position.Relative),
				button,
			),
		),
		g.makeOptionList(),
	)
}

func (g *Generic) makeHiddenSelect() vecty.ComponentOrHTML {
	hiddenOptions := make([]vecty.ComponentOrHTML, 0)
	for i, o := range g.Options {
		option := o
		index := i

		//for every option we make a hidden one
		hiddenOptions = append(hiddenOptions,
			elem.Option(
				vecty.Markup(
					vecty.Property("disabled", option.Disabled),
					vecty.Property("value", index),
				),
			),
		)
	}

	return elem.Select(
		vecty.Markup(
			position.Absolute,
			style.Opacity(0),
			style.Height(1),
			style.Top(0),
			prop.ID(g.selectID),
			vecty.Property("disabled", g.Disabled),
			vecty.Property("size", strconv.Itoa(len(g.Options))),
			vecty.MarkupIf(!g.Disabled,
				event.FocusIn(func(v *vecty.Event) {
					g.focused = true
					vecty.Rerender(g)
				}),
				event.Change(func(v *vecty.Event) {
					x, _ := strconv.Atoi(v.Target.Get("value").String())
					u := -1
					for i := 0; i < len(g.Options); i++ {
						if g.Options[i].Hightlight != false {
							u = i
							break
						}
					}

					if x != u && !g.Options[x].Disabled {
						if u >= 0 {
							g.Options[u].Hightlight = false
						}
						g.Options[x].Hightlight = true
						g.SelectedPos = x + 1
						vecty.Rerender(g)
					}
				}),
				event.Blur(func(v *vecty.Event) {
					//we're loosing focus time to collapse the
					g.blurRerender = true
					go func() {
						time.Sleep(100 * time.Millisecond)
						if g.blurRerender {
							g.expanded = false
							g.focused = false
							vecty.Rerender(g)
						}
					}()
				}),
				event.KeyDown(func(v *vecty.Event) {
					switch v.Get("code").String() {
					case "Enter":
						g.expanded = !g.expanded
						vecty.Rerender(g)
					case "Space":
						g.expanded = true
						vecty.Rerender(g)
					case "Tab":
						if g.expanded {
							//here we just want it to close the list
							// but stay on this item
							g.expanded = false
							g.focused = true
							go func() {
								time.Sleep(100 * time.Millisecond)
								internal.RequestFocusEvent(g.selectID)
							}()

						}
						vecty.Rerender(g)
					}
				}),
			),
		),
		vecty.List(hiddenOptions),
	)
}

func (g *Generic) makeOptionList() vecty.ComponentOrHTML {
	optionsHolder := make([]vecty.MarkupOrChild, 0)
	optionsHolderMarkups := make([]vecty.Applyer, 0)

	if g.expanded {
		optionsHolderMarkups = append(optionsHolderMarkups, visibility.Visible)
	} else {
		optionsHolderMarkups = append(optionsHolderMarkups, visibility.Hidden)
	}

	optionsHolderMarkups = append(optionsHolderMarkups,
		position.Absolute,
		style.Width("100%"),
	)

	//now we're at the styles
	if len(g.OptionsStyles) != 0 {
		optionsHolderMarkups = append(optionsHolderMarkups, g.OptionsStyles...)
	} else {
		optionsHolderMarkups = append(optionsHolderMarkups, DefaultOptionsStyles...)
	}
	optionsHolder = append(optionsHolder, vecty.Markup(optionsHolderMarkups...))

	for i, o := range g.Options {
		option := o
		index := i

		// and now to make a visible one
		optionMarkups := make([]vecty.Applyer, 0)

		//there's two things we take care here
		//1) all events
		//2) setting the position and width
		optionMarkups = append(optionMarkups,
			position.Relative,
			style.Width("100%"),
			event.MouseOver(func(v *vecty.Event) {
				if option.Disabled {
					return
				}
				//we need to clear previous highlight
				if index != g.SelectedPos-1 {
					if g.SelectedPos > 0 && len(g.Options) >= g.SelectedPos {
						g.Options[g.SelectedPos-1].Hightlight = false
					}

					option.Hightlight = true
					g.SelectedPos = index + 1
					vecty.Rerender(g)

					//lets update the hidden select too
					go func() {
						g.setSelectIndex(g.selectID, index)
					}()
				}
				if option.MouseMove != nil {
					option.MouseMove(v)
				}
			}),
			event.Click(func(v *vecty.Event) {
				if option.Disabled {
					return
				}

				option.Hightlight = false
				g.SelectedPos = index + 1
				if g.SelectedEvent != nil {
					g.SelectedEvent(index, g)
				}
				g.expanded = false
				vecty.Rerender(g)
				go func() {
					internal.RequestFocusEvent(g.selectID)
				}()

			}).StopPropagation(),
		)
		optionsHolder = append(optionsHolder, elem.Div(
			vecty.Markup(optionMarkups...),
			option,
		))
	}

	return elem.Div(optionsHolder...)
}

func (g *Generic) setSelectIndex(id string, index int) {
	for i := 0; i < 6; i++ {
		d := js.Global().Get("document").Call("getElementById", id)
		if js.Null() != d.JSValue() {
			d.Get("options").Index(index).Set("selected", true)
		}
		time.Sleep(10 * time.Duration(i) * time.Millisecond)
	}
}
