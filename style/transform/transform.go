package transform

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
)

type Type string

const (
	None    Type = "none"
	Initial Type = "initial"
	Inherit Type = "inherit"
	Unset   Type = "unset"
)

var styleNames = []string{"transform", "-webkit-transform"}

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

type Value string

func (v Value) Apply(h *vecty.HTML) {
	for _, s := range styleNames {
		vecty.Style(s, string(v)).Apply(h)
	}
}

func (v Value) AddTo(m map[string]string) {
	for _, s := range styleNames {
		m[s] = string(v)
	}
}

func Percent(percent interface{}) string {
	return internal.Stringify(percent, "%")
}

func Pixels(pixels interface{}) string {
	return internal.Stringify(pixels, "px")
}

func Deg(degrees interface{}) string {
	return internal.Stringify(degrees, "deg")
}

func Rad(rad interface{}) string {
	return internal.Stringify(rad, "rad")
}

//<matrix()> = matrix( <number>#{6} )
func Matrix(number0,
	number1,
	number2,
	number3,
	number4,
	number5 interface{}) Type {
	return "matrix(" +
		internal.Stringify(number0, "") + "," +
		internal.Stringify(number1, "") + "," +
		internal.Stringify(number2, "") + "," +
		internal.Stringify(number3, "") + "," +
		internal.Stringify(number4, "") + "," +
		internal.Stringify(number5, "") + "," +
		")"
}

//<matrix3d()> = matrix3d( <number>#{16} )
func Matrix3d(number0,
	number1,
	number2,
	number3,
	number4,
	number5,
	number6,
	number7,
	number8,
	number9,
	number10,
	number11,
	number12,
	number13,
	number14,
	number15 interface{}) Type {
	return "matrix3d(" +
		internal.Stringify(number0, "") + "," +
		internal.Stringify(number1, "") + "," +
		internal.Stringify(number2, "") + "," +
		internal.Stringify(number3, "") + "," +
		internal.Stringify(number4, "") + "," +
		internal.Stringify(number5, "") + "," +
		internal.Stringify(number6, "") + "," +
		internal.Stringify(number7, "") + "," +
		internal.Stringify(number8, "") + "," +
		internal.Stringify(number9, "") + "," +
		internal.Stringify(number10, "") + "," +
		internal.Stringify(number11, "") + "," +
		internal.Stringify(number12, "") + "," +
		internal.Stringify(number13, "") + "," +
		internal.Stringify(number14, "") + "," +
		internal.Stringify(number15, "") + "," +
		")"
}

//<translate()> = translate( <length-percentage> , <length-percentage>? )
//translate favors pixel translation if not specified
func Translate(lengthOrPercentage ...interface{}) Type {
	l := len(lengthOrPercentage)
	switch l {
	case 0:
		return ""
	case 1:
		return "translate(" + internal.Stringify(lengthOrPercentage[0], "px") + ")"
	default:
		return "translate(" + internal.Stringify(lengthOrPercentage[0], "px") + "," + internal.Stringify(lengthOrPercentage[1], "px") + ")"
	}
}

//<translateX()> = translateX( <length-percentage> )
func TranslateX(lengthOrPercentage interface{}) Type {
	return "translateX(" + internal.Stringify(lengthOrPercentage, "px") + ")"
}

//<translateY()> = translateY( <length-percentage> )
func TranslateY(lengthOrPercentage interface{}) Type {
	return "translateY(" + internal.Stringify(lengthOrPercentage, "px") + ")"
}

//<translateZ()> = translateZ( <length> )
func TranslateZ(length interface{}) Type {
	return "translateZ(" + internal.Stringify(length, "px") + ")"
}

//<translate3d()> = translate3d( <length-percentage> , <length-percentage> , <length> )
func Translate3d(lenPercentX, lenPercentY, lenZ interface{}) Type {
	return "translate3d(" + internal.Stringify(lenPercentX, "px") + "," + internal.Stringify(lenPercentY, "px") + "," + internal.Stringify(lenZ, "px") + ")"
}

//<scale()> = scale( <number> , <number>? )
func Scale(numbers ...interface{}) Type {
	l := len(numbers)
	switch l {
	case 0:
		return ""
	case 1:
		return "scale(" + internal.Stringify(numbers[0], "") + ")"
	default:
		//if equal to 2 or more we only take the first two
		return "scale(" + internal.Stringify(numbers[0], "") + "," + internal.Stringify(numbers[1], "") + ")"
	}
}

//<scaleX()> = scaleX( <number> )
func ScaleX(number interface{}) Type {
	return "scaleX(" + internal.Stringify(number, "") + ")"
}

//<scaleY()> = scaleY( <number> )
func ScaleY(number interface{}) Type {
	return "scaleY(" + internal.Stringify(number, "") + ")"
}

//<scaleZ()> = scaleZ( <number> )
func ScaleZ(number interface{}) Type {
	return "scaleZ(" + internal.Stringify(number, "") + ")"
}

//<scale3d()> = scale3d( <number> , <number> , <number> )
func Scale3d(numberX, numberY, numberZ interface{}) Type {
	return "scale3d(" + internal.Stringify(numberX, "") + "," + internal.Stringify(numberY, "") + "," + internal.Stringify(numberZ, "") + ")"
}

//<skew()> = skew( [ <angle> | <zero> ] , [ <angle> | <zero> ]? )
func Skew(angles ...interface{}) Type {
	l := len(angles)
	switch l {
	case 0:
		return ""
	case 1:
		return "skew(" + internal.Stringify(angles[0], "deg") + ")"
	default:
		//if equal to 2 or more we only take the first two
		return "skew(" + internal.Stringify(angles[0], "deg") + "," + internal.Stringify(angles[1], "deg") + ")"
	}
}

//<skewX()> = skewX( [ <angle> | <zero> ] )
func SkewX(angle interface{}) Type {
	return "skewX(" + internal.Stringify(angle, "deg") + ")"
}

//<skewY()> = skewY( [ <angle> | <zero> ] )
func SkewY(angle interface{}) Type {
	return "skewY(" + internal.Stringify(angle, "deg") + ")"
}

//<rotate()> = rotate( [ <angle> | <zero> ] )
func Rotate(angle interface{}) Type {
	return "rotate(" + internal.Stringify(angle, "deg") + ")"
}

//<rotateX()> = rotateX( [ <angle> | <zero> ] )
func RotateX(angle interface{}) Type {
	return "rotateX(" + internal.Stringify(angle, "deg") + ")"
}

//<rotateY()> = rotateY( [ <angle> | <zero> ] )
func RotateY(angle interface{}) Type {
	return "rotateY(" + internal.Stringify(angle, "deg") + ")"
}

//<rotateZ()> = rotateZ( [ <angle> | <zero> ] )
func RotateZ(angle interface{}) Type {
	return "rotateZ(" + internal.Stringify(angle, "deg") + ")"
}

//<rotate3d()> = rotate3d( <number> , <number> , <number> , [ <angle> | <zero> ] )
func Rotate3d(numberX, numberY, numberZ, angle interface{}) Type {
	return "rotate3d(" + internal.Stringify(numberX, "") + "," + internal.Stringify(numberY, "") + "," + internal.Stringify(numberZ, "") + "," + internal.Stringify(angle, "deg") + ")"
}

//<perspective()> = perspective( <length> )
func Perspective(length interface{}) Type {
	return "perspective(" + internal.Stringify(length, "px") + ")"
}
