package repeat

import (
	"github.com/nathanhack/vectyUI/internal"
)

type CountType string

const (
	AutoFill CountType = "auto-fill"
	AutoFit  CountType = "auto-fit"
)

type Func string

func New(count, fragment interface{}) Func {
	return Func("repeat(" + internal.Stringify(count, "") + "," + internal.Stringify(fragment, "px") + ")")
}
