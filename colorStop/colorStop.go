package colorStop

import (
	color2 "github.com/nathanhack/vectyUI/color"
	"github.com/nathanhack/vectyUI/internal"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Type string //

func From(c color2.Type) Type {
	return Type(c)
}

//Percent returns a colorStop.Type and
func Percent(c color2.Type, percents ...interface{}) Type {
	sb := strings.Builder{}
	sb.WriteString(string(c))
	sb.WriteString(" ")
	for _, p := range percents[:min(2, len(percents))] {
		sb.WriteString(internal.Stringify(p) + "%")
	}

	return Type(sb.String())
}

func Length(c color2.Type, lengths ...interface{}) Type {
	sb := strings.Builder{}
	sb.WriteString(string(c))
	sb.WriteString(" ")
	for _, l := range lengths[:min(2, len(lengths))] {
		sb.WriteString(internal.Stringify(l) + "px")
	}

	return Type(sb.String())
}

func HintPercent(percent interface{}) Type {
	return Type(internal.Stringify(percent) + "%")
}

func HintLength(len interface{}) Type {
	return Type(internal.Stringify(len) + "px")
}
