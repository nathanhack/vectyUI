package style

import (
	"fmt"
	"github.com/nathanhack/vectyUI/internal"
	"github.com/nathanhack/vectyUI/style/alignItems"
	"github.com/nathanhack/vectyUI/style/alignSelf"
	"github.com/nathanhack/vectyUI/style/background"
	"github.com/nathanhack/vectyUI/style/backgroundAttachment"
	"github.com/nathanhack/vectyUI/style/backgroundClip"
	"github.com/nathanhack/vectyUI/style/backgroundColor"
	"github.com/nathanhack/vectyUI/style/backgroundImage"
	"github.com/nathanhack/vectyUI/style/backgroundOrigin"
	"github.com/nathanhack/vectyUI/style/backgroundPosition"
	"github.com/nathanhack/vectyUI/style/backgroundRepeat"
	"github.com/nathanhack/vectyUI/style/backgroundSize"
	"github.com/nathanhack/vectyUI/style/border"
	"github.com/nathanhack/vectyUI/style/borderColor"
	"github.com/nathanhack/vectyUI/style/borderRadius"
	"github.com/nathanhack/vectyUI/style/borderStyle"
	"github.com/nathanhack/vectyUI/style/borderWidth"
	"github.com/nathanhack/vectyUI/style/bottom"
	"github.com/nathanhack/vectyUI/style/boxShadow"
	"github.com/nathanhack/vectyUI/style/color"
	"github.com/nathanhack/vectyUI/style/cursor"
	"github.com/nathanhack/vectyUI/style/display"
	"github.com/nathanhack/vectyUI/style/flexDirection"
	"github.com/nathanhack/vectyUI/style/flexWrap"
	"github.com/nathanhack/vectyUI/style/float"
	"github.com/nathanhack/vectyUI/style/fontFamily"
	"github.com/nathanhack/vectyUI/style/fontSize"
	"github.com/nathanhack/vectyUI/style/fontWeight"
	"github.com/nathanhack/vectyUI/style/height"
	"github.com/nathanhack/vectyUI/style/justifyContent"
	"github.com/nathanhack/vectyUI/style/left"
	"github.com/nathanhack/vectyUI/style/lineHeight"
	"github.com/nathanhack/vectyUI/style/margin"
	"github.com/nathanhack/vectyUI/style/marginBottom"
	"github.com/nathanhack/vectyUI/style/marginLeft"
	"github.com/nathanhack/vectyUI/style/marginRight"
	"github.com/nathanhack/vectyUI/style/marginTop"
	"github.com/nathanhack/vectyUI/style/opacity"
	"github.com/nathanhack/vectyUI/style/overflow"
	"github.com/nathanhack/vectyUI/style/padding"
	"github.com/nathanhack/vectyUI/style/position"
	"github.com/nathanhack/vectyUI/style/right"
	"github.com/nathanhack/vectyUI/style/textAlign"
	"github.com/nathanhack/vectyUI/style/top"
	"github.com/nathanhack/vectyUI/style/transition"
	"github.com/nathanhack/vectyUI/style/transitionDelay"
	"github.com/nathanhack/vectyUI/style/transitionDuration"
	"github.com/nathanhack/vectyUI/style/transitionProperty"
	"github.com/nathanhack/vectyUI/style/transitionTimingFunction"
	"github.com/nathanhack/vectyUI/style/userSelect"
	"github.com/nathanhack/vectyUI/style/verticalAlign"
	"github.com/nathanhack/vectyUI/style/visibility"
	"github.com/nathanhack/vectyUI/style/width"
	"github.com/nathanhack/vectyUI/style/zIndex"
	"strings"
)

func AlignItems(value alignItems.Type) alignItems.Value {
	return alignItems.Value(value)
}

func AlignSelf(value alignSelf.Type) alignSelf.Value {
	return alignSelf.Value(value)
}

//Background is used for weird case not handled by Background.
// Expects either strings or something that casts to a string. Then basically concatenates the strings.
// NOTE items are not ordered. Also the size value must be prefixed by "/"
func Background(nonRepeatingBackgroundValues ...interface{}) background.Value {
	sb := strings.Builder{}
	for _, bv := range nonRepeatingBackgroundValues {
		sb.WriteString(internal.Stringify(bv, ""))
		sb.WriteString(" ")

	}
	return background.Value(sb.String())
}

//Backgrounds
// NOTE only the last layer can include background-color
func Backgrounds(backgrounds ...background.Value) background.Value {
	sb := strings.Builder{}
	for i, bv := range backgrounds {
		sb.WriteString(internal.Stringify(bv, ""))
		if i < len(backgrounds)-1 {
			sb.WriteString(",")
		}
	}
	return background.Value(sb.String())
}

func BackgroundAttachment(attaches backgroundAttachment.Type) backgroundAttachment.Value {
	return backgroundAttachment.Value(attaches)
}

func BackgroundColor(color color.Type) backgroundColor.Value {
	return backgroundColor.Value(color)
}

func BackgroundClip(clipping backgroundClip.Type) backgroundClip.Value {
	return backgroundClip.Value(clipping)
}

func BackgroundImage(backgrounds ...backgroundImage.Type) backgroundImage.Value {
	return backgrounds
}

func BackgroundOrigin(origin backgroundOrigin.Type) backgroundOrigin.Value {
	return backgroundOrigin.Value(origin)
}

func BackgroundRepeat(repeats ...backgroundRepeat.Type) backgroundRepeat.Value {
	return repeats
}

func BackgroundPosition(positions ...backgroundPosition.Type) backgroundPosition.Value {
	return positions
}

func BackgroundSize(sizes ...backgroundSize.Type) backgroundSize.Value {
	return sizes
}

func Border(width interface{}, style borderStyle.Type, color color.Type) border.Value {
	return border.Value{
		Width: internal.Stringify(width, "px"),
		Style: style,
		Color: color,
	}
}

func BorderColor(colors ...interface{}) borderColor.Value {
	values := make([]borderColor.Type, 0)
	for _, c := range colors {
		switch c.(type) {
		case borderColor.Type:
			values = append(values, c.(borderColor.Type))
		case color.Type:
			values = append(values, borderColor.Type(c.(color.Type)))
		case string:
			values = append(values, borderColor.Type(c.(string)))
		default:
			panic(fmt.Sprintf("unsupported color type %T with value: %v", c, c))
		}
	}
	return values
}

//BorderRadius
// For the case of ellipses just pass in "/" ex:
// style.BorderRadius("10px","/","40px")
// style.BorderRadius("10% / 40em")
func BorderRadius(values ...borderRadius.Type) borderRadius.Value {
	return values
}

func BorderStyle(styles ...borderStyle.Type) borderStyle.Value {
	return styles
}

func BorderWidth(widths ...interface{}) borderWidth.Value {
	return internal.StringifyList("px", widths...)
}

//Bottom defaults to pixels if a number is passed in, otherwise it accepts as is
func Bottom(value interface{}) bottom.Value {
	return bottom.Value(internal.Stringify(value, "px"))
}

func BoxShadow(offsetsBlurSpreadColorAndInset ...interface{}) boxShadow.Value {
	return boxShadow.Value([]string{strings.Join(internal.StringifyList("px", offsetsBlurSpreadColorAndInset...), " ")})
}

func BoxShadows(boxes ...boxShadow.Value) boxShadow.Value {
	s := make([]string, 0)
	for _, b := range boxes {
		s = append(s, b...)
	}
	return s
}

func Color(c color.Type) color.Value {
	return color.Value(c)
}

//Cursor sets mouse cursor
// for URL() support pass in a string ex: "URL(..), URL(..), auto"
func Cursor(value cursor.Type) cursor.Value {
	return cursor.Value(value)
}

func Display(value display.Type) display.Value {
	return display.Value(value)
}

func FlexDirection(dir flexDirection.Type) flexDirection.Value {
	return flexDirection.Value(dir)
}

func FlexWrap(wrap flexWrap.Type) flexWrap.Value {
	return flexWrap.Value(wrap)
}

func Float(fl float.Type) float.Value {
	return float.Value(fl)
}

func FontFamily(fontNames ...fontFamily.Type) fontFamily.Value {
	return fontNames
}

//FontSize defaults to pixels if a numbers are passed in, otherwise it accepts as is
func FontSize(size interface{}) fontSize.Value {
	return fontSize.Value(internal.Stringify(size, "px"))
}

func FontWeight(value fontWeight.Type) fontWeight.Value {
	return fontWeight.Value(value)
}

//Height defaults to pixels if a numbers are passed in, otherwise it accepts as is
func Height(length interface{}) height.Value {
	return height.Value(internal.Stringify(length, "px"))
}

func JustifyContent(value justifyContent.Type) justifyContent.Value {
	return justifyContent.Value(value)
}

//Left defaults to pixels if a number is passed in, otherwise it accepts as is
func Left(value interface{}) left.Value {
	return left.Value(internal.Stringify(value, "px"))
}

//LineHeight defaults to pixels if a numbers are passed in, otherwise it accepts as is
func LineHeight(value interface{}) lineHeight.Value {
	return lineHeight.Value(internal.Stringify(value, "px"))
}

func Margin(lengths ...interface{}) margin.Value {
	return internal.StringifyList("px", lengths...)
}

func MarginBottom(length interface{}) marginBottom.Value {
	return marginBottom.Value(internal.Stringify(length, "px"))
}

func MarginLeft(length interface{}) marginLeft.Value {
	return marginLeft.Value(internal.Stringify(length, "px"))
}

//MarginRight defaults to pixels if a numbers are passed in, otherwise it accepts as is
func MarginRight(length interface{}) marginRight.Value {
	return marginRight.Value(internal.Stringify(length, "px"))
}

//MarginTop defaults to pixels if a numbers are passed in, otherwise it accepts as is
func MarginTop(length interface{}) marginTop.Value {
	return marginTop.Value(internal.Stringify(length, "px"))
}

func Opacity(value interface{}) opacity.Value {
	return opacity.Value(opacity.Number(value))
}

func Overflow(p overflow.Type) overflow.Value {
	return overflow.Value(p)
}

//Padding defaults to pixels if a numbers are passed in, otherwise it accepts as is
func Padding(lengths ...interface{}) padding.Value {
	return internal.StringifyList("px", lengths...)
}

func Position(pos position.Type) position.Value {
	return position.Value(pos)
}

//Right defaults to pixels if a numbers are passed in, otherwise it accepts as is
func Right(length interface{}) right.Value {
	return right.Value(internal.Stringify(length, "px"))
}

func TextAlign(value textAlign.Type) textAlign.Value {
	return textAlign.Value(value)
}

func Top(length interface{}) top.Value {
	return top.Value(internal.Stringify(length, "px"))
}

//Background is used for weird case not handled by Background.
// Expects either strings or something that casts to a string. Then basically concatenates the strings.
func Transition(nonRepeatingTransitionValues ...interface{}) transition.Value {
	sb := strings.Builder{}
	for _, bv := range nonRepeatingTransitionValues {
		sb.WriteString(internal.Stringify(bv, ""))
		sb.WriteString(" ")

	}
	return transition.Value(sb.String())
}

//Transitions is used when specifying multiple transitions
func Transitions(transitions ...transition.Value) transition.Value {
	sb := strings.Builder{}
	for i, bv := range transitions {
		sb.WriteString(string(bv))
		if i < len(transitions)-1 {
			sb.WriteString(",")
		}
	}
	return transition.Value(sb.String())
}

func TransitionDelay(value transitionDelay.Type) transitionDelay.Value {
	return transitionDelay.Value(value)
}

func TransitionDelays(delays ...transitionDelay.Type) transitionDelay.Value {
	sb := strings.Builder{}
	for i, bv := range delays {
		sb.WriteString(string(bv))
		if i < len(delays)-1 {
			sb.WriteString(",")
		}
	}
	return transitionDelay.Value(sb.String())
}

func TransitionDuration(value transitionDuration.Type) transitionDuration.Value {
	return transitionDuration.Value(value)
}

func TransitionDurations(delays ...transitionDuration.Type) transitionDuration.Value {
	sb := strings.Builder{}
	for i, bv := range delays {
		sb.WriteString(string(bv))
		if i < len(delays)-1 {
			sb.WriteString(",")
		}
	}
	return transitionDuration.Value(sb.String())
}

func TransitionProperty(value transitionProperty.Type) transitionProperty.Value {
	return transitionProperty.Value(value)
}

func TransitionProperties(delays ...transitionProperty.Type) transitionProperty.Value {
	sb := strings.Builder{}
	for i, bv := range delays {
		sb.WriteString(string(bv))
		if i < len(delays)-1 {
			sb.WriteString(",")
		}
	}
	return transitionProperty.Value(sb.String())
}

func TransitionTimingFunction(value transitionTimingFunction.Type) transitionTimingFunction.Value {
	return transitionTimingFunction.Value(value)
}

func TransitionTimingFunctions(delays ...transitionTimingFunction.Type) transitionTimingFunction.Value {
	sb := strings.Builder{}
	for i, bv := range delays {
		sb.WriteString(string(bv))
		if i < len(delays)-1 {
			sb.WriteString(",")
		}
	}
	return transitionTimingFunction.Value(sb.String())
}

func UserSelect(value userSelect.Type) userSelect.Value {
	return userSelect.Value(value)
}

func VerticalAlign(value verticalAlign.Type) verticalAlign.Value {
	return verticalAlign.Value(value)
}

func Visibility(value visibility.Type) visibility.Value {
	return visibility.Value(value)
}

func Width(length interface{}) width.Value {
	return width.Value(internal.Stringify(length, "px"))
}

func ZIndex(typeOrNumber interface{}) zIndex.Value {
	return zIndex.Value(internal.Stringify(typeOrNumber, ""))
}
