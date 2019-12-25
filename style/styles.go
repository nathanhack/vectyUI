package style

import (
	"fmt"
	"github.com/nathanhack/vectyUI/internal"
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
	"github.com/nathanhack/vectyUI/style/borderStyle"
	"github.com/nathanhack/vectyUI/style/borderWidth"
	"github.com/nathanhack/vectyUI/style/bottom"
	"github.com/nathanhack/vectyUI/style/color"
	"github.com/nathanhack/vectyUI/style/display"
	"github.com/nathanhack/vectyUI/style/flexDirection"
	"github.com/nathanhack/vectyUI/style/flexWrap"
	"github.com/nathanhack/vectyUI/style/float"
	"github.com/nathanhack/vectyUI/style/fontFamily"
	"github.com/nathanhack/vectyUI/style/fontSize"
	"github.com/nathanhack/vectyUI/style/height"
	"github.com/nathanhack/vectyUI/style/left"
	"github.com/nathanhack/vectyUI/style/lineHeight"
	"github.com/nathanhack/vectyUI/style/margin"
	"github.com/nathanhack/vectyUI/style/marginBottom"
	"github.com/nathanhack/vectyUI/style/marginLeft"
	"github.com/nathanhack/vectyUI/style/marginRight"
	"github.com/nathanhack/vectyUI/style/marginTop"
	"github.com/nathanhack/vectyUI/style/overflow"
	"github.com/nathanhack/vectyUI/style/padding"
	"github.com/nathanhack/vectyUI/style/position"
	"github.com/nathanhack/vectyUI/style/right"
	"github.com/nathanhack/vectyUI/style/textAlign"
	"github.com/nathanhack/vectyUI/style/top"
	"github.com/nathanhack/vectyUI/style/userSelect"
	"github.com/nathanhack/vectyUI/style/verticalAlign"
	"github.com/nathanhack/vectyUI/style/width"
	"github.com/nathanhack/vectyUI/style/zIndex"
	"strings"
)

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

func Color(c color.Type) color.Value {
	return color.Value(c)
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

//Height defaults to pixels if a numbers are passed in, otherwise it accepts as is
func Height(length interface{}) height.Value {
	return height.Value(internal.Stringify(length, "px"))
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

func UserSelect(value userSelect.Type) userSelect.Value {
	return userSelect.Value(value)
}

func VerticalAlign(value verticalAlign.Type) verticalAlign.Value {
	return verticalAlign.Value(value)
}

func Width(length interface{}) width.Value {
	return width.Value(internal.Stringify(length, "px"))
}

func ZIndex(typeOrNumber interface{}) zIndex.Value {
	return zIndex.Value(internal.Stringify(typeOrNumber, ""))
}
