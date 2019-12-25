package style

import (
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
	"github.com/nathanhack/vectyUI/style/color"
	"github.com/nathanhack/vectyUI/style/display"
	"github.com/nathanhack/vectyUI/style/flexDirection"
	"github.com/nathanhack/vectyUI/style/flexWrap"
	"github.com/nathanhack/vectyUI/style/float"
	"github.com/nathanhack/vectyUI/style/fontFamily"
	"github.com/nathanhack/vectyUI/style/fontSize"
	"github.com/nathanhack/vectyUI/style/height"
	"github.com/nathanhack/vectyUI/style/left"
	"github.com/nathanhack/vectyUI/style/margin"
	"github.com/nathanhack/vectyUI/style/marginBottom"
	"github.com/nathanhack/vectyUI/style/marginLeft"
	"github.com/nathanhack/vectyUI/style/marginRight"
	"github.com/nathanhack/vectyUI/style/marginTop"
	"github.com/nathanhack/vectyUI/style/overflow"
	"github.com/nathanhack/vectyUI/style/padding"
	"github.com/nathanhack/vectyUI/style/position"
	"github.com/nathanhack/vectyUI/style/right"
	"github.com/nathanhack/vectyUI/style/top"
	"github.com/nathanhack/vectyUI/style/userSelect"
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

func BackgroundAttachment(attaches backgroundAttachment.Type) backgroundAttachment.Value {
	return backgroundAttachment.Value(attaches)
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

func BorderColor(colors ...color.Type) borderColor.Value {
	return colors
}

func BorderStyle(styles ...borderStyle.Type) borderStyle.Value {
	return borderStyle.Value(styles)
}

func BorderWidth(widths ...interface{}) borderWidth.Value {
	return internal.StringifyList("px", widths...)
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

func FontFamily(fontNames ...string) fontFamily.Value {
	return fontNames
}

func FontSize(size interface{}) fontSize.Value {
	return fontSize.Value(internal.Stringify(size, "px"))
}

func Height(length interface{}) height.Value {
	return height.Value(internal.Stringify(length, "px"))
}

func Left(length interface{}) left.Value {
	return left.Value(internal.Stringify(length, "px"))
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

func MarginRight(length interface{}) marginRight.Value {
	return marginRight.Value(internal.Stringify(length, "px"))
}

func MarginTop(length interface{}) marginTop.Value {
	return marginTop.Value(internal.Stringify(length, "px"))
}

func Overflow(p overflow.Type) overflow.Value {
	return overflow.Value(p)
}

func Padding(lengths ...interface{}) padding.Value {
	return internal.StringifyList("px", lengths...)
}

func Position(pos position.Type) position.Value {
	return position.Value(pos)
}

func Right(length interface{}) right.Value {
	return right.Value(internal.Stringify(length, "px"))
}

func Top(length interface{}) top.Value {
	return top.Value(internal.Stringify(length, "px"))
}

func UserSelect(value userSelect.Type) userSelect.Value {
	return userSelect.Value(value)
}

func Width(length interface{}) width.Value {
	return width.Value(internal.Stringify(length, "px"))
}

func ZIndex(typeOrNumber interface{}) zIndex.Value {
	return zIndex.Value(internal.Stringify(typeOrNumber, ""))
}
