package internal

import (
	"reflect"
	"strconv"
)

func Stringify(value interface{}) string {
	return stringify(reflect.TypeOf(value), reflect.ValueOf(value))
}

//stringify calls recursively until we get the type and value
func stringify(typeOf reflect.Type, valueOf reflect.Value) string {
	switch typeOf.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(valueOf.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(valueOf.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(valueOf.Float(), 'f', -1, 64)
	case reflect.String:
		return valueOf.String()
	case reflect.Ptr:
		i := reflect.Indirect(valueOf)
		return stringify(i.Type(), i)

	default:
		return ""
	}
}

func StringifyPost(postfix string, values ...interface{}) []string {
	stringValues := make([]string, 0)
	for _, l := range values {
		stringValues = append(stringValues, Stringify(l)+postfix)
	}
	return stringValues
}
