package fitContent

import (
	"github.com/nathanhack/vectyUI/internal"
)

type Func string

func New(lenOrPercent interface{}) Func {
	return Func("fit-content(" + internal.Stringify(lenOrPercent, "px") + ")")
}
