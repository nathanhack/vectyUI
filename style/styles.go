package style

import (
	"github.com/nathanhack/vectyUI/style/background"
	"github.com/nathanhack/vectyUI/style/border"
	"github.com/nathanhack/vectyUI/style/borderColor"
	"github.com/nathanhack/vectyUI/style/borderStyle"
	"github.com/nathanhack/vectyUI/style/borderWidth"
	"github.com/nathanhack/vectyUI/style/color"
	"github.com/nathanhack/vectyUI/style/fontFamily"
	"github.com/nathanhack/vectyUI/style/fontName"
	"reflect"
	"strconv"
	"strings"
)

func Background(color string) background.Type {
	return background.Type(color)
}

func Border(width int, style borderStyle.Type, color string) border.Border {
	return border.Border{
		Width: width,
		Style: style,
		Color: color,
	}
}

func BorderColor(colors ...string) borderColor.Color {
	return colors
}

func BorderWidth(widths ...interface{}) borderWidth.Width {
	borderWidth := make([]string, 0)
	for _, w := range widths {
		v := reflect.TypeOf(w)
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			borderWidth = append(borderWidth, strconv.FormatInt(reflect.ValueOf(w).Int(), 10))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			borderWidth = append(borderWidth, strconv.FormatUint(reflect.ValueOf(w).Uint(), 10))
		case reflect.String:
			borderWidth = append(borderWidth, reflect.ValueOf(w).String())
		}
	}
	return borderWidth
}

func Color(c string) color.Type {
	return color.Type(c)
}

func FontFamily(fonts ...fontName.Type) fontFamily.FontFamilyType {
	sb := strings.Builder{}
	for i, f := range fonts {
		sb.WriteString(string(f))
		if i < len(fonts)-1 {
			sb.WriteString(",")
		}
	}
	return fontFamily.FontFamilyType(sb.String())
}
