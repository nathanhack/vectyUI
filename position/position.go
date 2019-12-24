package position

import (
	"github.com/nathanhack/vectyUI/internal"
)

type Type string

const (
	Left   Type = "left"
	Center Type = "center"
	Right  Type = "right"
	Top    Type = "top"
	Bottom Type = "bottom"
)

func Len(length interface{}) Type {
	return Type(internal.Stringify(length) + "px")
}

func Percent(percent interface{}) Type {
	return Type(internal.Stringify(percent) + "%")
}

type Value []Type

func New(position ...Type) Value {
	return Value(position)
}
