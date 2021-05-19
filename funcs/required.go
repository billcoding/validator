package funcs

import (
	"reflect"
	"time"
)

// requiredFunc struct
type requiredFunc struct {
}

// RequiredFunc method
func RequiredFunc() VFunc {
	return &requiredFunc{}
}

// Accept method
func (r *requiredFunc) Accept(typ reflect.Type) bool {
	return true
}

// Pass method
func (r *requiredFunc) Pass(value reflect.Value) bool {
	switch value.Type().Kind() {
	case reflect.Ptr:
		return !value.IsNil()
	case reflect.Array, reflect.Slice:
		return !value.IsNil() && !value.IsZero() && value.Len() > 0
	}
	if value.Type() == reflect.TypeOf(time.Time{}) {
		return !value.Interface().(time.Time).IsZero()
	}
	return value.IsValid()
}
