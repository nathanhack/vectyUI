package minmax

import (
	"github.com/nathanhack/vectyUI/internal"
)

type Type string

const (
	Auto       Type = "auto"
	MaxContent Type = "max-content"
	MinContent Type = "min-content"
)

type Func string

func Flex(frac int) Type {
	return Type(internal.Stringify(frac, "fr"))
}

func New(value1, value2 interface{}) Func {
	return Func("minmax(" + internal.Stringify(value1, "px") + "," + internal.Stringify(value2, "px") + ")")
}
