package reflectutils

import (
	"fmt"
	"reflect"
	"runtime"
)

// GetFunctionName returns the fully-qualified name and dynamic type of the
// function i, e.g. "pkg.Foo func()". If i is not a function, or its name
// cannot be resolved, only the type is returned (e.g. "int").
func GetFunctionName(i any) string {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Func {
		return fmt.Sprintf("%T", i)
	}

	fn := runtime.FuncForPC(v.Pointer())
	if fn == nil {
		return fmt.Sprintf("%T", i)
	}

	return fmt.Sprintf("%s %T", fn.Name(), i)
}
