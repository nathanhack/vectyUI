package sel

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/nathanhack/vectyUI/style"
	"github.com/nathanhack/vectyUI/style/alignItems"
	"github.com/nathanhack/vectyUI/style/display"
	"github.com/nathanhack/vectyUI/style/flexWrap"
	"github.com/nathanhack/vectyUI/style/justifyContent"
	"github.com/nathanhack/vectyUI/style/outlineStyle"
	"github.com/nathanhack/vectyUI/style/position"
	"github.com/nathanhack/vectyUI/style/visibility"
	"github.com/nathanhack/vectyUI/style/zIndex"
)

//GenericOption is effectively a div container for the option
// Notes:
// 1) User should not directly create the following events: MouseMove and Click.
// 2) Each call to GenericOption.Option() should return a distinct vecty.ComponentOrHTML.
type GenericOption struct {
	vecty.Core
	Option    func(disabled bool, highlight bool, selected bool) vecty.ComponentOrHTML `vecty:"prop"`
	MouseMove func(v *vecty.Event)                                                     `vecty:"prop"`
	Disabled  bool                                                                     `vecty:"prop"`
	Highlight bool                                                                     `vecty:"prop"`
}

func (g *GenericOption) Render() vecty.ComponentOrHTML {
	return g.Option(g.Disabled, g.Highlight, false)
}

const genericSelectID = "GenericSelectID"

var genericSelectCount = 0

var DefaultStyles = []vecty.Applyer{
	vecty.Attribute("tabindex", 0),
	style.OutlineStyle(outlineStyle.None),
}

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
// 3) The following should be avoided in styles or components: position, display and justifyContent.
// 4) To create without a button, set the Button to and empty Div.
// 5) Normally disabled options are unselectable, at startup this can be overridden by setting .SelectPos (use carefully).
type Generic struct {
	vecty.Core
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
	hightlightPos                       int                                                                     // pos of highlighting
	expanded                            bool                                                                    // will render option list visible
	focused                             bool                                                                    // will render component as focused
	mouseOver                           bool                                                                    // used to identify when mouse is over placeholder/selected or button
}

func (g *Generic) Render() vecty.ComponentOrHTML {
	markups := make([]vecty.Applyer, 0)
	markups = append(markups,
		event.FocusIn(func(v *vecty.Event) {
			g.focused = true
			if g.SelectedPos == 0 {
				for i := 0; i < len(g.Options); i++ {
					if !g.Options[i].Disabled {
						g.SelectedPos = i + 1
						if g.expanded {
							g.hightlightPos = g.SelectedPos
						}
						vecty.Rerender(g)
						g.SelectedEvent(i, g)
						break
					}
				}
			}
			vecty.Rerender(g)
		}),
		event.FocusOut(func(v *vecty.Event) {
			g.focused = false
			g.expanded = false

			vecty.Rerender(g)
		}),
		event.KeyDown(func(v *vecty.Event) {
			switch v.Get("code").String() {
			case "Escape":
				g.expanded = false
				vecty.Rerender(g)
				v.Value.Call("stopPropagation")
				v.Value.Call("preventDefault")
			case "Enter":
				if !g.expanded {
					g.hightlightPos = g.SelectedPos
				}
				g.expanded = !g.expanded
				vecty.Rerender(g)
				v.Value.Call("stopPropagation")
				v.Value.Call("preventDefault")
			case "Space":
				if !g.expanded {
					g.hightlightPos = g.SelectedPos
				}
				g.expanded = true
				vecty.Rerender(g)
				v.Value.Call("stopPropagation")
				v.Value.Call("preventDefault")
			case "Tab":
				if g.expanded {
					//here we just want it to close the list
					// but stay on this item
					g.expanded = false
					vecty.Rerender(g)
					v.Value.Call("stopPropagation")
					v.Value.Call("preventDefault")
				}
			case "ArrowDown":
				for i := 0; i < len(g.Options); i++ {
					if i+1 > g.SelectedPos && !g.Options[i].Disabled {
						g.SelectedPos = i + 1
						if g.expanded {
							g.hightlightPos = g.SelectedPos
						}
						vecty.Rerender(g)
						g.SelectedEvent(i, g)
						break
					}
				}

				v.Value.Call("stopPropagation")
				v.Value.Call("preventDefault")
			case "ArrowUp":
				for i := len(g.Options) - 1; i >= 0; i-- {
					if i+1 < g.SelectedPos && !g.Options[i].Disabled {
						g.SelectedPos = i + 1
						if g.expanded {
							g.hightlightPos = g.SelectedPos
						}
						vecty.Rerender(g)
						g.SelectedEvent(i, g)
						break
					}
				}

				v.Value.Call("stopPropagation")
				v.Value.Call("preventDefault")
			case "Home":
				for i := 0; i < len(g.Options); i++ {
					if !g.Options[i].Disabled {
						g.SelectedPos = i + 1
						if g.expanded {
							g.hightlightPos = g.SelectedPos
						}
						vecty.Rerender(g)
						g.SelectedEvent(i, g)
						break
					}
				}

				v.Value.Call("stopPropagation")
				v.Value.Call("preventDefault")
			case "End":
				for i := len(g.Options) - 1; i >= 0; i-- {
					if !g.Options[i].Disabled {
						g.SelectedPos = i + 1
						if g.expanded {
							g.hightlightPos = g.SelectedPos
						}
						vecty.Rerender(g)
						g.SelectedEvent(i, g)
						break
					}
				}

				v.Value.Call("stopPropagation")
				v.Value.Call("preventDefault")

			}

		}),
	)
	if len(g.Styles) == 0 {
		markups = append(markups, DefaultStyles...)
	} else {
		markups = append(markups, g.Styles...)
	}

	return elem.Div(
		vecty.Markup(markups...),
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

	var ph vecty.ComponentOrHTML
	switch {
	case g.SelectedPos > 0 && len(g.Options) >= g.SelectedPos:
		ph = g.Options[g.SelectedPos-1].Option(g.Disabled, false, true)
	case g.Placeholder.Option != nil:
		ph = g.Placeholder.Option(g.Disabled, false, true)
	default:
		ph = elem.Div()
	}

	if !g.Disabled {
		markups = append(markups,
			//we always take care of clicks not the user
			event.Click(func(v *vecty.Event) {
				//as soon as we
				if g.hightlightPos == 0 && len(g.Options) >= 1 {
					g.hightlightPos++
				}

				//when we click we always want to toggle showing the options
				g.expanded = !g.expanded
				g.focused = true
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
				ph,
			),

			elem.Div(
				vecty.Markup(position.Relative),
				button,
			),
		),
		g.makeOptionList(),
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

		if g.hightlightPos == i+1 {
			option.Highlight = true
		} else {
			option.Highlight = false
		}
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
				if index != g.hightlightPos-1 {
					g.hightlightPos = index + 1
					vecty.Rerender(g)
				}
				if option.MouseMove != nil {
					option.MouseMove(v)
				}
			}),
			event.Click(func(v *vecty.Event) {
				if option.Disabled {
					return
				}

				g.SelectedPos = index + 1
				if g.SelectedEvent != nil {
					g.SelectedEvent(index, g)
				}
				g.expanded = false
				vecty.Rerender(g)

			}).StopPropagation(),
		)
		optionsHolder = append(optionsHolder, elem.Div(
			vecty.Markup(optionMarkups...),
			option,
		))
	}

	return elem.Div(optionsHolder...)
}
