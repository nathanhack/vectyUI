package internal

import (
	"reflect"
	"strconv"
)

func Stringify(value interface{}, nonStringPostFix string) string {
	return stringify(reflect.TypeOf(value), reflect.ValueOf(value), nonStringPostFix)
}

//stringify calls recursively until we get the type and value
func stringify(typeOf reflect.Type, valueOf reflect.Value, nonStringPostFix string) string {
	switch typeOf.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(valueOf.Int(), 10) + nonStringPostFix
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(valueOf.Uint(), 10) + nonStringPostFix
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(valueOf.Float(), 'f', -1, 64) + nonStringPostFix
	case reflect.String:
		return valueOf.String()
	case reflect.Ptr:
		i := reflect.Indirect(valueOf)
		return stringify(i.Type(), i, nonStringPostFix)

	default:
		return ""
	}
}

func StringifyList(nonStringPostFix string, values ...interface{}) []string {
	stringValues := make([]string, 0)
	for _, value := range values {
		stringValues = append(stringValues, stringify(reflect.TypeOf(value), reflect.ValueOf(value), nonStringPostFix))
	}
	return stringValues
}
