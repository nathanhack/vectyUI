package style

import (
	"github.com/nathanhack/vectyUI/style/background"
	"github.com/nathanhack/vectyUI/style/border"
	"github.com/nathanhack/vectyUI/style/borderColor"
	"github.com/nathanhack/vectyUI/style/borderStyle"
	"github.com/nathanhack/vectyUI/style/borderWidth"
	"github.com/nathanhack/vectyUI/style/color"
	"github.com/nathanhack/vectyUI/style/fontFamily"
	"github.com/nathanhack/vectyUI/style/margin"
	"github.com/nathanhack/vectyUI/style/padding"
	"reflect"
	"strconv"
)

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

func BorderWidth(widths ...interface{}) borderWidth.Value {
	println("BorderWidth")
	return stringify("px", widths)
}

func Margin(lengths ...interface{}) margin.Value {
	println("Margin")
	return stringify("px", lengths)
}

func Padding(lengths ...interface{}) padding.Value {
	println("Padding")
	return stringify("px", lengths)
}

func Color(c string) color.Value {
	return color.Value(c)
}

func FontFamily(fontNames ...string) fontFamily.Value {
	return fontNames
}

func stringify(postfix string, values ...interface{}) []string {
	stringValues := make([]string, 0)
	for _, l := range values {
		v := reflect.TypeOf(l)
		println(v.Kind())
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			stringValues = append(stringValues, strconv.FormatInt(reflect.ValueOf(l).Int(), 10)+postfix)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			stringValues = append(stringValues, strconv.FormatUint(reflect.ValueOf(l).Uint(), 10)+postfix)
		case reflect.Float32, reflect.Float64:
			stringValues = append(stringValues, strconv.FormatFloat(reflect.ValueOf(l).Float(), 'f', -1, 64)+postfix)
		case reflect.String:
			stringValues = append(stringValues, reflect.ValueOf(l).String())
		}
	}
	println(stringValues)
	return stringValues
}
