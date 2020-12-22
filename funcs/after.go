package funcs

import (
	"github.com/billcoding/reflectx"
	"reflect"
	"time"
)

//Define afterFunc struct
type afterFunc struct {
	after time.Time
}

//AfterFunc
func AfterFunc(after time.Time) VFunc {
	return &afterFunc{after}
}

//Accept
func (f *afterFunc) Accept(typ reflect.Type) bool {
	return typ.Kind() == reflect.String || typ == reflect.TypeOf(time.Time{})
}

//Pass
func (f *afterFunc) Pass(value reflect.Value) bool {
	if f.after.IsZero() {
		return true
	}
	if value.Type().Kind() == reflect.String {
		return f.Pass(reflect.ValueOf(reflectx.ParseTime(value.String())))
	} else if value.Type() == reflect.TypeOf(time.Time{}) {
		return f.after.Before(value.Interface().(time.Time))
	}
	return true
}
