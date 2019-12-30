package backgroundImage

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
	"github.com/nathanhack/vectyUI/style/backgroundImage/colorStop"
	"github.com/nathanhack/vectyUI/style/position"
	"strings"
)

type Type string

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("background-image", string(t)).Apply(h)
}

type Value []Type

func (v Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for i, l := range v {
		sb.WriteString(string(l))
		if i < len(v)-1 {
			sb.WriteString(",")
		}
	}
	vecty.Style("background-image", sb.String()).Apply(h)
}

//URL make
func URL(url string) Type {
	return Type("url(" + url + ")")
}

type Dir string

var ToRight Dir = "to right"
var ToLeft Dir = "to left"
var ToBottom Dir = "to bottom"
var ToTop Dir = "to top"

func Degree(angle interface{}) Dir {
	return Dir(internal.Stringify(angle, "deg"))
}

//LinearGradient
// Note colors can be color2.Type too.
func LinearGradient(dir Dir, colors ...colorStop.Type) Type {
	return linear("linear-gradient", dir, colors...)
}

func RepeatingLinearGradient(dir Dir, colors ...colorStop.Type) Type {
	return linear("repeating-linear-gradient", dir, colors...)
}

func linear(linearType string, dir Dir, colors ...colorStop.Type) Type {
	sb := strings.Builder{}
	sb.WriteString(linearType + "(")
	sb.WriteString(string(dir))

	for _, c := range colors {
		sb.WriteString(",")
		sb.WriteString(string(c))
	}
	sb.WriteString(")")
	return Type(sb.String())
}

type SizeOrExtent string

var ClosestCorner SizeOrExtent = "closest-corner"
var ClosestSide SizeOrExtent = "closest-side"
var FarthestCorner SizeOrExtent = "farthest-corner"
var FarthestSide SizeOrExtent = "farthest-side"

func Pixels(len interface{}) SizeOrExtent {
	return SizeOrExtent(internal.Stringify(len, "px"))
}

func Percent(percent interface{}) SizeOrExtent {
	return SizeOrExtent(internal.Stringify(percent, "%"))
}

//RadialGradientCircle
// sizeOrExtent is optional, only Pixels() or Extent const are allowed (Percent is NOT allowed)
// position is optional
func RadialGradientCircle(sizeOrExtent *SizeOrExtent, position *position.Value, colorStops ...colorStop.Type) Type {
	return radial("radial-gradient", "circle", sizeOrExtent, position, colorStops...)
}

//RadialGradientEllipse
// sizeOrExtent is optional, requires two Pixels()'s or Percent()'s (ex. 20% 20%), only one extent const
// position is optional
func RadialGradientEllipse(sizeOrExtent *SizeOrExtent, position *position.Value, colorStops ...colorStop.Type) Type {
	return radial("radial-gradient", "ellipse", sizeOrExtent, position, colorStops...)
}

//RepeatingRadialGradientCircle
// sizeOrExtent is optional, only Pixels() or Extent const are allowed (Percent is NOT allowed)
// position is optional
func RepeatingRadialGradientCircle(sizeOrExtent *SizeOrExtent, position *position.Value, colorStops ...colorStop.Type) Type {
	return radial("repeating-radial-gradient", "circle", sizeOrExtent, position, colorStops...)
}

//RepeatingRadialGradientEllipse
// sizeOrExtent is optional, requires two Pixels()'s or Percent()'s (ex. 20% 20%), only one extent const
// position is optional
func RepeatingRadialGradientEllipse(sizeOrExtent *SizeOrExtent, position *position.Value, colorStops ...colorStop.Type) Type {
	return radial("repeating-radial-gradient", "ellipse", sizeOrExtent, position, colorStops...)
}

func radial(conicType string, radialType string, sizeOrExtent *SizeOrExtent, position *position.Value, colorStops ...colorStop.Type) Type {
	sb := strings.Builder{}
	sb.WriteString(conicType + "(" + radialType)

	if sizeOrExtent != nil {
		sb.WriteString("," + string(*sizeOrExtent))
	}
	if position != nil {
		sb.WriteString(", at")
		for _, p := range *position {
			sb.WriteString(" " + string(p))
		}
	}

	for _, c := range colorStops {
		sb.WriteString("," + string(c))
	}
	sb.WriteString(")")
	return Type(sb.String())
}
