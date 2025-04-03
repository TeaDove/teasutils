package reflect_utils

import (
	"fmt"
	"reflect"
	"runtime"
)

func GetTypesStringRepresentation(v any) string {
	if v == nil {
		return "nil"
	}

	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}

	return t.Name()
}

func GetFunctionName(i any) string {
	return fmt.Sprintf("%s %T", runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), i)
}
