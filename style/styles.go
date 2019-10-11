package style

import (
	"github.com/nathanhack/vectyUI/style/background"
	"github.com/nathanhack/vectyUI/style/border"
	"github.com/nathanhack/vectyUI/style/borderColor"
	"github.com/nathanhack/vectyUI/style/borderStyle"
	"github.com/nathanhack/vectyUI/style/borderWidth"
	"github.com/nathanhack/vectyUI/style/color"
	"github.com/nathanhack/vectyUI/style/fontFamily"
	"github.com/nathanhack/vectyUI/style/fontSize"
	"github.com/nathanhack/vectyUI/style/height"
	"github.com/nathanhack/vectyUI/style/left"
	"github.com/nathanhack/vectyUI/style/margin"
	"github.com/nathanhack/vectyUI/style/marginBottom"
	"github.com/nathanhack/vectyUI/style/marginLeft"
	"github.com/nathanhack/vectyUI/style/marginRight"
	"github.com/nathanhack/vectyUI/style/marginTop"
	"github.com/nathanhack/vectyUI/style/padding"
	"github.com/nathanhack/vectyUI/style/right"
	"github.com/nathanhack/vectyUI/style/top"
	"github.com/nathanhack/vectyUI/style/width"
	"github.com/nathanhack/vectyUI/style/zIndex"
	"reflect"
	"strconv"
)

func stringify(postfix string, values ...interface{}) []string {
	stringValues := make([]string, 0)
	for _, l := range values {
		stringValues = append(stringValues, stringifyInterface(postfix, l))
	}
	return stringValues
}

func stringifyInterface(postfix string, value interface{}) string {
	v := reflect.TypeOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(reflect.ValueOf(value).Int(), 10) + postfix
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(reflect.ValueOf(value).Uint(), 10) + postfix
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(reflect.ValueOf(value).Float(), 'f', -1, 64) + postfix
	case reflect.String:
		return reflect.ValueOf(value).String()
	default:
		return ""
	}
}

func Background(color string) background.Value {
	return background.Value(color)
}

func Border(width int, style borderStyle.Type, color string) border.Value {
	return border.Value{
		Width: width,
		Style: style,
		Color: color,
	}
}

func BorderColor(colors ...string) borderColor.Value {
	return colors
}

func BorderStyle(styles ...borderStyle.Type) borderStyle.Value {
	return borderStyle.Value(styles)
}

func BorderWidth(widths ...interface{}) borderWidth.Value {
	return stringify("px", widths...)
}

func Color(c string) color.Value {
	return color.Value(c)
}

func FontFamily(fontNames ...string) fontFamily.Value {
	return fontNames
}

func FontSize(size interface{}) fontSize.Value {
	return fontSize.Value(stringifyInterface("px", size))
}

func Height(length interface{}) height.Value {
	return height.Value(stringifyInterface("px", length))
}

func Left(length interface{}) left.Value {
	return left.Value(stringifyInterface("px", length))
}

func Margin(lengths ...interface{}) margin.Value {
	return stringify("px", lengths...)
}

func MarginBottom(length interface{}) marginBottom.Value {
	return marginBottom.Value(stringifyInterface("px", length))
}

func MarginLeft(length interface{}) marginLeft.Value {
	return marginLeft.Value(stringifyInterface("px", length))
}

func MarginRight(length interface{}) marginRight.Value {
	return marginRight.Value(stringifyInterface("px", length))
}

func MarginTop(length interface{}) marginTop.Value {
	return marginTop.Value(stringifyInterface("px", length))
}

func Padding(lengths ...interface{}) padding.Value {
	return stringify("px", lengths...)
}

func Right(length interface{}) right.Value {
	return right.Value(stringifyInterface("px", length))
}

func Top(length interface{}) top.Value {
	return top.Value(stringifyInterface("px", length))
}

func Width(length interface{}) width.Value {
	return width.Value(stringifyInterface("px", length))
}

func ZIndex(typeOrNumber interface{}) zIndex.Value {
	return zIndex.Value(stringifyInterface("", typeOrNumber))
}
