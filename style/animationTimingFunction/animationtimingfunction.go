package animationTimingFunction

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
	"github.com/nathanhack/vectyUI/style/animationTimingFunction/jumpTerm"
	"strconv"
	"strings"
)

type Type string

const (
	Linear    Type = "linear"
	Ease      Type = "ease"
	EaseIn    Type = "ease-in"
	EaseOut   Type = "ease-out"
	EaseInOut Type = "ease-in-out"
	StepStart Type = "step-start"
	StepEnd   Type = "step-end"
	Initial   Type = "initial"
	Inherit   Type = "inherit"
	Unset     Type = "unset"
)

var styleNames = []string{"-webkit-transition-timing-function", "transition-timing-function"}

func (t Type) Apply(h *vecty.HTML) {
	for _, s := range styleNames {
		vecty.Style(s, string(t)).Apply(h)
	}
}

func (t Type) AddTo(m map[string]string) {
	for _, s := range styleNames {
		m[s] = string(t)
	}
}

func Steps(steps int, dir jumpTerm.Type) Type {
	return Type("steps(" + strconv.Itoa(steps) + "," + string(dir) + ")")
}

//CubicBezier defines a cubic-bezier curve
func CubicBezier(n0, n1, n2, n3 interface{}) Type {
	sb := strings.Builder{}
	sb.WriteString("cubic-bezier(")
	sb.WriteString(internal.Stringify(n0, ""))
	sb.WriteString(internal.Stringify(n1, ""))
	sb.WriteString(internal.Stringify(n2, ""))
	sb.WriteString(internal.Stringify(n3, ""))
	sb.WriteString(")")

	return Type(sb.String())
}

type Value string

func (v Value) Apply(h *vecty.HTML) {
	for _, s := range styleNames {
		vecty.Style(s, string(v)).Apply(h)
	}
}
