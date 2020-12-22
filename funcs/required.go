package funcs

import (
	"reflect"
	"time"
)

//Define RequiredFunc struct
type requiredFunc struct {
}

//RequiredFunc
func RequiredFunc() VFunc {
	return &requiredFunc{}
}

//Accept
func (r *requiredFunc) Accept(typ reflect.Type) bool {
	return true
}

//Pass
func (r *requiredFunc) Pass(value reflect.Value) bool {
	switch value.Type().Kind() {
	case reflect.Array, reflect.Slice:
		return !value.IsNil() && value.Len() > 0
	}
	if value.Type() == reflect.TypeOf(time.Time{}) {
		return !value.Interface().(time.Time).IsZero()
	}
	return value.IsValid()
}
